package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notebook struct {
	ID       string `gorm:"primaryKey"`
	UserID   string
	User     User
	Title    string
	PID      *string
	Notebook *Notebook `gorm:"foreignKey:pid"`
	Notes    []Note
	Password string
}

func (nb *Notebook) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	nb.ID = uuid.NewString()
	return
}
