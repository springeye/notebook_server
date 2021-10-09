package db

import (
	"context"
	redis "github.com/go-redis/redis/v8"
)

var RedisContext = context.Background()
var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
