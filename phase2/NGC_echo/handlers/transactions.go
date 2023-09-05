package handlers

import (
	"fmt"
	"ngc_echo/helpers"
	"ngc_echo/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TransactionsHandler struct {
	DB *gorm.DB
}

func (h TransactionsHandler) Buy(c echo.Context) error {
	// init tx for transaction by gorm
	tx := h.DB.Begin()
	
	// bind transaction
	var transaction models.Transactions
	err := c.Bind(&transaction)
	if err != nil {
		helpers.WriteResponse(c, 400,"Failed bind")
		tx.Rollback()
		return nil	
	}
	
	// find product
	var product models.Products
	result := tx.Where("id = ?",transaction.ProductID).First(&product)
	if result.Error != nil {
		helpers.WriteResponse(c, 400,"Data does not exist")
		tx.Rollback()
		return nil
	}
	
	// check if stock unavailable
	if product.Stock < transaction.Quantity{
		helpers.WriteResponse(c, 400,"Stock unavailable")
		tx.Rollback()
		return nil
	}
	
	// if stock available
	product.Stock -= transaction.Quantity
	result = tx.Save(&product)
	if result.Error != nil {
		helpers.WriteResponse(c,500,result.Error.Error())
		tx.Rollback()
		return nil
	}

	// get logged in user from context
	user := c.Get("user").(models.Users)

	// check if user fund less than total amount transaction
	transaction.TotalAmount = product.Price * float64(transaction.Quantity)
	if user.DepositAmount < transaction.TotalAmount{
		helpers.WriteResponse(c,400,"The fund is insufficient")
		tx.Rollback()
		return nil
	}

	// if fund is sufficient
	user.DepositAmount -= transaction.TotalAmount
	result = tx.Save(&user)
	if result.Error != nil{
		helpers.WriteResponse(c,500,result.Error.Error())
		tx.Rollback()
		return nil
	}

	// create transaction information
	transaction.UserID = user.ID
	result = tx.Create(&transaction)
	if result.Error != nil{
		helpers.WriteResponse(c,500,result.Error.Error())
		tx.Rollback()
		return nil
	}

	// success purchase
	tx.Commit()
	message := fmt.Sprintf("%s successfully purchased '%s'", user.Username,product.Name)
	helpers.WriteResponseWithData(c, 201,message, transaction)

	return nil
}