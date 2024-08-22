package controller

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/models"
	mongikDB "github.com/FrosTiK-SD/mongik/db"
	mongikModels "github.com/FrosTiK-SD/mongik/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateHostelStaffAllotment(mongikClient *mongikModels.Mongik, hostelStaffAllotment *models.HostelStaffAllotment) error {

	if hostelStaffAllotment.ID.IsZero() {
		hostelStaffAllotment.ID = primitive.NewObjectID()
	}

	_, err := mongikDB.InsertOne[models.HostelStaffAllotment](mongikClient, constants.DB, constants.COLLECTION_HOSTEL_STAFF_ALLOTMENTS, *hostelStaffAllotment)

	return err
}
