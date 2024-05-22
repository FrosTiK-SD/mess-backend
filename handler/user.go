package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler *Handler) CreateUser(ctx *fiber.Ctx) error {
	var User models.User

	if err := ctx.BodyParser(&User); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	User.ID = primitive.NewObjectID()

	if User.ManagingDetails.Hostels == nil {
		User.ManagingDetails.Hostels = make([]primitive.ObjectID, 0)
	}

	if User.ManagingDetails.Messes == nil {
		User.ManagingDetails.Messes = make([]primitive.ObjectID, 0)
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	if result, err := collection.InsertOne(ctx.Context(), User); err != nil {
		return err
	} else {
		ctx.JSON(interfaces.GetGenericResponse(true, "User Created", result, nil))
	}

	return nil
}

func (handler *Handler) ManageHostelMess(ctx *fiber.Ctx) error {
	var allocation interfaces.HostelMessManageToUser

	if err := ctx.BodyParser(&allocation); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	filter := bson.D{{Key: "_id", Value: allocation.UserID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "managingDetails.hostels", Value: allocation.HostelIDs}, {Key: "managingDetails.messes", Value: allocation.MessIDs}}}}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	if result, err := collection.UpdateOne(ctx.Context(), filter, update); err != nil {
		return err
	} else {
		ctx.JSON(interfaces.GetGenericResponse(true, "Replaced User's hostels and messes with current selection", result, nil))
	}

	return nil
}
