package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Roles []string           `json:"roles" bson:"roles"`
}
