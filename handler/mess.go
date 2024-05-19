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

func (handler *Handler) CreateMess(ctx *fiber.Ctx) error {
	var Mess models.Mess

	// Parse JSON body
	if err := ctx.BodyParser(&Mess); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Mess.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MESSES)
	if result, err := collection.InsertOne(ctx.Context(), Mess); err != nil {
		return err
	} else {
		ctx.JSON(interfaces.GetGenericResponse(true, "Mess Created", result, nil))
	}

	return nil
}
