package interfaces

import "go.mongodb.org/mongo-driver/bson/primitive"

type AssignHostelToUsersRequestBody struct {
	Hostel primitive.ObjectID   `json:"hostel" binding:"required"`
	Users  []primitive.ObjectID `json:"users" binding:"required"`
}

type AssignMessToUsersRequestBody struct {
	Mess  primitive.ObjectID   `json:"mess" binding:"required"`
	Users []primitive.ObjectID `json:"users" binding:"required"`
}

type BatchCreateHostelRoomsRequest struct {
	RangeStart int  `json:"rangeStart" binding:"required"`
	RangeEnd   int  `json:"rangeEnd" binding:"required"`
	Floor      int  `json:"floor" binding:"required"`
	Occupancy  int  `json:"occupancy" binding:"required"`
	Available  bool `json:"available" binding:"required"`
}
