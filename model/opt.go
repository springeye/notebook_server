package model

type Otp struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserId uint
	User   User
	Key    string
	Words  string
}
