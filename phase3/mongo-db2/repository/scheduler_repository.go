package repository

import (
	"context"
	"ngc2_p3/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r Repository) CountEmployees() (int,error){
	// ctx
	ctx,cancel := context.WithTimeout(context.Background(),time.Second*10)
	defer cancel()

	// find (many)
	coll := r.DB.Collection("employees")
	filter := bson.D{}

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return 0,err
	}

	var employees []model.Employees

	for cursor.Next(ctx) {
		var employee model.Employees
		cursor.Decode(&employee)
		if err != nil {
			return 0,err
		}
		employees = append(employees, employee)
	}

	
	return len(employees),nil
}