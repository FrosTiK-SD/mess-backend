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

func CreateNewUser(mongikClient *mongikModels.Mongik, user *models.User) error {
	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	_, err := mongikDB.InsertOne[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, *user)

	return err
}

func GetUserFromFilter(mongikClient *mongikModels.Mongik, userFilter *interfaces.UserFilter) ([]models.User, error) {

	matchStatement := bson.M{}

	if len(userFilter.RollNo) != 0 {
		matchStatement["rollNo"] = bson.M{
			"$in": userFilter.RollNo,
		}
	}

	if len(userFilter.EndYear) != 0 {
		matchStatement["endYear"] = bson.M{
			"$in": userFilter.EndYear,
		}
	}

	if len(userFilter.Course) != 0 {
		matchStatement["course"] = bson.M{
			"$in": userFilter.Course,
		}
	}

	if len(userFilter.Department) != 0 {
		matchStatement["department"] = bson.M{
			"$in": userFilter.Department,
		}
	}

	if len(userFilter.StartYear) != 0 {
		matchStatement["startYear"] = bson.M{
			"$in": userFilter.StartYear,
		}
	}

	pipeline := []bson.M{{
		"$match": matchStatement,
	}}
	users, err := mongikDB.Aggregate[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, pipeline, false)

	return users, err
}

func AssignHostelToUsers(mongikClient *mongikModels.Mongik, hostel primitive.ObjectID, users []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{"$in": users},
	}
	update := bson.M{
		"$set": bson.M{
			"allocationDetails.hostel": hostel,
		}}
	_, err := mongikDB.UpdateMany[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, filter, update)

	return err
}

func AssignMessToUsers(mongikClient *mongikModels.Mongik, mess primitive.ObjectID, users []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{"$in": users},
	}
	update := bson.M{
		"$set": bson.M{
			"allocationDetails.mess": mess,
		},
	}
	_, err := mongikDB.UpdateMany[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, filter, update)

	return err
}

func GetUserByRollNo(mongikClient *mongikModels.Mongik, rollNo int64, noCache bool) (models.User, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"rollNo": rollNo}},
	}

	user, err := mongikDB.AggregateOne[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, pipeline, noCache)

	return user, err
}

func GetUserByRole(mongikClient *mongikModels.Mongik, role constants.Role) ([]models.User, error) {
	pipeline := []bson.M{{
		"$match": bson.M{
			"role": role,
		},
	}}

	users, err := mongikDB.Aggregate[models.User](mongikClient, constants.DB, constants.COLLECTION_USERS, pipeline, false)

	if users == nil {
		users = []models.User{}
	}
	return users, err
}
