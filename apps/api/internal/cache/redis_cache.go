package cache

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

type Cache interface {
    Get(ctx context.Context, key string) (string, error)
    Set(ctx context.Context, key string, value string, ttl time.Duration) error
}

type redisCache struct {
    rdb *redis.Client
}

func New(rdb *redis.Client) Cache {
    return &redisCache{rdb: rdb}
}

func (c *redisCache) Get(ctx context.Context, key string) (string, error) {
    return c.rdb.Get(ctx, key).Result()
}

func (c *redisCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
    return c.rdb.Set(ctx, key, value, ttl).Err()
}
