package weather

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"weather-service/cache"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	cacheClient cache.CacheClient
	apiKey      string
}

func NewWeatherService(cacheClient cache.CacheClient, apiKey string) *Service {
	return &Service{cacheClient: cacheClient, apiKey: apiKey}
}

func generateCacheKey(address string) string {
	hash := sha256.Sum256([]byte(address))
	return fmt.Sprintf("weather:%s", hex.EncodeToString(hash[:]))
}

func (ws *Service) GetWeather(address string) (map[string]interface{}, error) {
	ctx := context.Background()
	cacheKey := generateCacheKey(address)

	// Try to get weather information from cache
	val, err := ws.cacheClient.Get(ctx, cacheKey)
	if errors.Is(err, redis.Nil) {
		// Address not found in cache, make an HTTP call to get weather information
		weatherInfo, err := ws.fetchWeatherFromAPI(address)
		if err != nil {
			return nil, err
		}

		// Store the weather information in cache with a random expiration time between 5 and 15 minutes
		weatherJSON, _ := json.Marshal(weatherInfo)
		expiration := time.Duration(5+rand.Intn(11)) * time.Minute
		ws.cacheClient.Set(ctx, cacheKey, string(weatherJSON), expiration)

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

func (ws *Service) fetchWeatherFromAPI(address string) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=us&include=days&key=%s&contentType=json", address, ws.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: %s", resp.Status)
	}

	var weatherInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&weatherInfo); err != nil {
		return nil, err
	}

	return weatherInfo, nil
}
