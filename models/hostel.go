package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hostel struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name" binding:"required"`
}
