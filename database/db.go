package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log2 "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	config2 "notebook/config"
	"notebook/model"
	"os"
	"path/filepath"
	"time"
)

var Database *gorm.DB

func init() {
	dbLogger := log2.WithFields(log2.Fields{})
	var err error
	logLevel := logger.Error
	if gin.Mode() != "release" {
		logLevel = logger.Info
	}
	config := gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logLevel,    // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		),
	}

	dbconf := config2.Conf.Database
	log.Printf("Database type: %s", dbconf.Type)
	if dbconf.Type == config2.SQLITE {
		b, _ := json.Marshal(dbconf.Sqlite)
		dbLogger.Info("Database conf: %s", string(b))
		dbpath, err := filepath.Abs(dbconf.Sqlite.File)
		if err != nil {
			panic(err)
		}
		Database, err = gorm.Open(sqlite.Open(dbpath), &config)
	} else if dbconf.Type == config2.MYSQL {
		b, _ := json.Marshal(dbconf.MySQL)
		dbLogger.Printf("Database conf: %s", string(b))
		db := dbconf.MySQL
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			db.Username, db.Password,
			db.Host,
			db.Port,
			db.Name,
		)
		Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else {
		panic(errors.New(fmt.Sprintf("not support database %s", dbconf.Type)))
	}

	if err != nil {
		panic("failed to connect database")
	}
	Database.AutoMigrate(
		&model.User{},
		&model.Otp{},
		&model.Notebook{},
		&model.Note{})
}
