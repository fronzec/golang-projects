package handlers

import (
	"encoding/json"
	"net/http"
	"weather-service/types"
	"weather-service/weather"
)

var weatherService *weather.Service

func InitWeatherHandler(ws *weather.Service) {
	weatherService = ws
}

/// ErrorResponse structure for error responses
type ErrorResponse struct {
	Message string `json:"message"`
}

/// writeJSONResponse writes successful JSON responses
func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

/// writeJSONError writes JSON error responses
func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

/// WeatherHandler handles the /weather endpoint
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		writeJSONError(w, "Missing required query parameter: address", http.StatusBadRequest)
		return
	}

	weatherInfo, err := weatherService.GetWeather(address)
	if err != nil {
		if errorResponse, ok := err.(*types.ErrorResponse); ok {
			writeJSONError(w, errorResponse.Message, errorResponse.HTTPCode)
			return
		}
		writeJSONError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, weatherInfo, http.StatusOK)
}
