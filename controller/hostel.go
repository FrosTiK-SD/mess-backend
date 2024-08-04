package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllHostels(mongikClient *mongikModels.Mongik) ([]models.Hostel, error) {
	hostels, err := mongikDB.Find[models.Hostel](mongikClient, constants.DB, constants.COLLECTION_HOSTELS, bson.M{}, false)

	return hostels, err
}

func GetHostelById(mongikClient *mongikModels.Mongik, hostelId primitive.ObjectID) (models.Hostel, error) {
	pipeline := []bson.M{{
		"$match": bson.M{
			"_id": hostelId,
		},
	}}
	hostel, err := mongikDB.AggregateOne[models.Hostel](mongikClient, constants.DB, constants.COLLECTION_HOSTELS, pipeline, false)

	return hostel, err
}

func CreateHostel(mongikClient *mongikModels.Mongik, hostel *models.Hostel) error {
	if hostel.ID.IsZero() {
		hostel.ID = primitive.NewObjectID()
	}

	_, err := mongikDB.InsertOne[models.Hostel](mongikClient, constants.DB, constants.COLLECTION_HOSTELS, *hostel)

	return err

}
