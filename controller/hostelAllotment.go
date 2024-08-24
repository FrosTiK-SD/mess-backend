package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BatchCreateHostelAllotments(mongikClient mongikModels.Mongik, semesterID primitive.ObjectID, hostelID primitive.ObjectID, userIDs []primitive.ObjectID) error {
	var hostelAllotments = make([]models.HostelAllotment, len(userIDs))
	for idx := range hostelAllotments {
		hostelAllotments[idx].ID = primitive.NewObjectID()
		hostelAllotments[idx].Hostel = hostelID
		hostelAllotments[idx].Semester = semesterID
		hostelAllotments[idx].User = userIDs[idx]
	}
	_, err := mongikDB.InsertMany[models.HostelAllotment](&mongikClient, constants.DB, constants.COLLECTION_HOSTEL_ALLOTMENTS, hostelAllotments)
	return err
}
