package weather

import (
	"context"
	"encoding/json"
	"time"
	"weather-service/cache"

	"github.com/redis/go-redis/v9"
)

type WeatherService struct {
	cacheClient cache.CacheClient
}

func NewWeatherService(cacheClient cache.CacheClient) *WeatherService {
	return &WeatherService{cacheClient: cacheClient}
}

func (ws *WeatherService) GetWeather(address string) (map[string]interface{}, error) {
	ctx := context.Background()

	// Try to get weather information from cache
	val, err := ws.cacheClient.Get(ctx, address)
	if err == redis.Nil {
		// Address not found in cache, make an HTTP call to get weather information
		weatherInfo, err := fetchWeatherFromAPI(address)
		if err != nil {
			return nil, err
		}

		// Store the weather information in cache
		weatherJSON, _ := json.Marshal(weatherInfo)
		ws.cacheClient.Set(ctx, address, string(weatherJSON), 10*time.Minute)

		return weatherInfo, nil
	} else if err != nil {
		return nil, err
	}

	// Address found in cache, unmarshal the JSON
	var weatherInfo map[string]interface{}
	if err := json.Unmarshal([]byte(val), &weatherInfo); err != nil {
		return nil, err
	}

	return weatherInfo, nil
}

func fetchWeatherFromAPI(address string) (map[string]interface{}, error) {
	// Simulate an HTTP call to get weather information
	// In a real implementation, you would make an actual HTTP request here
	return map[string]interface{}{
		"location":    address,
		"temperature": 25.5,
		"humidity":    60,
		"condition":   "Partly cloudy",
	}, nil
}
