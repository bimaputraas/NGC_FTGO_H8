package main

import (
	"ngc-cms/config"
	"ngc-cms/handler"
	"ngc-cms/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	// db
	db := config.InitDatabase()

	// handler
	userHandler := handler.UserHandler{
		Handler: repository.UserQuery{HandlerDB: db},
	}

	// gin
  	r := gin.Default()
	r.GET("/users",userHandler.ViewAll)
	r.GET("/users/:id",userHandler.View)
	r.POST("/users/register",userHandler.Register)
  	
  	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}