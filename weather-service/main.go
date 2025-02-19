package main

import (
	"log"
	"net/http"
	"weather-service/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Register handler for /ping
	mux.HandleFunc("/ping", handlers.PingHandler)

	// Register handler for /weather
	mux.HandleFunc("/weather", handlers.WeatherHandler)

	// Start the server
	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
