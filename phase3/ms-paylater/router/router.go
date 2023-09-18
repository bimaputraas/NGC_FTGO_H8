package router

import (
	"ms-paylater/handler"
	"ms-paylater/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *echo.Echo{
	// init echo
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

	// init handler
	handler := handler.InitHandler(db)

	// init middleware auth
	authMiddleware := middleware.InitAuth(db)

	// routing
	v1 := e.Group("/v1/ms-paylater")
	{
		v1.POST("/register", handler.RegisterUser)
		v1.POST("/login", handler.LoginUser)
		v1.POST("/loan", handler.AddLoan,authMiddleware.Authentication)
		v1.GET("/limit", handler.ViewLimit,authMiddleware.Authentication)
		v1.POST("/withdraw", handler.Withdraw,authMiddleware.Authentication)
		v1.POST("/pay", handler.Pay,authMiddleware.Authentication)
	}

	return e
}