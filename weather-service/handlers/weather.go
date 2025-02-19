package handlers

import (
	"encoding/json"
	"net/http"
)

// WeatherHandler handles the /weather endpoint
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the required query parameter "address" is present
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing required query parameter: address", http.StatusBadRequest)
		return
	}

	// Hardcoded response example as specified in the API_Specification_v1.yml
	response := map[string]interface{}{
		"location":    address, // Use the provided address in the response
		"temperature": 25.5,
		"humidity":    60,
		"condition":   "Partly cloudy",
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
