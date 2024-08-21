package interfaces

import "github.com/FrosTiK-SD/mess-backend/models"

type PopulatedRoom struct {
	models.Room
	Allotments int `json:"allotments" bson:"allotments"`
}
