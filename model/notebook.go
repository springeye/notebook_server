package model

type Notebook struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserId   uint
	User     User
	Title    string
	Pid      *uint
	Notebook *Notebook `gorm:"foreignKey:pid"`
	Notes    []Note
	Password string
}
