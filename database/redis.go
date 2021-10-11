package database

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"os"
)

var RedisContext = context.Background()
var Redis *redis.Client

func init() {
	host := "localhost"
	port := "6379"
	if os.Getenv("REDIS_HOST") != "" {
		host = os.Getenv("REDIS_HOST")
	}
	if os.Getenv("REDIS_PORT") != "" {
		port = "6379"
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
