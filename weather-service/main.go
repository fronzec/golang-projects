package main

import (
	"log"
	"net/http"
	"os"
	"weather-service/cache"
	"weather-service/handlers"
	"weather-service/redisclient"
	"weather-service/weather"
)

func main() {
	// Load environment variables
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR environment variable is required")
	}

	// Initialize Redis client
	client, err := redisclient.NewRedisClient(redisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer client.Close()

	// Initialize cache client
	cacheClient := cache.NewRedisCacheClient(client)

	// Initialize weather service
	weatherService := weather.NewWeatherService(cacheClient)

	// Initialize handlers
	handlers.InitWeatherHandler(weatherService)

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
