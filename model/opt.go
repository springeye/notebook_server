package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Otp struct {
	ID     string `gorm:"primaryKey"`
	UserID string
	User   User
	Key    string
	Words  string
}

func (nb *Otp) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	nb.ID = uuid.NewString()
	return
}
