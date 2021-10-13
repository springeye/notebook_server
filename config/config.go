package config

import (
	"github.com/spf13/viper"
	"time"
)

type CacheType string

const (
	Memory CacheType = "memory"
	Redis            = "redis"
)

type AppConfig struct {
	Server *struct {
		Port int
	}
	Redis *struct {
		Host string
		Port int
	}
	Cache *struct {
		Type       CacheType
		Expiration time.Duration
	}
}

var Config *AppConfig

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("./")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	err := v.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}
