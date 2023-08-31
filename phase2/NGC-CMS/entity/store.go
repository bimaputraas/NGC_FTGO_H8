package entity

type Store struct {
	ID        int    `gorm:"primaryKey,autoIncrement" json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	StoreType string `json:"store_type"`
}