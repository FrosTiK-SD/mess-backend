package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type HostelMessManageToUser struct {
	UserID    primitive.ObjectID   `json:"userID"`
	HostelIDs []primitive.ObjectID `json:"hostelIDs"`
	MessIDs   []primitive.ObjectID `json:"messIDs"`
}
