package interfaces

import "github.com/FrosTiK-SD/mess-backend/models"

type PopulatedRoom struct {
	models.Room
	AllottedTo []models.User `json:"allottedTo"`
}
