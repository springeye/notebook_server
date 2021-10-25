package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type Notebook struct {
	ID       string `gorm:"primaryKey"`
	UserID   string
	Title    string
	PID      *string
	Notes    []Note
	Password string
}

func (nb *Notebook) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	nb.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return
}
