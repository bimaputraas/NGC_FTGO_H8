package models

type Transactions struct {
	ID          uint    `gorm:"primaryKey,autoIncrement"`
	UserID      uint    `json:"user_id" binding:"required"`
	ProductID   uint    `json:"product_id" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	TotalAmount float64 `json:"total_amount"`
}
