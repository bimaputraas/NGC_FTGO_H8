package repository

import (
	"errors"
	"ngc-cms/entity"

	"gorm.io/gorm"
)

type UserQuery struct {
	HandlerDB *gorm.DB
}

// register
func (h UserQuery) Insert(store entity.Store) (entity.Store,error){	
	result := h.HandlerDB.Create(&store)
	// if error
	if result.Error != nil {
		panic(result.Error)
	}

	// if error insert
	if result.Error != nil{
		return store,result.Error
	}
	return store,nil
}

// find all
func (h UserQuery) FindAll() (stores []entity.Store,err error){
	result := h.HandlerDB.Find(&stores)

	// if error
	if result.Error != nil {
		panic(result.Error)
	}

	// if data 0
	if result.RowsAffected == 0{
		return stores, errors.New("No data")
	} 
	return stores,nil
}

// find by id
func (h UserQuery) FindbyId(id int) (store entity.Store,err error){
	result := h.HandlerDB.First(&store,id)

	// if error
	if result.Error != nil {
		panic(result.Error)
	}

	// if not found
	if result.RowsAffected == 0{
		return store, errors.New("No data")
	} 

	return store,nil
}

// find by email
func (h UserQuery) FindbyEmail(email string) (store entity.Store,err error){
	result:= h.HandlerDB.Where("email = ?", email).Find(&store)

	// if error
	if result.Error != nil {
		panic(result.Error)
	}

	// if not found
	if result.RowsAffected == 0 {
		return store,errors.New("No data")
	}
	return store,nil
}




