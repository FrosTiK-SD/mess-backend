package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type AssignHostelToUsersRequestBody struct {
	Hostel primitive.ObjectID   `json:"hostel" binding:"required"`
	Users  []primitive.ObjectID `json:"users" binding:"required"`
}
