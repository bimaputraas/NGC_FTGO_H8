package handler

import (
	"ms-paylater/helper"
	"ms-paylater/model"

	"github.com/labstack/echo/v4"
)



func (h Handler) RegisterUser(c echo.Context) error {
	// bind
	var reqBody model.ReqBodyUserRegister
	c.Bind(&reqBody)

	// validate
	helper.ValidateStruct(reqBody)

	// hash
	hash,err := helper.HashPassword(reqBody.Password)
	if err != nil {
		return helper.WriteErrorResponse(500,err.Error())
	}
	reqBody.Password = hash

	// create
	newUser,err := h.Repository.CreateUser(reqBody.Name,reqBody.Email,reqBody.Password)
	if err != nil {
		return helper.WriteErrorResponse(400,err.Error())
	}

	return helper.WriteResponseWithData(c, 201, "Success Register",newUser)
}

func (h Handler) LoginUser(c echo.Context) error {
	// bind
	var reqBody model.ReqBodyUserLogin
	c.Bind(&reqBody)

	// validate
	helper.ValidateStruct(reqBody)

	// get user
	user,err := h.Repository.GetUserByEmail(reqBody.Email)
	if err != nil {
		return helper.WriteErrorResponse(400,"wrong email or password")
	}

	// compare password and hash
	if !helper.CheckPasswordHash(reqBody.Password,user.Password){
		return helper.WriteErrorResponse(400,"wrong email or password")
	}

	// generate token
	tokenString,err := helper.GenerateJWT(int(user.ID))
	if err != nil {
		return helper.WriteErrorResponse(500,err.Error())
	}

	return helper.WriteResponseWithData(c, 201, "Success Login "+user.Email,tokenString)
}

func (h Handler) AddLoan(c echo.Context) error {

	return helper.WriteResponse(c, 201, "Masuk post loan")
}

func (h Handler) Withdraw(c echo.Context) error {

	return helper.WriteResponse(c, 201, "Masuk post loan")
}

func (h Handler) ViewLimit(c echo.Context) error {

	return helper.WriteResponse(c, 201, "Masuk post loan")
}

func (h Handler) Pay(c echo.Context) error {

	return helper.WriteResponse(c, 201, "Masuk post loan")
}