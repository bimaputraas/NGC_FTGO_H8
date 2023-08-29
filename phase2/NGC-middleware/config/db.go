package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// data source
var(
	dbUser = "root"
	dbPassword = ""
	dbHost = "localhost:3306"
	dbName = "ngc5_p2"
)

// function init database
func InitDatabase() *sql.DB{
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",dbUser,dbPassword,dbHost,dbName) 
	db,err := sql.Open("mysql",dataSource)
	if err != nil {
		log.Fatal(err)
	}
	return db
}