package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups []primitive.ObjectID `json:"groups" bson:"groups"`

	// Not nil objectID in case of student
	AllocationDetails AllocationDetails `json:"allocationDetails" bson:"allocationDetails"`
	InstituteProfile  InstituteProfile  `json:"instituteProfile" bson:"instituteProfile"`

	//Not empty array in case of caretaker
	ManagingDetails ManagingDetails `json:"managingDetails" bson:"managingDetails"`

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
	Hostel primitive.ObjectID `json:"hostel" bson:"hostel"`
	Mess   primitive.ObjectID `json:"mess" bson:"mess"`
	Room   primitive.ObjectID `json:"room" bson:"room"`
}

type ManagingDetails struct {
	Hostels []primitive.ObjectID `json:"hostels" bson:"hostels"`
	Messes  []primitive.ObjectID `json:"messes" bson:"messes"`
}

type UserMini struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
}

type StudentMini struct {
	UserMini
	InstituteProfile InstituteProfile `json:"instituteProfile"`
}
