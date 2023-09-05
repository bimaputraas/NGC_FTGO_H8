package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
  
func InitDatabase() *gorm.DB{
	dsn := "host=localhost user=postgres password=postgres dbname=ngc_echo port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to DB")
	}
	
	// db.AutoMigrate(&models.Users{})
	return db
}

//migration
//migrate -database "postgres://postgres:postgres@localhost:5432/ngc_echo?sslmode=disable" -path migrations up
//migrate -database "postgres://postgres:postgres@localhost:5432/ngc_echo?sslmode=disable" -path migrations down
//migrate -database "postgres://postgres:postgres@localhost:5432/ngc_echo?sslmode=disable" -path migrations force 20230905124633