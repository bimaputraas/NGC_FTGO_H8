package main

import (
	"log"
	"net/http"
	"ngc2-webserver/cli"
	"ngc2-webserver/config"
	"ngc2-webserver/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// init db mysql
	db := config.InitDatabase()
	defer db.Close()

	// assign new handler for user
	handler := handler.MarvelHandler{Handler: db}

	// init router
	router := httprouter.New()

	// HEROES AND VILLAINS
	router.GET("/heroes",handler.HandleHeroes)
	router.GET("/villains",handler.HandleVillains)
	
	// INVENTORY
	// task 1 get inventories
	router.GET("/inventories",handler.HandleInventories)
	// task 2 get inventory by param id
	router.GET("/inventories/:id",handler.HandleInventoryById)
	// task 3 post inventories
	router.POST("/inventories",handler.HandleCreateInventories)
	// task 4 PUT /inventories/:id
	router.PUT("/inventories/:id",handler.HandleEditInventory)
	// task 5 DELETE /inventories/:id
	router.DELETE("/inventories/:id",handler.HandleDeleteInventory)

	// CRIMINAL REPORT
	// read
	router.GET("/criminal_reports",handler.HandleViewReports)
	// create
	router.POST("/criminal_reports",handler.HandleCreateReports)
	// edit
	router.PUT("/criminal_reports/:id",handler.HandleEditReport)
	// Delete
	router.DELETE("/criminal_reports/:id",handler.HandleDeleteReport)


	// run server
	cli.StartServerCLI("8383")
	server := http.Server{
		Addr: "localhost:8383",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// Name: OpenAPI (Swagger) Editor
// Id: 42Crunch.vscode-openapi
// Description: OpenAPI extension for Visual Studio Code
// Version: 4.18.6
// Publisher: 42Crunch
// VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi

