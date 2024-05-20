package handler

import (
	"net/http"
	"time"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (handler *Handler) CreateMess(ctx *fiber.Ctx) error {
	var Mess models.Mess

	if err := ctx.BodyParser(&Mess); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Mess.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MESSES)
	if result, err := collection.InsertOne(ctx.Context(), Mess); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "Mess Created", result, nil))
	}
}

func (handler *Handler) GetMess(ctx *fiber.Ctx) error {
	var Mess models.Mess
	messID, errObjID := primitive.ObjectIDFromHex(ctx.GetReqHeaders()["messID"][0])
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MESSES)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": messID}).Decode(Mess); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found Mess with the given Mess ID", Mess, nil))
}
