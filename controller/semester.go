package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllSemesters(mongikClient *mongikModels.Mongik) ([]models.Semester, error) {
	pipeline := []bson.M{{
		"$sort": bson.M{
			"end": -1,
		},
	}}
	semesters, err := mongikDB.Aggregate[models.Semester](mongikClient, constants.DB, constants.COLLECTION_SEMESTERS, pipeline, false)

	if semesters == nil {
		semesters = []models.Semester{}
	}

	return semesters, err
}

func GetSemesterById(mongikClient *mongikModels.Mongik, semesterID primitive.ObjectID) (*models.Semester, error) {

	pipeline := []bson.M{{
		"$match": bson.M{
			"_id": semesterID,
		},
	}}

	semester, err := mongikDB.AggregateOne[models.Semester](mongikClient, constants.DB, constants.COLLECTION_SEMESTERS, pipeline, false)

	return &semester, err
}

func CreateSemester(mongikClient *mongikModels.Mongik, semester *models.Semester) error {
	if semester.ID.IsZero() {
		semester.ID = primitive.NewObjectID()
	}

	_, err := mongikDB.InsertOne[models.Semester](mongikClient, constants.DB, constants.COLLECTION_SEMESTERS, *semester)

	return err
}

func UpdateSemester(mongikClient *mongikModels.Mongik, semesterID primitive.ObjectID, semester *models.Semester) error {
	semester.ID = semesterID
	query := bson.M{
		"_id": semesterID,
	}
	update := bson.M{
		"$set": *semester,
	}
	_, err := mongikDB.UpdateOne[models.Semester](mongikClient, constants.DB, constants.COLLECTION_SEMESTERS, query, update)

	return err
}
