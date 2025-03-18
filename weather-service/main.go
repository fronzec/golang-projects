package main

import (
	"net/http"
	"time"
	"weather-service/cache"
	"weather-service/config"
	"weather-service/handlers"
	"weather-service/ratelimit"
	"weather-service/redisclient"
	"weather-service/weather"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// zerolog config in UTC
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	cfg := config.NewConfig()

	client, err := redisclient.NewRedisClient(cfg)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to connect to Redis")
	}
	defer client.Close()

	cacheClient := cache.NewRedisCacheClient(client)
	weatherService := weather.NewWeatherService(cacheClient, cfg.Get("ENV_VISUAL_CROSSING_API_KEY"))
	handlers.InitWeatherHandler(weatherService)

	// Create a new rate limiter: 2 requests per second with burst of 5
	limiter := ratelimit.NewIPRateLimiter(2, 5)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.PingHandler)
	// Apply rate limiting middleware to the weather endpoint
	mux.HandleFunc("/weather", ratelimit.Middleware(handlers.WeatherHandler, limiter))

	log.Info().Msg("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error starting the server")
	}
}
