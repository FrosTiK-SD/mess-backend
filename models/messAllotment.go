package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MessAllotment struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Mess     primitive.ObjectID `json:"mess" bson:"mess" binding:"required"`
	User     primitive.ObjectID `json:"user" bson:"user" binding:"required"`
	Semester primitive.ObjectID `json:"semester" bson:"semester" binding:"required"`
}
