package dto

type ReqBodyAddEmployee struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address"`
}

type ReqBodyUpdateEmployee struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}