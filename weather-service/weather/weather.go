package weather

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "time"
    "weather-service/cache"

    "github.com/redis/go-redis/v9"
)

type WeatherService struct {
    cacheClient cache.CacheClient
    apiKey      string
}

func NewWeatherService(cacheClient cache.CacheClient, apiKey string) *WeatherService {
    return &WeatherService{cacheClient: cacheClient, apiKey: apiKey}
}

func (ws *WeatherService) GetWeather(address string) (map[string]interface{}, error) {
    ctx := context.Background()

    // Try to get weather information from cache
    val, err := ws.cacheClient.Get(ctx, address)
    if errors.Is(err, redis.Nil) {
        // Address not found in cache, make an HTTP call to get weather information
        weatherInfo, err := ws.fetchWeatherFromAPI(address)
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

func (ws *WeatherService) fetchWeatherFromAPI(address string) (map[string]interface{}, error) {
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