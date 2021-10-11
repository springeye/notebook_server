package model

type User struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Username string
	Password string `json:"-"`
	Salt     string
	OptId    *uint
	Opt      *Otp `json:"-"`
}
