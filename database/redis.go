package database

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	config2 "notebook/config"
)

var RedisContext = context.Background()
var Redis *redis.Client

func init() {
	cfg := config2.Config

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
