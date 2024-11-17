package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BatchCreateMessAllotment(mongikClient *mongikModels.Mongik, messID primitive.ObjectID, semesterID primitive.ObjectID, userIDs []primitive.ObjectID) error {
	var messAllotments = make([]models.MessAllotment, len(userIDs))
	for idx := range messAllotments {
		messAllotments[idx].ID = primitive.NewObjectID()
		messAllotments[idx].Mess = messID
		messAllotments[idx].Semester = semesterID
		messAllotments[idx].User = userIDs[idx]
	}
	_, err := mongikDB.InsertMany(mongikClient, constants.DB, constants.COLLECTION_MESS_ALLOTMENTS, messAllotments)

	return err
}
