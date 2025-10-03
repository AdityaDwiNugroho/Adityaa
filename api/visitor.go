package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
)

// Simple in-memory counter (will reset on redeploy, use Redis/DB for production)
var visitorCount uint64

func Handler(w http.ResponseWriter, r *http.Request) {
	// Increment visitor count
	count := atomic.AddUint64(&visitorCount, 1)
	
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html")
	
	// Return HTML fragment for HTMX (for floating stat)
	response := fmt.Sprintf(`
		<div class="stat-number">%d</div>
		<div class="stat-label">Visitors</div>
	`, count)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
