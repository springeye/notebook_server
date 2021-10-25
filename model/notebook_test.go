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

func TestCreateNotebook(t *testing.T) {
	config := gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel:                  logger.Info,
				SlowThreshold:             time.Second, // Slow SQL threshold
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			},
		),
	}
	db, err := gorm.Open(sqlite.Open("file::memory:?store=shared"), &config)
	db.AutoMigrate(&Notebook{}, &Note{})
	assert.Equal(t, err, nil)
	t.Run("create notebook", func(t *testing.T) {
		tx := db.Begin()
		notebook := Notebook{Title: "test"}
		tx.Create(&notebook)
		assert.NotEqual(t, notebook.ID, "")
		tx.Rollback()
	})
	t.Run("create note by notebook", func(t *testing.T) {

		tx := db.Begin()
		notebook := Notebook{Title: "test"}
		tx.Create(&notebook)
		assert.NotEqual(t, notebook.ID, "")
		note := Note{Title: "test", NotebookID: notebook.ID}
		err = tx.Create(&note).Error

		assert.Equal(t, err, nil)
		assert.NotEqual(t, note.ID, "")
		var result Notebook
		err = tx.Preload("Notes").First(&result).Error
		assert.Equal(t, err, nil)
		assert.Equal(t, len(result.Notes), 1)
		tx.Rollback()
	})

}
