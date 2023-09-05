package main

import (
	"mygram/config"
	"mygram/handler"
	"mygram/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	// init db
	config.InitDatabase()
	db := config.GetDB()

	// init handler
	userHandler := handler.UserHandler{
		Repository: repository.UserRepository{
			DB: db,
		},
	}

	// setup router and server
  	r := gin.Default()
  	
	// /register and login
	r.POST("/register", userHandler.Register)
	r.POST("/login", func(c *gin.Context) {
	   	c.JSON(http.StatusOK, gin.H{
	     "message": "login",
	   	})
	 })
	 r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}