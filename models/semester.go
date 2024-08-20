package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Semester struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name" binding:"required"`
	Start primitive.DateTime `json:"start" bson:"start" binding:"required"`
	End   primitive.DateTime `json:"end" bson:"end" binding:"required"`
}
