package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateRooms(mongikClient *mongikModels.Mongik, hostelId primitive.ObjectID, rangeStart int, rangeEnd int) error {
	newRooms := make([]models.Room, rangeEnd-rangeStart+1)
	for idx := range newRooms {
		newRooms[idx] = models.Room{
			ID:        primitive.NewObjectID(),
			Hostel:    hostelId,
			Number:    rangeStart + idx,
			Floor:     0,
			Available: true,
			Note:      "",
		}
	}
	_, err := mongikDB.InsertMany[models.Room](mongikClient, constants.DB, constants.COLLECTION_ROOMS, newRooms)

	return err
}
