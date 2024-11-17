package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HostelAllotment struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Hostel   primitive.ObjectID `json:"hostel" bson:"hostel" binding:"required"`
	User     primitive.ObjectID `json:"user" bson:"user" binding:"required"`
	Semester primitive.ObjectID `json:"semester" bson:"semester" binding:"required"`
}
