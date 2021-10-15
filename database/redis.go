package database

import (
	"context"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	log2 "github.com/sirupsen/logrus"
	config2 "notebook/config"
)

var RedisContext = context.Background()
var Redis *redis.Client

func init() {
	dbLogger := log2.WithFields(log2.Fields{})
	cfg := config2.Conf
	dbLogger.Debugf("%v\n", cfg.Redis)
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
