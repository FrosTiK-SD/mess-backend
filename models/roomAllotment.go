package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomAllotment struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Room    primitive.ObjectID `json:"room" bson:"room" binding:"required"`
	User    primitive.ObjectID `json:"user" bson:"user" binding:"required"`
	Session primitive.ObjectID `json:"session" bson:"session" binding:"required"`
}
