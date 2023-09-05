package handlers

import (
	"ngc_echo/helpers"
	"ngc_echo/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (h ProductHandler) View(c echo.Context) error {
	user := c.Get("user").(models.Users)
	
	var products []models.Products
	result := h.DB.Find(&products)
	if result.Error != nil{
		helpers.WriteResponse(c,500,result.Error.Error())
		return nil
	}
	if result.RowsAffected == 0{
		helpers.WriteResponse(c,500,"Data does not exist")
		return nil
	}

	helpers.WriteResponseWithData(c, 200, "logged in user "+user.Username, products)
	return nil
}