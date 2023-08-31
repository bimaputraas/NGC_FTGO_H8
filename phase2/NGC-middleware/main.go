package main

import (
	"fmt"
	"net/http"
	"ngc5-p2/config"
	"ngc5-p2/handler"
	"ngc5-p2/middleware"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// db
	db := config.InitDatabase()
	defer db.Close()
	
    // init handler
	userHandler := handler.MemberHandler{HandlerDB: db}

	// init router
	router := httprouter.New()

	// before login
	// register
	router.POST("/register",userHandler.Register)
	// login
	router.POST("/login",userHandler.Login)

	// after login
	productHandler := handler.ProductHandler{}
	

	// init server then listen and serve
	server := http.Server{
		Addr: "localhost:9090",
		Handler: middleware.LogMiddleware(router),
	}
	fmt.Println("Server is running on localhost port 9090")
	server.ListenAndServe()
}

