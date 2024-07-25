package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserByEmail(mongikClient *mongikModels.Mongik, email *string, noCache bool) (models.User, error) {

	pipline := []bson.M{
		{
			"$match": bson.M{
				"email": *email,
			},
		},
	}

	user, err := mongikDB.AggregateOne[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, pipline, noCache)

	return user, err

}

func GetUserPopulatedByEmail(mongikClient *mongikModels.Mongik, email *string, noCache bool) (interfaces.UserPopulated, error) {

	pipline := []bson.M{
		{
			"$match": bson.M{
				"email": *email,
			},
		}, {
			"$lookup": bson.M{
				"from":         "groups",
				"localField":   "groups",
				"foreignField": "_id",
				"as":           "groups",
			},
		},
	}

	user, err := mongikDB.AggregateOne[interfaces.UserPopulated](mongikClient, constants.DB, constants.COLLECTION_USERS, pipline, noCache)

	return user, err

}

func CreateNewUser(mongikClient *mongikModels.Mongik, user *models.User) error {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}
	if user.Groups == nil {
		user.Groups = make([]primitive.ObjectID, 0)
	}

	_, err := mongikDB.InsertOne[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, *user)

	return err
}
