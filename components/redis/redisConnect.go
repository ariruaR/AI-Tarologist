package redis

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type redisClient struct {
	client *redis.Client
}

func NewClient() *redisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &redisClient{
		client: rdb,
	}
}

func (r *redisClient) Setter(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
func (r *redisClient) Getter(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
