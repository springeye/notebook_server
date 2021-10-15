package database

import (
	"errors"
	"fmt"
	log2 "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	conf "notebook/config"
	"notebook/model"
	"os"
	"time"
)

var Database *gorm.DB

func init() {
	dbLogger := log2.WithFields(log2.Fields{})
	var err error
	logLevel := logger.Error
	if conf.Conf.Database.Logger.Level == "info" {
		logLevel = logger.Info
	} else if conf.Conf.Database.Logger.Level == "warn" {
		logLevel = logger.Warn
	} else if conf.Conf.Database.Logger.Level == "error" {
		logLevel = logger.Error
	} else if conf.Conf.Database.Logger.Level == "silent" {
		logLevel = logger.Silent
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

	dbconf := conf.Conf.Database
	dbLogger.Debugf("Database type: %s", dbconf.Type)
	if dbconf.Type == conf.SQLITE {
		dbLogger.Debugf("Database conf: %v", dbconf.Sqlite)

		Database, err = gorm.Open(sqlite.Open(dbconf.Sqlite.File), &config)
	} else if dbconf.Type == conf.MYSQL {
		dbLogger.Debugf("Database conf: %v", dbconf.MySQL)
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
