package api

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
}

func DiaryAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(LoginResponse{Success: false})
		return
	}
	
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		json.NewEncoder(w).Encode(LoginResponse{Success: false})
		return
	}
	
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		json.NewEncoder(w).Encode(LoginResponse{Success: false})
		return
	}
	
	rdb := redis.NewClient(opt)
	defer rdb.Close()
	
	storedPassword, err := rdb.Get(ctx, "diary_password").Result()
	if err != nil {
		json.NewEncoder(w).Encode(LoginResponse{Success: false})
		return
	}
	
	if req.Password == storedPassword {
		json.NewEncoder(w).Encode(LoginResponse{
			Success: true,
			Token:   "authenticated",
		})
	} else {
		json.NewEncoder(w).Encode(LoginResponse{Success: false})
	}
}
