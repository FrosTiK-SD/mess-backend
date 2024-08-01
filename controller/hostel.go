package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateHostel(mongikClient *mongikModels.Mongik, hostel *models.Hostel) error {
	if hostel.ID.IsZero() {
		hostel.ID = primitive.NewObjectID()
	}

	_, err := mongikDB.InsertOne[models.Hostel](mongikClient, constants.DB, constants.COLLECTION_HOSTELS, *hostel)

	return err

}
