package main

import (
	"github.com/labstack/echo/v4"
	"ngc_echo/config"
	"ngc_echo/internal/handlers"
	"ngc_echo/internal/middleware"
	"ngc_echo/internal/repository"
	"ngc_echo/internal/usecase"
	"ngc_echo/pkg"
)

func main() {
	// exec godotenv
	pkg.LoadEnv()

	// init router
	e := echo.New()

	// init db
	db := config.InitDatabase()

	// init repository
	repo := repository.NewRepository(db)

	// init usecase
	uc := usecase.NewUsecase(repo)

	// init handler
	h := handlers.NewHandler(uc)

	// users
	usersGroup := e.Group("/users")
	usersGroup.POST("/register", h.Register)
	usersGroup.POST("/login", h.Login)

	// products
	e.GET("/products", middleware.Authentication(h.ProductsHandler))

	// transaction
	e.POST("/transactions", h.Buy, middleware.Authentication)

	e.Logger.Fatal(e.Start(":8080"))
}
