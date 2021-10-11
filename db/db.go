package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"notebook/model"
	"os"
	"time"
)

var Database *gorm.DB

func init() {
	var err error
	Database, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Warn, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}
	Database.AutoMigrate(
		&model.User{},
		&model.Otp{},
		&model.Notebook{},
		&model.Note{})
}
