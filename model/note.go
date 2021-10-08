package model

import (
	"time"
)

type Note struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Uuid        string
	UserId      uint
	User        User
	Title       string
	NotebookId  uint
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
