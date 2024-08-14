package interfaces

import "github.com/FrosTiK-SD/mess-backend/models"

type PopulatedRoom struct {
	models.Room
	AllocatedTo []models.User `json:"allocatedTo"`
}
