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

func CreateRoomAllotment(mongikClient *mongikModels.Mongik, roomAllotment *models.RoomAllotment) error {
	if roomAllotment.ID.IsZero() {
		roomAllotment.ID = primitive.NewObjectID()
	}
	_, err := mongikDB.InsertOne[models.RoomAllotment](mongikClient, constants.DB, constants.COLLECTION_ROOM_ALLOTMENTS, *roomAllotment)
	return err
}

func GetSemesterRoomAllotmentsWithUser(mongikClient *mongikModels.Mongik, semesterID primitive.ObjectID, roomID primitive.ObjectID) ([]interfaces.RoomAllotmentWithUser, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"room":     roomID,
				"semester": semesterID,
			},
		},
		{
			"$lookup": bson.M{
				"from":         constants.COLLECTION_USERS,
				"localField":   "user",
				"foreignField": "_id",
				"as":           "user",
			},
		},
		{
			"$unwind": bson.M{
				"path": "$user",
			},
		},
	}

	roomAllotments, err := mongikDB.Aggregate[interfaces.RoomAllotmentWithUser](mongikClient, constants.DB, constants.COLLECTION_ROOM_ALLOTMENTS, pipeline, false)

	if roomAllotments == nil {
		roomAllotments = []interfaces.RoomAllotmentWithUser{}
	}

	return roomAllotments, err
}
