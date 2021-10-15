package config

import (
	"github.com/spf13/viper"
	"time"
)

type CacheType string

type DataBaseType string

const (
	Memory CacheType = "memory"
	Redis            = "redis"
)
const (
	SQLITE DataBaseType = "sqlite"
	MYSQL               = "mysql"
)

type AppConfig struct {
	Server *struct {
		Port int
	}
	Database *struct {
		Type   DataBaseType
		Logger *struct {
			Level string
		}
		Sqlite *struct {
			File string
		}
		MySQL *struct {
			Host     string
			Port     int
			Name     string
			Username string
			Password string
		}
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

var Conf *AppConfig

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath("/app")
	v.AddConfigPath("./")
	v.AddConfigPath("../")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	err := v.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
