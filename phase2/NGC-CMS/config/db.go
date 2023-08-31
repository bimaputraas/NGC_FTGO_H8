package config

import (
	"log"
	"ngc-cms/entity"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	// dataSource := os.Getenv("DBSOURCENAME")
	dsn := "host=localhost user=postgres password=postgres dbname=ngc_mcs_p2 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&entity.Store{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}