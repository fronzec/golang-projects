package main

import (
	"net/http"
	"time"
	"weather-service/cache"
	"weather-service/config"
	"weather-service/handlers"
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
	weatherService := weather.NewWeatherService(cacheClient, cfg.Get("SECRET_VISUAL_CROSSING"))
	handlers.InitWeatherHandler(weatherService)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.PingHandler)
	mux.HandleFunc("/weather", handlers.WeatherHandler)

	log.Info().Msg("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error starting the server")
	}
}
