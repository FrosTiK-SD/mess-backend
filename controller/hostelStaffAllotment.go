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

func CreateHostelStaffAllotment(mongikClient *mongikModels.Mongik, hostelStaffAllotment *models.HostelStaffAllotment) error {

	if hostelStaffAllotment.ID.IsZero() {
		hostelStaffAllotment.ID = primitive.NewObjectID()
	}

	_, err := mongikDB.InsertOne[models.HostelStaffAllotment](mongikClient, constants.DB, constants.COLLECTION_HOSTEL_STAFF_ALLOTMENTS, *hostelStaffAllotment)

	return err
}

func GetHostelStaffAllotmentWithUser(mongikClient *mongikModels.Mongik, hostelID primitive.ObjectID) ([]interfaces.HostelStaffAllotmentWithUser, error) {
	pipeline := []bson.M{{
		"$match": bson.M{
			"hostel": hostelID,
		},
	}, {
		"$lookup": bson.M{
			"from":         constants.COLLECTION_USERS,
			"localField":   "user",
			"foreignField": "_id",
			"as":           "user",
		},
	}, {
		"$unwind": bson.M{
			"path": "$user",
		},
	}}

	staffAllotments, err := mongikDB.Aggregate[interfaces.HostelStaffAllotmentWithUser](mongikClient, constants.DB, constants.COLLECTION_HOSTEL_STAFF_ALLOTMENTS, pipeline, false)

	if staffAllotments == nil {
		staffAllotments = []interfaces.HostelStaffAllotmentWithUser{}
	}
	return staffAllotments, err
}
