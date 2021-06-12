package main

import (
	"context"
	"github.com/go-redis/redis"
	"time"
)

type IRedisCache interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Set(ctx context.Context, key string, value interface{}, expirationTime time.Duration) (interface{}, error)
	Remove(ctx context.Context, key string) error
}

type RedisCache struct {
	RedisClient *redis.Client
}

func NewRedisCache(redisClient *redis.Client) IRedisCache {
	return &RedisCache{RedisClient: redisClient}
}

func(c *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	value, err := c.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}

func(c *RedisCache) Set(ctx context.Context, key string, value interface{}, expirationTime time.Duration) (interface{}, error) {
	value, err := c.RedisClient.Set(ctx, key, value, expirationTime).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}

func(c *RedisCache) Remove(ctx context.Context, key string) error {
	return c.RedisClient.Del(ctx, key).Err()
}