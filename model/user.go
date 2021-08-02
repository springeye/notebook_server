package model

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string
	Salt     string
	OptId    *uint
	Opt      *Otp
}
