package main

import (
	"net/http"
	"ngc5-p2/config"
	"ngc5-p2/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// db
	db := config.InitDatabase()
	defer db.Close()
	
    // init handler
	handler := handler.Handler{HandlerDB: db}

	// init router
	router := httprouter.New()
	// register
	router.POST("/register",handler.Register)
	// login
	router.POST("/register",handler.Login)

	// init server then listen and serve
	server := http.Server{
		Addr: "localhost:9090",
		Handler: router,
	}
	server.ListenAndServe()
}