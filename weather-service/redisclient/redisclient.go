package redisclient

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) (*redis.Client, error) {
	options := &redis.Options{
		Addr: addr,
	}
	client := redis.NewClient(options)

	// Test the connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
