package main

import (
	"log"
	"net/http"
	"weather-service/cache"
	"weather-service/config"
	"weather-service/handlers"
	"weather-service/redisclient"
	"weather-service/weather"
)

func main() {
	cfg := config.NewConfig()
	redisAddr := cfg.Get("REDIS_ADDR")

	client, err := redisclient.NewRedisClient(redisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer client.Close()

	cacheClient := cache.NewRedisCacheClient(client)
	weatherService := weather.NewWeatherService(cacheClient)
	handlers.InitWeatherHandler(weatherService)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.PingHandler)
	mux.HandleFunc("/weather", handlers.WeatherHandler)

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
