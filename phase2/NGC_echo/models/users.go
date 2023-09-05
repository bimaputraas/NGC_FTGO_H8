package models

type Users struct {
	ID            uint    `gorm:"primaryKey,autoIncrement"`
	Username      string  `json:"username" gorm:"unique"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
}
