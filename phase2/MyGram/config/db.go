package config

import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbname   = "mygram"
	db       *gorm.DB
	err      error
   )
func InitDatabase(){
	// dsn := "host=localhost user=postgres password=postgres dbname=mygram port=5432"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }

	// db.AutoMigrate(&models.User{})
	// return db
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
 	dsn := config
 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

 	if err != nil {
  	log.Fatal("errror connecting to database :", err)
 	}

 	fmt.Println("sukses koneksi ke database")
 	db.AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})

	
}

func GetDB() *gorm.DB {
	return db
   }


