package models

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGroup struct {
	ID          primitive.ObjectID     `json:"_id" bson:"_id"`
	Name        string                 `json:"name" bson:"name"`
	Description string                 `json:"description" bson:"description"`
	Permissions []constants.Permission `json:"permissions" bson:"permissions"`
}
