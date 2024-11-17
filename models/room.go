package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Hostel    primitive.ObjectID `json:"hostel" bson:"hostel" binding:"required"`
	Number    int                `json:"number" bson:"number" binding:"required"`
	Floor     int                `json:"floor" bson:"floor" binding:"required"`
	Occupancy int                `json:"occupancy" bson:"occupancy" binding:"required"`
	Available bool               `json:"available" bson:"available" binding:"required"`
	Note      string             `json:"note,omitempty" bson:"note,omitempty"`
}
