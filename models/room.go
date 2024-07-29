package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Room struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Hostel    primitive.ObjectID `json:"hostel" bson:"hostel" binding:"required"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Floor     int                `json:"floor" bson:"floor,omitempty"`
	Available bool               `json:"available" bson:"available" binding:"required"`
	Remarks   string             `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

type PopulatedRoom struct {
	Room                      // Embedding Room struct
	AllocatedTo []StudentMini `json:"allocatedTo"`
}
