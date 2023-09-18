package model

type Users struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password string
}