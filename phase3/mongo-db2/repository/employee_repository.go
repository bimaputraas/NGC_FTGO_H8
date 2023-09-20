package repository

import (
	"context"
	"ngc2_p3/dto"
	"ngc2_p3/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Repository) InsertEmployee(reqBody dto.ReqBodyAddEmployee) (*model.Employees,error){
	// ctx
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	// doc
	doc := model.Employees{
		Name: reqBody.Name,
		Age: reqBody.Age,
		Email: reqBody.Email,
		Address: reqBody.Address,
	}

	// insert
	coll := r.DB.Collection("employees")
	result,err := coll.InsertOne(ctx,doc)
	doc.ID = result.InsertedID.(primitive.ObjectID)
	if err != nil {
		return nil, err
	}

	return &doc,nil
}

func (r Repository) FindEmployees() ([]model.Employees,error){
	// ctx
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	// find (many)
	coll := r.DB.Collection("employees")
	filter := bson.D{}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil,err
	}

	var employees []model.Employees

	for cursor.Next(ctx) {
		var employee model.Employees
		cursor.Decode(&employee)
		if err != nil {
			return nil,err
		}
		employees = append(employees, employee)
	}

	return employees,nil
}

func (r Repository) FindEmployeById(id primitive.ObjectID) (*model.Employees,error){
	// ctx
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	// find (one)
	coll := r.DB.Collection("employees")
	filter := bson.D{
		{Key: "_id",Value: id},
	}
	result := coll.FindOne(ctx, filter)
	var employee model.Employees
	err := result.Decode(&employee)
	if err != nil {
		return nil, err
	}
	
	return &employee,nil
}

func (r Repository) UpdateEmployeeById(id primitive.ObjectID, reqBody dto.ReqBodyUpdateEmployee) error{
	// ctx
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	// find (one)
	coll := r.DB.Collection("employees")
	filter := bson.D{
		{Key: "_id",Value: id},
	}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "name", Value: reqBody.Name},
				{Key: "email", Value: reqBody.Email},
				{Key: "age", Value: reqBody.Age},
				{Key: "address", Value: reqBody.Address},
			},
		},
	}

	result := coll.FindOneAndUpdate(ctx, filter,update)
	if result.Err() != nil {
		return result.Err()
	}
	
	return nil
}

func (r Repository) DeleteEmployeeById(id primitive.ObjectID) (error){
	// ctx
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	// find (one)
	coll := r.DB.Collection("employees")
	filter := bson.D{
		{Key: "_id",Value: id},
	}
	result := coll.FindOneAndDelete(ctx, filter)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}


