package redisclient

import (
	"context"
	"os"
	"strconv"
	"weather-service/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	redisAddrKey := cfg.Get("ENV_REDIS_ADDR")
	redisPasswordKey := cfg.Get("ENV_REDIS_PASSWORD")
	redisDBKey := cfg.Get("ENV_REDIS_DB")

	options := &redis.Options{
		Addr: os.Getenv(redisAddrKey),
	}

	if cfg.Exists("ENV_REDIS_PASSWORD") {
		options.Password = os.Getenv(redisPasswordKey)
	}

	if cfg.Exists("ENV_REDIS_DB") {
		dbValue, err := strconv.Atoi(os.Getenv(redisDBKey))
		if err == nil {
			options.DB = dbValue
		}
	}

	client := redis.NewClient(options)

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
