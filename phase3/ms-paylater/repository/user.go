package repository

import "ms-paylater/model"

// for register user handler
func (r Repository) CreateUser(name, email, password string) (*model.Users,error) {
	user := model.Users{
		Name: name,
		Email: email,
		Password: password,
	}
	result := r.DB.Create(&user)
	if result.Error != nil{
		return nil,result.Error
	}
	return &user,nil
}

// for login user handler
func (r Repository) GetUserByEmail(email string) (*model.Users,error) {
	var user model.Users
	result := r.DB.Where("email = ?",email).First(&user)
	if result.Error != nil{
		return nil,result.Error
	}
	return &user,nil
}

// for authentication user
func (r Repository) GetUserById(id int) (*model.Users,error) {
	var user model.Users
	result := r.DB.First(&user,id)
	if result.Error != nil{
		return nil,result.Error
	}
	return &user,nil
}