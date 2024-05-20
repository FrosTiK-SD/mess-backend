package handler

import (
	"net/http"
	"time"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler *Handler) CreateRoom(ctx *fiber.Ctx) error {
	var Room models.Room

	// Parse JSON body
	if err := ctx.BodyParser(&Room); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Room.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_ROOMS)
	if result, err := collection.InsertOne(ctx.Context(), Room); err != nil {
		return err
	} else {
		ctx.JSON(interfaces.GetGenericResponse(true, "Mess Created", result, nil))
	}

	return nil
}

func (handler *Handler) RegisterUser(ctx *fiber.Ctx) error {
	var User models.User

	// Parse JSON body
	if err := ctx.BodyParser(&User); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_USERS)
	if result, err := collection.InsertOne(ctx.Context(), User); err != nil {
		return err
	} else {
		ctx.JSON(interfaces.GetGenericResponse(true, "User Created", result, nil))
	}

	return nil
}

func (handler *Handler) GetManagedMessess(ctx *fiber.Ctx) error {
	ctx.JSON(interfaces.GetGenericResponse(true, "hello", nil, nil))
	return nil
}
