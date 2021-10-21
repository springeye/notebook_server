package store

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisAdapter struct {
	Client *redis.Client
}

func (g RedisAdapter) Get(key string) interface{} {
	cmd := g.Client.Get(context.Background(), key)
	cmd.Err()
	return cmd.Val()
}

func (g RedisAdapter) Set(key string, object interface{}) {
	g.Client.Set(context.Background(), key, object, -1)
}

func (g RedisAdapter) Remove(key string) {
	g.Client.Del(context.Background(), key)
}
