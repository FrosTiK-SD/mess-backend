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

	if hostels == nil {
		hostels = make([]models.Hostel, 0)
	}

	return hostels, err
}

func GetHostelById(mongikClient *mongikModels.Mongik, hostelID primitive.ObjectID) (models.Hostel, error) {
	pipeline := []bson.M{{
		"$match": bson.M{
			"_id": hostelID,
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

func UpdateHostel(mongikClient *mongikModels.Mongik, hostelID primitive.ObjectID, updatedHostel *models.Hostel) error {
	updatedHostel.ID = hostelID

	query := bson.M{
		"_id": hostelID,
	}
	update := bson.M{
		"$set": updatedHostel,
	}
	_, err := mongikDB.UpdateOne[models.Hostel](mongikClient, constants.DB, constants.COLLECTION_HOSTELS, query, update)

	return err
}
