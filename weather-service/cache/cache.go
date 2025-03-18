package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
}

type RedisCacheClient struct {
	client *redis.Client
}

func NewRedisCacheClient(client *redis.Client) *RedisCacheClient {
	return &RedisCacheClient{client: client}
}

func (r *RedisCacheClient) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisCacheClient) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
