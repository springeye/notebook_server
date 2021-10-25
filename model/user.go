package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	Username string
	Password string `json:"-"`
	Salt     string
	OptID    *string
	Opt      *Otp `json:"-"`
}

func (nb *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	nb.ID = uuid.NewString()
	return
}
