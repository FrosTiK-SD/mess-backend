package models

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Role constants.Role     `json:"role" bson:"role"`

	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName" bson:"middleName"`
	LastName   string `json:"lastName" bson:"lastName"`

	//Contact Details
	Email  string `json:"email" bson:"email" binding:"required"`
	Mobile string `json:"mobile" bson:"mobile"`

	//Institute Details
	StartYear  int    `json:"startYear,omitempty" bson:"startYear,omitempty"`
	EndYear    int    `json:"endYear,omitempty" bson:"endYear,omitempty"`
	RollNo     int    `json:"rollNo,omitempty" bson:"rollNo,omitempty"`
	Department string `json:"department,omitempty" bson:"department,omitempty"`
	Course     string `json:"course,omitempty" bson:"course,omitempty"`
}
