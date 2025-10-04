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
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
	Public  bool      `json:"public"`
}

type DiaryEntriesResponse struct {
	Entries []DiaryEntry `json:"entries"`
}

func DiaryEntries(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	isAuthenticated := r.Header.Get("X-Diary-Auth") == "authenticated"
	
	redisURL := os.Getenv("REDIS_URL")
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
	
	if len(entries) == 0 {
		entries = []DiaryEntry{
			{
				ID:      1,
				Title:   "Starting Fresh",
				Content: "Built this portfolio from scratch today. Go + HTMX, no frameworks. Feels good to keep things simple.\n\nThe old money aesthetic just clicked. Not flashy, just solid. Like good code should be.\n\nNext up: more projects, maybe write about Rust learnings.",
				Date:    time.Date(2025, 10, 4, 23, 47, 0, 0, time.UTC),
				Public:  true,
			},
			{
				ID:      2,
				Title:   "Private Thoughts",
				Content: "This is a private entry. Only visible when logged in. Contains personal reflections and sensitive information.",
				Date:    time.Date(2025, 10, 4, 23, 50, 0, 0, time.UTC),
				Public:  false,
			},
		}
	}
	
	if !isAuthenticated {
		for i := range entries {
			if !entries[i].Public {
				entries[i].Content = "[This entry is private. Login to view full content.]"
			}
		}
	}
	
	json.NewEncoder(w).Encode(DiaryEntriesResponse{Entries: entries})
}
