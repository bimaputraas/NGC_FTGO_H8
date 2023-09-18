package main

import (
	"ms-paylater/config"
	"ms-paylater/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.InitDB()
	e := router.InitRouter(db)

	e.Logger.Fatal(e.Start(":8080"))
}