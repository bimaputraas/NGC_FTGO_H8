package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employees struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name    string             `bson:"name,omitempty" json:"name"`
	Age     int                `bson:"age,omitempty" json:"age"`
	Email   string             `bson:"email,omitempty" json:"email"`
	Address string             `bson:"address,omitempty" json:"address"`
}