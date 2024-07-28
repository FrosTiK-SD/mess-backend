package interfaces

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPopulated struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Groups []models.Group     `json:"groups" bson:"groups"` // populated
	Roles  []constants.Role   `json:"roles" bson:"roles"`

	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName" bson:"middleName"`
	LastName   string `json:"lastName" bson:"lastName"`

	// Not nil objectID in case of student
	AllocationDetails models.AllocationDetails `json:"allocationDetails" bson:"allocationDetails"`
	InstituteProfile  models.InstituteProfile  `json:"instituteProfile" bson:"instituteProfile"`

	//Not empty array in case of caretaker
	ManagingDetails models.ManagingDetails `json:"managingDetails" bson:"managingDetails"`

	//Contact Details
	Email  string `json:"email" bson:"email" binding:"required"`
	Mobile string `json:"mobile" bson:"mobile"`
}

type UserFilter struct {
	StartYear  []int                  `json:"startYear" binding:"required"`
	EndYear    []int                  `json:"endYear" binding:"required"`
	RollNos    []int                  `json:"rollNos" binding:"required"`
	Department []constants.Department `json:"department" binding:"required"`
	Courses    []constants.Course     `json:"courses" binding:"required"`
}
