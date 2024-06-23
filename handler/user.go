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

func (handler *Handler) GetUser(ctx *fiber.Ctx) error {
	var User models.User
	userID, errObjID := primitive.ObjectIDFromHex(ctx.Get("userID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": userID}).Decode(&User); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found User with the given Mess ID", User, nil))
}

func (handler *Handler) GetUserPopulated(ctx *fiber.Ctx) error {
	var Users []map[string]interface{}
	userID, errObjID := primitive.ObjectIDFromHex(ctx.Get("userID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)

	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: userID}}}},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: constants.COLLECTION_GROUPS},
				{Key: "localField", Value: "groups"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "groups"},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: constants.COLLECTION_HOSTELS},
				{Key: "localField", Value: "allocationDetails.hostel"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "allocationDetails.hostel"},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: constants.COLLECTION_MESSES},
				{Key: "localField", Value: "allocationDetails.mess"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "allocationDetails.mess"},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: constants.COLLECTION_ROOMS},
				{Key: "localField", Value: "allocationDetails.room"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "allocationDetails.room"},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: constants.COLLECTION_HOSTELS},
				{Key: "localField", Value: "managingDetails.hostels"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "managingDetails.hostels"},
			}},
		},
		bson.D{
			{Key: "$lookup", Value: bson.D{
				{Key: "from", Value: constants.COLLECTION_MESSES},
				{Key: "localField", Value: "managingDetails.messes"},
				{Key: "foreignField", Value: "_id"},
				{Key: "as", Value: "managingDetails.messes"},
			}},
		},
	}

	if cursor, err := collection.Aggregate(ctx.Context(), pipeline); err != nil {
		return err
	} else {
		if errBind := cursor.All(ctx.Context(), &Users); errBind != nil {
			return errBind
		}
		return ctx.JSON(interfaces.GetGenericResponse(true, "Fetched Mess Dashboard", Users, nil))
	}
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
