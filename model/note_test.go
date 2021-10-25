package model

import (
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"time"
)

func TestCreateNote(t *testing.T) {
	config := gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		),
	}

	db, err := gorm.Open(sqlite.Open("file::memory:?store=shared"), &config)
	db.AutoMigrate(&Note{})
	assert.Equal(t, err, nil)
	tx := db.Begin()
	note := Note{Title: "test", Content: "test"}
	tx.Create(&note)
	t.Logf("create note result id is %s\n", note.ID)
	assert.NotEqual(t, note.ID, "")
	tx.Rollback()
}
