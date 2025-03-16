package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"
	"weather-service/cache"
	"weather-service/config"
	"weather-service/handlers"
	"weather-service/redisclient"
	"weather-service/weather"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
)

// / IPRateLimiter Create a custom limiter structure
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// / NewIPRateLimiter Create a new rate limiter that allows 2 requests per second with burst of 5
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}
}

// / GetLimiter Get or create a limiter for an IP address
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter, exists := i.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(i.r, i.b)
		i.ips[ip] = limiter
	}

	return limiter
}

// / rateLimitMiddleware Middleware function to handle rate limiting
func rateLimitMiddleware(next http.HandlerFunc, limiter *IPRateLimiter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		if !limiter.GetLimiter(ip).Allow() {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"error":   "Rate limit exceeded",
				"message": "Too many requests, please try again later",
			})
			return
		}
		next.ServeHTTP(w, r)
	}
}

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
	limiter := NewIPRateLimiter(2, 5)

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handlers.PingHandler)
	// Apply rate limiting middleware to the weather endpoint
	mux.HandleFunc("/weather", rateLimitMiddleware(handlers.WeatherHandler, limiter))

	log.Info().Msg("Server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal().
			Err(err).
			Msg("Error starting the server")
	}
}
