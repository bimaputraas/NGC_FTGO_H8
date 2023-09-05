package repository

import (
	"mygram/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r UserRepository) Insert(user models.User) (models.User,error){
	result := r.DB.Create(&user)
	
	// if fail query
	if result.Error != nil{
		return user, result.Error
	}
	
	return user,nil
}