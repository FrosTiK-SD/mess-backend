package interfaces

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserFilter struct {
	StartYear  []int                  `json:"startYear" binding:"required"`
	EndYear    []int                  `json:"endYear" binding:"required"`
	RollNos    []int                  `json:"rollNos" binding:"required"`
	Department []constants.Department `json:"department" binding:"required"`
	Courses    []constants.Course     `json:"courses" binding:"required"`
}

type UserMealForADay struct {
	Date  primitive.DateTime `json:"date"`
	Meals []MealMini         `json:"meals"`
}

type MealMini struct {
	ID     primitive.ObjectID `json:"_id"`
	Name   string             `json:"name"`
	Status string             `json:"status"`
}

const (
	ACTIVE    string = "ACTIVE"
	CANCELLED string = "CANCELLED"
)
