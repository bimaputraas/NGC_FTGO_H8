package main

import (
	"ngc2_p3/config"
	"ngc2_p3/router"

	_ "github.com/joho/godotenv/autoload"
)


func main() {
	// db
	db := config.NewMongoDB("ngc2_p3")

	// init echo
    e := router.NewEcho(db)

	// start echo
    e.Logger.Fatal(e.Start(":8080"))
}
