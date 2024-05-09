package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"` //for uniquely identifying a room
	Name      string             `json:"name" bson:"name" binding:"required"`
	Available bool               `json:"available" bson:"available" binding:"required"`
	Remarks   string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
}
