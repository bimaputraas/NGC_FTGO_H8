package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{}

// view
func (p ProductHandler) View(c *gin.Context) {
	c.JSON(http.StatusAccepted,gin.H{
		"Message" : "Masuk view products",
	})
}

// view by id
func (p ProductHandler) ViewbyId(c *gin.Context) {
	c.JSON(http.StatusAccepted,gin.H{
		"Message" : "Masuk view product by id",
	})
}

// create
func (p ProductHandler) Create(c *gin.Context) {
	c.JSON(http.StatusAccepted,gin.H{
		"Message" : "Masuk create product",
	})
}

// update
func (p ProductHandler) Update(c *gin.Context) {
	c.JSON(http.StatusAccepted,gin.H{
		"Message" : "Masuk update product",
	})
}

// delete
func (p ProductHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusAccepted,gin.H{
		"Message" : "Masuk delete product",
	})
}