package models

type Stores struct {
	ID           uint         `gorm:"primaryKey"`
	Name         string       `json:"name"`
	Address      string       `json:"address"`
	StoreDetails StoreDetails `json:"store_details" gorm:"foreignKey:StoreID"`
}

type StoreDetails struct {
	ID         uint    `gorm:"primaryKey"`
	StoreID    uint    `json:"store_id"`
	Weather    string  `json:"weather"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	TotalSales float64 `json:"total_sales"`
	Rating     float64 `json:"rating"`
}