package models

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID   `json:"_id" bson:"_id"`
	Groups []primitive.ObjectID `json:"groups" bson:"groups"`
	Roles  []constants.Role     `json:"roles" bson:"roles"`

	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName" bson:"middleName"`
	LastName   string `json:"lastName" bson:"lastName"`

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
	StartYear  int    `json:"startYear" bson:"startYear"`
	EndYear    int    `json:"endYear" bson:"endYear"`
	RollNo     int    `json:"rollNo" bson:"rollNo"`
	Department string `json:"department" bson:"department"`
	Course     string `json:"course" bson:"course"`
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
