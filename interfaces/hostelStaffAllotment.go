package interfaces

import (
	"github.com/FrosTiK-SD/mess-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HostelStaffAllotmentWithUser struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Hostel primitive.ObjectID `json:"hostel" bson:"hostel" binding:"required"`
	User   models.User        `json:"user" bson:"user" binding:"required"`
}
