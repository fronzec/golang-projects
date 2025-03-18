package handlers

import (
	"encoding/json"
	"net/http"
)

// PingHandler responds to the /ping route with an empty JSON response.
func PingHandler(w http.ResponseWriter, r *http.Request) {
	// Verify that the method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Respond with an empty JSON
	_ = json.NewEncoder(w).Encode(map[string]interface{}{})
}
