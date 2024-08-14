package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HostelAllotment struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Hostel  primitive.ObjectID `json:"hostel" bson:"hostel" binding:"required"`
	User    primitive.ObjectID `json:"user" bson:"user" binding:"required"`
	Session primitive.ObjectID `json:"session" bson:"session" binding:"required"`
}
