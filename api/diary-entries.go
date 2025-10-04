package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type DiaryEntry struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Date      time.Time `json:"date"`
	IsPrivate bool      `json:"isPrivate"`
}

func DiaryEntries(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	// Check if user is authenticated
	authHeader := r.Header.Get("Authorization")
	isAuthenticated := authHeader != "" && authHeader == "Bearer true"
	
	redisURL := os.Getenv("REDIS_URL")
	
	// Handle POST - Create new entry (requires auth)
	if r.Method == "POST" {
		if !isAuthenticated {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
			return
		}
		
		var newEntry DiaryEntry
		if err := json.NewDecoder(r.Body).Decode(&newEntry); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
			return
		}
		
		// Get existing entries
		var entries []DiaryEntry
		if redisURL != "" {
			opt, _ := redis.ParseURL(redisURL)
			rdb := redis.NewClient(opt)
			defer rdb.Close()
			
			data, err := rdb.Get(ctx, "diary_entries").Result()
			if err == nil {
				json.Unmarshal([]byte(data), &entries)
			}
		}
		
		// Add new entry
		newEntry.ID = len(entries) + 1
		newEntry.Date = time.Now()
		entries = append([]DiaryEntry{newEntry}, entries...) // Prepend
		
		// Save to Redis
		if redisURL != "" {
			opt, _ := redis.ParseURL(redisURL)
			rdb := redis.NewClient(opt)
			defer rdb.Close()
			
			data, _ := json.Marshal(entries)
			rdb.Set(ctx, "diary_entries", string(data), 0)
		}
		
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"entry":   newEntry,
		})
		return
	}
	
	// Handle GET - List entries
	var entries []DiaryEntry
	
	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err == nil {
			rdb := redis.NewClient(opt)
			defer rdb.Close()
			
			data, err := rdb.Get(ctx, "diary_entries").Result()
			if err == nil {
				json.Unmarshal([]byte(data), &entries)
			}
		}
	}
	
	// Default entries if nothing in Redis
	if len(entries) == 0 {
		entries = []DiaryEntry{
			{
				ID:        1,
				Title:     "Starting Fresh",
				Content:   "Built this portfolio from scratch today. Go + HTMX, no frameworks. Feels good to keep things simple.\n\nThe old money aesthetic just clicked. Not flashy, just solid. Like good code should be.\n\nNext up: more projects, maybe write about Rust learnings.",
				Date:      time.Date(2025, 10, 4, 23, 47, 0, 0, time.UTC),
				IsPrivate: false,
			},
			{
				ID:        2,
				Title:     "Private Thoughts",
				Content:   "This is a private entry. Only visible when logged in. Contains personal reflections and sensitive information.",
				Date:      time.Date(2025, 10, 4, 23, 50, 0, 0, time.UTC),
				IsPrivate: true,
			},
		}
	}
	
	// Censor private entries if not authenticated
	if !isAuthenticated {
		for i := range entries {
			if entries[i].IsPrivate {
				entries[i].Content = "This entry is private. Login to view."
			}
		}
	}
	
	json.NewEncoder(w).Encode(entries)
}
