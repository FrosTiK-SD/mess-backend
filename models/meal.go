package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Mess     primitive.ObjectID `json:"mess" bson:"mess"`
	MealType primitive.ObjectID `json:"mealType" bson:"mealType"`

	Date primitive.DateTime `json:"date" bson:"date"`
	Day  string             `json:"day" bson:"day"`

	Menu              []primitive.ObjectID `json:"menu" bson:"menu"`
	CancelledStudents []primitive.ObjectID `json:"cancelledStudents" bson:"cancelledStudents"`
	AttendedStudents  []primitive.ObjectID `json:"attendedStudents" bson:"attendedStudents"`
}

type MealType struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Mess primitive.ObjectID `json:"mess,omitempty" bson:"mess"`

	Name      string `json:"name,omitempty" bson:"name"`
	StartTime string `json:"startTime,omitempty" bson:"endTime"` // time.Kitchen format
	EndTime   string `json:"endTime,omitempty" bson:"endTime"`   // time.Kitchen format
	Cost      int    `json:"cost,omitempty" bson:"cost"`         // in paisa
}

type MenuItem struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Mess primitive.ObjectID `json:"mess,omitempty" bson:"mess"`

	Name   string  `json:"name,omitempty" bson:"name,omitempty"`
	ImgUrl *string `json:"imgUrl,omitempty" bson:"imgUrl,omitempty"`
	Cost   int     `json:"cost,omitempty" bson:"cost,omitempty"` // in paisa
}
