package handlers

import (
	"encoding/json"
	"net/http"
	"weather-service/weather"
)

var weatherService *weather.WeatherService

func InitWeatherHandler(ws *weather.WeatherService) {
	weatherService = ws
}

// WeatherHandler handles the /weather endpoint
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the required query parameter "address" is present
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing required query parameter: address", http.StatusBadRequest)
		return
	}

	// Get weather information
	weatherInfo, err := weatherService.GetWeather(address)
	if err != nil {
		http.Error(w, "Failed to get weather information", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header and write response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weatherInfo); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
