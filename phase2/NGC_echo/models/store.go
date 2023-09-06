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
	Weather    Weather `json:"weather" gorm:"foreignKey:StoreDetailsId"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	TotalSales float64 `json:"total_sales"`
	Rating     float64 `json:"rating"`
}

type Weather struct {
	StoreDetailsId uint    `json:"store_details_id"`
	CloudPct       float64 `json:"cloud_pct"`
	Temp           float64 `json:"temp"`
	FeelsLike      float64 `json:"feels_like"`
	Humidity       float64 `json:"humidity"`
	MinTemp        float64 `json:"min_temp"`
	MaxTemp        float64 `json:"max_temp"`
	WindSpeed      float64 `json:"wind_speed"`
	WindDegrees    float64 `json:"wind_degrees"`
	Sunrise        float64 `json:"sunrise"`
	Sunset         float64 `json:"sunset"`
}