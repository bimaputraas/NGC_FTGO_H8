package main

import (
	"ngc-cms/config"
	"ngc-cms/handler"
	"ngc-cms/repository"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/override/docs"
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
  	
	// swagger and router
	v1 := r.Group("/v1")
	{
		usersGroup := v1.Group("/users")
		{
			usersGroup.GET("/",userHandler.ViewAll)
			usersGroup.GET("/:id",userHandler.View)
			usersGroup.POST("/register",userHandler.Register)
			usersGroup.POST("/login",userHandler.Login)
			// usersGroup.POST("/login",userHandler.Login)
		}
	}
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	
  	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}