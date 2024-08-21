package interfaces

import (
	"github.com/FrosTiK-SD/mess-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomAllotmentWithUser struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Room     primitive.ObjectID `json:"room" bson:"room" binding:"required"`
	User     models.User        `json:"user" bson:"user" binding:"required"`
	Semester primitive.ObjectID `json:"semester" bson:"semester" binding:"required"`
}
