package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID   `json:"_id,omitempty" bson:"_id"`
	Groups []primitive.ObjectID `json:"groups" bson:"groups"`

	// Not nil objectID in case of student
	AllocationDetails AllocationDetails `json:"allocationDetails,omitempty" bson:"allocationDetails,omitempty"`
	InstituteProfile  InstituteProfile  `json:"instituteProfile,omitempty" bson:"instituteProfile,omitempty"`

	//Not empty array in case of caretaker
	ManagingDetails ManagingDetails `json:"managingDetails,omitempty" bson:"managingDetails,omitempty"`

	//Contact Details
	Email  string `json:"email" bson:"email" binding:"required"`
	Mobile string `json:"mobile" bson:"mobile"`
}

type InstituteProfile struct {
	StartYear  int    `json:"startYear,omitempty" bson:"startYear,omitempty"`
	EndYear    int    `json:"endYear,omitempty" bson:"endYear,omitempty"`
	RollNo     int    `json:"rollNo,omitempty" bson:"rollNo,omitempty"`
	Department string `json:"department,omitempty" bson:"department,omitempty"`
	Course     string `json:"course,omitempty" bson:"course,omitempty"`
}

type AllocationDetails struct {
	AllocatedHostel primitive.ObjectID `json:"allocatedHostel" bson:"allocatedHostel,omitempty"`
	AllocatedMess   primitive.ObjectID `json:"allocatedMess,omitempty" bson:"allocatedMess,omitempty"`
	AllocatedRoom   primitive.ObjectID `json:"allocatedRoom,omitempty" bson:"allocatedRoom,omitempty"`
}

type ManagingDetails struct {
	ManagingHostels []primitive.ObjectID `json:"managingHostels,omitempty" bson:"managingHostels,omitempty"`
	ManagingMesses  []primitive.ObjectID `json:"managingMesses,omitempty" bson:"managingMesses,omitempty"`
}
