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

func BatchCreateRooms(mongikClient *mongikModels.Mongik, hostelId primitive.ObjectID, batchReq *interfaces.BatchCreateHostelRoomsRequest) error {
	newRooms := make([]models.Room, batchReq.RangeEnd-batchReq.RangeStart+1)
	for idx := range newRooms {
		newRooms[idx] = models.Room{
			ID:        primitive.NewObjectID(),
			Hostel:    hostelId,
			Number:    batchReq.RangeStart + idx,
			Floor:     batchReq.Floor,
			Occupancy: batchReq.Occupancy,
			Available: batchReq.Available,
		}
	}
	_, err := mongikDB.InsertMany[models.Room](mongikClient, constants.DB, constants.COLLECTION_ROOMS, newRooms)

	return err
}

func GetRoomsByHostelId(mongikClient *mongikModels.Mongik, hostelID primitive.ObjectID) ([]interfaces.PopulatedRoom, error) {

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"hostel": hostelID,
			},
		},
		{
			"$set": bson.M{
				"allottedTo": []interfaces.PopulatedRoom{},
			},
		},
	}
	rooms, err := mongikDB.Aggregate[interfaces.PopulatedRoom](mongikClient, constants.DB, constants.COLLECTION_ROOMS, pipeline, false)
	if rooms == nil {
		rooms = make([]interfaces.PopulatedRoom, 0)
	}

	return rooms, err
}
