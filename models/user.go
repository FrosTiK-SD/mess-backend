package models

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonalDetails struct {
	// TODO add stuff here
}
type CaretakerDetails struct {
	// Struct might require renaming
	Hostels []primitive.ObjectID `json:"hostels" bson:"hostels"`
}

type User struct {
	ID               primitive.ObjectID     `json:"_id,omitempty" bson:"_id"`
	Permissions      []constants.Permission `json:"permissions" bson:"permissions" binding:"required"`
	Groups           []UserGroup            `json:"groups" bson:"groups"`
	Hostel           primitive.ObjectID     `json:"hostel" bson:"hostel"`
	Mess             primitive.ObjectID     `json:"mess" bson:"mess"`
	Room             primitive.ObjectID     `json:"room" bson:"room"`
	CaretakerDetails CaretakerDetails       `json:"caretakerDetails,omitempty" bson:"caretakerDetails,omitempty"`
	PersonalDetails  PersonalDetails        `json:"personalDetails,omitempty" bson:"personalDetails,omitempty"`
}
