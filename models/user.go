package models

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID     `json:"_id,omitempty" bson:"_id"`
	Permissions     []constants.Permission `json:"permissions" bson:"permissions" binding:"required"`
	Groups          []primitive.ObjectID   `json:"groups" bson:"groups"` //Object ID of groups to which the user belongs
	AllocatedHostel primitive.ObjectID     `json:"allocatedHostel" bson:"allocatedHostel,omitempty"`
	AllocatedMess   primitive.ObjectID     `json:"allocatedMess,omitempty" bson:"allocatedMess,omitempty"`
	AllocatedRoom   primitive.ObjectID     `json:"allocatedRoom,omitempty" bson:"allocatedRoom,omitempty"`

	//Academic Details
	StartYear  uint   `json:"startYear,omitempty" bson:"startYear,omitempty"`
	EndYear    uint   `json:"endYear,omitempty" bson:"endYear,omitempty"`
	RollNo     string `json:"rollNo,omitempty" bson:"rollNo,omitempty"`
	Department string `json:"department,omitempty" bson:"department,omitempty"`
	Course     string `json:"course,omitempty" bson:"course,omitempty"`

	//Managing Details
	ManagingHostels []primitive.ObjectID
	ManagingMesses  []primitive.ObjectID

	//Contact Details
	Email  string `json:"email" bson:"email" binding:"required"`
	Mobile string `json:"mobile" bson:"mobile"`
}
