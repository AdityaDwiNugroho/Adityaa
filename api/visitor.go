package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Visitor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html")
	
	kvURL := os.Getenv("KV_REST_API_URL")
	kvToken := os.Getenv("KV_REST_API_TOKEN")
	
	var count int64 = 1
	
	if kvURL != "" && kvToken != "" {
		resp, err := http.Get(kvURL + "/incr/visitor_count?_token=" + kvToken)
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			var result struct {
				Result int64 `json:"result"`
			}
			if err := json.NewDecoder(resp.Body).Decode(&result); err == nil {
				count = result.Result
			}
		}
	}
	
	response := fmt.Sprintf(`
		<div class="stat-number">%d</div>
		<div class="stat-label">Visitors</div>
	`, count)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
