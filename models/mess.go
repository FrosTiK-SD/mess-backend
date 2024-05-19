package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mess struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Hostel primitive.ObjectID `json:"hostel" bson:"hostel" binding:"required"`
	Name   string             `json:"name" bson:"name"`
}
