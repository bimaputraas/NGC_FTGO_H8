package entity

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Price       int    `json:"price"`
	StoreId     int    `json:"store_id"`
}