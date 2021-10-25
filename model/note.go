package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID          string `gorm:"primaryKey"`
	UserId      string
	User        User
	Title       string
	NotebookId  string
	Notebook    *Notebook
	Password    string
	Content     string
	VersionKey  string
	VersionCode uint64
	Encrypted   string
	Deleted     bool
	CreatedTime time.Time
	UpdateTime  time.Time `gorm:"autoUpdateTime"`
}

func (nb *Note) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	nb.ID = uuid.NewString()
	return
}
