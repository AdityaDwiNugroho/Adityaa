package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("KV_REST_API_URL"),
		Password: os.Getenv("KV_REST_API_TOKEN"),
	})
}

func Visitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html")
	
	rdb := getRedisClient()
	defer rdb.Close()
	
	count, err := rdb.Incr(ctx, "visitor_count").Result()
	if err != nil {
		count = 1
	}
	
	response := fmt.Sprintf(`
		<div class="stat-number">%d</div>
		<div class="stat-label">Visitors</div>
	`, count)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
