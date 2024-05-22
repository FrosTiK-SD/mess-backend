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

func (handler *Handler) CreateMeal(ctx *fiber.Ctx) error {
	var Meal models.Meal

	if err := ctx.BodyParser(&Meal); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Meal.ID = primitive.NewObjectID()

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MEALS)
	if result, err := collection.InsertOne(ctx.Context(), Meal); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "Mess Created", result, nil))
	}
}

func (handler *Handler) GetMeal(ctx *fiber.Ctx) error {
	var Meal models.Meal
	mealID, errObjID := primitive.ObjectIDFromHex(ctx.Get("mealID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MEALS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": mealID}).Decode(&Meal); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found Meal with the given Meal ID", Meal, nil))
}
