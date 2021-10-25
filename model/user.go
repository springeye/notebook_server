package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Password string `json:"-"`
	Salt     string
	Key      *string
	Words    *string
}

func (nb *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	nb.ID = strings.ReplaceAll(uuid.NewString(), "-", "")
	return
}
