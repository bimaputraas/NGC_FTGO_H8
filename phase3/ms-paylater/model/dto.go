package model

type ReqBodyUserRegister struct {
	Name     string `json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type ReqBodyUserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}