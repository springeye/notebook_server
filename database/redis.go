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
	if os.Getenv("notebook_redis") != "" {
		host = os.Getenv("notebook_redis")
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", host),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
