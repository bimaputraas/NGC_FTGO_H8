package config

import (
	"ms-paylater/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=ms_paylater_p3 port=5432 sslmode=disable"
  	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.Users{})
	if err != nil {
		panic(err)
	}

	return db
}
  