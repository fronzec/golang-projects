package redisclient

import (
	"context"
	"weather-service/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	options := &redis.Options{
		Addr: cfg.Get("REDIS_ADDR"),
	}

	if cfg.Exists("REDIS_PASSWORD") {
		options.Password = cfg.Get("REDIS_PASSWORD")
	}

	if cfg.Exists("REDIS_DB") {
		options.DB = cfg.GetInt("REDIS_DB")
	}

	client := redis.NewClient(options)

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
