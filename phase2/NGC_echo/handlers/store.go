package handlers

import (
	"ngc_echo/helpers"
	"ngc_echo/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type StoreHandler struct {
	DB *gorm.DB
}

func (h StoreHandler) View(c echo.Context) error {
	user := c.Get("user").(models.Users)

	var stores []models.Stores
	result := h.DB.Preload("StoreDetails").Find(&stores)
	if result.Error != nil {
		helpers.WriteResponse(c, 500, result.Error.Error())
		return nil
	}
	if result.RowsAffected == 0 {
		helpers.WriteResponse(c, 500, "Data does not exist")
		return nil
	}

	for i:=0;i<len(stores);i++{
		stores[i].StoreDetails.Weather = GetWeather(stores[i].Address)
	}

	helpers.WriteResponseWithData(c, 200, "logged in user "+user.Username, stores)
	return nil
}


func (h StoreHandler) ViewById(c echo.Context) error {
	user := c.Get("user").(models.Users)

	idStr := c.Param("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		helpers.WriteResponse(c, 400, "Invalid param path")
		return nil
	}

	
	var store models.Stores
	result := h.DB.Preload("StoreDetails").Find(&store,id)
	if result.Error != nil {
		helpers.WriteResponse(c, 500, result.Error.Error())
		return nil
	}

	store.StoreDetails.Weather = GetWeather(store.Address)
	
	helpers.WriteResponseWithData(c, 200, "logged in user "+user.Username, store)
	return nil
}