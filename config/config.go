package config

import (
	log2 "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
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
	Logger *struct {
		ShowSql bool `mapstructure:"show_sql"`
		Level   string
	}
	Database *struct {
		Type DataBaseType

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
	log2.SetLevel(log2.TraceLevel)
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

	if Conf.Logger.Level == "info" {
		log2.SetLevel(log2.InfoLevel)
	} else if Conf.Logger.Level == "warn" {
		log2.SetLevel(log2.WarnLevel)
	} else if Conf.Logger.Level == "error" {
		log2.SetLevel(log2.ErrorLevel)
	}
	log2.SetOutput(os.Stdout)
	log2.SetFormatter(&log2.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})

}
