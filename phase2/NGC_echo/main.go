package main

import (
	"ngc_echo/config"
	"ngc_echo/handlers"
	"ngc_echo/helpers"
	"ngc_echo/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	// exec godotenv
	helpers.LoadEnv()

	// init db
	db := config.InitDatabase()

	// init handler
	userHandler := handlers.UsersHandler{
		DB: db,
	}
	productHandler := handlers.ProductHandler{
		DB: db,
	}
	transactionHandler := handlers.TransactionsHandler{
		DB: db,
	}

	// routing
	e := echo.New()

	// users
	usersGroup := e.Group("/users")
	usersGroup.POST("/register",userHandler.Register)
	usersGroup.POST("/login",userHandler.Login)

	// products
	e.GET("/products",middleware.Authentication(productHandler.View))

	// transaction
	e.POST("/transactions",middleware.Authentication(transactionHandler.Buy))
	
    e.Logger.Fatal(e.Start(":8080"))
}