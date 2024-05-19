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

func (handler *Handler) CreateMeal(ctx *fiber.Ctx) error {
	var Meal models.Meal

	// Parse JSON body
	if err := ctx.BodyParser(&Meal); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Meal.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MEALS)
	if result, err := collection.InsertOne(ctx.Context(), Meal); err != nil {
		return err
	} else {
		ctx.JSON(interfaces.GetGenericResponse(true, "Mess Created", result, nil))
	}

	return nil
}
