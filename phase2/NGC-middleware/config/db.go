package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// function init database
func InitDatabase() (*sql.DB){
	dataSource := os.Getenv("DBSOURCE")
	db,err := sql.Open("mysql",dataSource)
	if err != nil {
		panic(err)
	}
	return db
}