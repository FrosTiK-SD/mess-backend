package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMess(mongikClient *mongikModels.Mongik, mess *models.Mess) error {
	if mess.ID.IsZero() {
		mess.ID = primitive.NewObjectID()
	}
	_, err := mongikDB.InsertOne[models.Mess](mongikClient, constants.DB, constants.COLLECTION_MESSES, *mess)

	return err
}

func GetAllMesses(mongikClient *mongikModels.Mongik) ([]models.Mess, error) {

	pipeline := []bson.M{}
	messes, err := mongikDB.Aggregate[models.Mess](mongikClient, constants.DB, constants.COLLECTION_MESSES, pipeline, false)

	if messes == nil {
		messes = []models.Mess{}
	}
	return messes, err
}
