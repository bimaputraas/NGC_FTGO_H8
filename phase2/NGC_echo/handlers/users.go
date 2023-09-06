package handlers

import (
	"ngc_echo/helpers"
	"ngc_echo/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsersHandler struct {
	DB *gorm.DB
}

func (h UsersHandler) Register(c echo.Context) error {
	// bind reqbody
	var reqBody models.Users
	err := c.Bind(&reqBody)
	if err != nil {
		helpers.WriteResponse(c,400,"Failed bind")

		return nil
	}
	
	// hash
	reqBody.Password,err = helpers.HashPassword(reqBody.Password)
	if err != nil {
		helpers.WriteResponse(c,500,"Failed hash")
	}

	// validate
	// err = helpers.Validate(reqBody)
	// if err != nil {
	// 	helpers.WriteResponse(c,400,err.Error())
	// 	return nil
	// }

	// create
	result := h.DB.Create(&reqBody)

	if result.Error != nil{
		helpers.WriteResponse(c,400,result.Error.Error())

		return nil
	}

	helpers.WriteResponseWithData(c,201,"Success create",reqBody)
	return nil
}

func (h UsersHandler) Login(c echo.Context) error {
	// bind reqbody
	var reqBody models.Users
	err := c.Bind(&reqBody)
	if err != nil {
		helpers.WriteResponse(c,400,"Failed bind")
		return nil
	}
	
	// find from db by username
	var user models.Users
	result := h.DB.Where("username = ?",reqBody.Username).Find(&user)
	if result.Error != nil{
		helpers.WriteResponse(c,400,result.Error.Error())
		return nil
	}
	if result.RowsAffected == 0{
		helpers.WriteResponse(c,400,"Data does not exist")
		return nil
	}

	// compare hash
	if !helpers.CheckPasswordHash(reqBody.Password,user.Password){
		helpers.WriteResponse(c,500,"Failed compare hash")
		return nil
	}

	// success
	// generate token
	tokenString,err:= helpers.GenerateToken(int(user.ID))
	if err != nil {
		helpers.WriteResponse(c,500,"Failed generate token")
		return nil
	}

	// create cookie and send token
	// cookie := new(http.Cookie)
    // cookie.Name = "Authorize-Token"
    // cookie.Value = tokenString
    // cookie.Expires = time.Now().Add(24 * time.Hour)
	// cookie.Domain = "http://localhost:8080"
	// c.SetCookie(cookie)
	// http.SetCookie(c.Response(),cookie)
	
	// send to header response
	c.Response().Header().Set("Authorization",tokenString)
	
	// response body
	helpers.WriteResponseWithData(c,200,"Success login",user)
    return nil
}