package api

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Visitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html")
	
	redisURL := os.Getenv("REDIS_URL")
	var count int64 = 1
	
	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err == nil {
			rdb := redis.NewClient(opt)
			count, _ = rdb.Incr(ctx, "visitor_count").Result()
			rdb.Close()
		}
	}
	
	response := fmt.Sprintf(`
		<div class="stat-number">%d</div>
		<div class="stat-label">Visitors</div>
	`, count)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
