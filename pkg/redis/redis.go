package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) Redis {
	return Redis{
		client: client,
	}
}

func (r *Redis) Inspect(ctx context.Context, prefix string, key string) (bool, []byte) {
	primeKey := createKey(prefix, key)
	result, err := r.client.Get(ctx, primeKey).Bytes()
	if err != nil {
		return false, nil
	}
	return true, result
}

func (r *Redis) Set(ctx context.Context, prefix string, expiration time.Duration, key string, value any) error {
	primeKey := createKey(prefix, key)
	err := r.client.Set(ctx, primeKey, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Validate(ctx context.Context, prefix string, key string, value string) bool {
	primeKey := createKey(prefix, key)
	result, err := r.client.Get(ctx, primeKey).Result()
	if err == nil {
		if result == value {
			return true
		}
	}
	return false
}

func (r *Redis) Remove(ctx context.Context, prefix string, key string) error {
	primeKey := createKey(prefix, key)

	_, err := r.client.Get(ctx, primeKey).Result()
	if err != nil {
		return err
	}
	err = r.client.Del(ctx, primeKey).Err()
	if err != nil {
		return err
	}
	return nil
}

func createKey(prefix string, key string) string {
	primeKey := fmt.Sprintf("%s:%s", prefix, key)
	return primeKey
}
