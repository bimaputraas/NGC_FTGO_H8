package main

import (
	"log"
	"net/http"
	"ngc2-webserver/config"
	"ngc2-webserver/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// db mysql
	db := config.InitDatabase()
	defer db.Close()

	// assign new handler for user
	handler := handler.MarvelHandler{Handler: db}

	// router
	router := httprouter.New()

	// get heroes and villains
	router.GET("/heroes",handler.HandlerHeroes)
	router.GET("/villains",handler.HandlerVillains)

	// task 1 get inventories
	router.GET("/inventories",handler.HandlerInventories)

	// task 2 get inventory by param id
	router.GET("/inventories/:id",handler.HandlerInventoryById)

	// task 3 post inventories
	router.POST("/inventories",handler.HandlerCreateInventories)
	
	// task 4 PUT /inventories/:id
	router.PUT("/inventories/:id",handler.HandlerEditInventory)

	// task 5 DELETE /inventories/:id
	router.DELETE("/inventories/:id",handler.HandlerDeleteInventory)



	// run server
	server := http.Server{
		Addr: "localhost:8383",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}



