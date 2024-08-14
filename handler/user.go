package handler

import (
	"net/http"
	"strconv"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/FrosTiK-SD/mess-backend/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler *Handler) CreateUser(ctx *fiber.Ctx) error {
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := controller.CreateNewUser(handler.MongikClient, &user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.SendStatus(201)
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

func (handler *Handler) UpdateUser(ctx *fiber.Ctx) error {
	var updatedUser models.User

	if err := ctx.BodyParser(&updatedUser); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userID, errObjID := primitive.ObjectIDFromHex(ctx.Get("userID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": updatedUser}

	if _, err := collection.UpdateOne(ctx.Context(), filter, update); err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "User Updated", nil, nil))
}

func (handler *Handler) DeleteUser(ctx *fiber.Ctx) error {
	userID, errObjID := primitive.ObjectIDFromHex(ctx.Get("userID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	filter := bson.M{"_id": userID}

	if _, err := collection.DeleteOne(ctx.Context(), filter); err != nil {
		return err
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "User Deleted", nil, nil))
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

// Authenticated
func (h *Handler) GetUserFromToken(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals(constants.SESSION).(models.User)

	if !ok {
		return fiber.NewError(fiber.StatusForbidden, "Authentication Failed")
	}

	return ctx.JSON(fiber.Map{
		"user": user,
	})
}

// Unauthenticated
func (h *Handler) CreateUserFromToken(ctx *fiber.Ctx) error {
	idToken := ctx.Get("token", "")
	noCache := ctx.Get(constants.CACHE_CONTROL_HEADER, "") == constants.NO_CACHE

	email, _, errStr := utils.VerifyToken(h.MongikClient.CacheClient, idToken, h.JwkSet, noCache)
	if errStr != nil {
		return fiber.NewError(fiber.StatusBadRequest, *errStr)
	}
	if !utils.IsInstituteEmail(*email) {
		return fiber.NewError(fiber.StatusBadRequest, "Token Registration can only be done with Institute Mail")
	}

	user := new(models.User)
	user.Email = *email

	err := controller.CreateNewUser(h.MongikClient, user)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"user": *user,
	})
}

func (h *Handler) GetFilteredUsers(ctx *fiber.Ctx) error {
	var userFilter interfaces.UserFilter
	if err := ctx.BodyParser(&userFilter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Improper Request Body")
	}

	users, err := controller.GetUserFromFilter(h.MongikClient, &userFilter)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"users": users,
	})
}

func (h *Handler) AssignHostelToUsers(ctx *fiber.Ctx) error {
	var reqBody interfaces.AssignHostelToUsersRequestBody

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err := controller.AssignHostelToUsers(h.MongikClient, reqBody.Hostel, reqBody.Users)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (h *Handler) AssignMessToUsers(ctx *fiber.Ctx) error {
	var reqBody interfaces.AssignMessToUsersRequestBody

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := controller.AssignMessToUsers(h.MongikClient, reqBody.Mess, reqBody.Users)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (h *Handler) AssignRoomToUser(ctx *fiber.Ctx) error {
	type RequestBody struct {
		UserID primitive.ObjectID `json:"userId"`
		Room   primitive.ObjectID `json:"room"`
	}
	var reqBody RequestBody

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := controller.AssignRoomToUser(h.MongikClient, reqBody.UserID, reqBody.Room)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return nil
}

func (h *Handler) GetUserByRollNo(ctx *fiber.Ctx) error {
	rollNo, err := strconv.ParseInt(ctx.Params("rollNo", ""), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Could not parse rollNo")
	}

	user, err := controller.GetUserByRollNo(h.MongikClient, rollNo, false)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"user": user,
	})
}
