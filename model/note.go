package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Note struct {
	ID          string `gorm:"primaryKey"`
	UserId      string
	Title       string
	NotebookID  string
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
	nb.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return
}
