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
		return ctx.JSON(interfaces.GetGenericResponse(true, "Meal Created", result, nil))
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

func (handler *Handler) CreateMealType(ctx *fiber.Ctx) error {
	var MealType models.MealType

	if err := ctx.BodyParser(&MealType); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	MealType.ID = primitive.NewObjectID()

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MEALS)
	if result, err := collection.InsertOne(ctx.Context(), MealType); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "MealType Created", result, nil))
	}
}

func (handler *Handler) GetMealType(ctx *fiber.Ctx) error {
	var MealType models.MealType
	mealID, errObjID := primitive.ObjectIDFromHex(ctx.Get("mealID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MEAL_TYPES)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": mealID}).Decode(&MealType); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found MealType with the given Meal ID", MealType, nil))
}

func (handler *Handler) GetAllMealTypesOfAMess(ctx *fiber.Ctx) error {
	var MealTypes []models.MealType
	messID, errObjID := primitive.ObjectIDFromHex(ctx.Get("messID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MEAL_TYPES)
	if cursor, errFind := collection.Find(ctx.Context(), bson.M{"mess": messID}); errFind != nil {
		return errFind
	} else {
		if errDecode := cursor.All(ctx.Context(), &MealTypes); errDecode != nil {
			return errDecode
		}
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found all MealTypes of the given mess", MealTypes, nil))
}

func (handler *Handler) UpdateMealsByDate(ctx *fiber.Ctx) error {
	u := new(interfaces.UpdateMeal)
	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MEALS)

	if err := ctx.BodyParser(u); err != nil {
		return err
	}

	filter := bson.M{
		"mess":     u.MessID,
		"mealType": u.MealType,
		"date": bson.M{
			"$in": u.Dates,
		},
	}

	var update bson.M
	if u.AppendOnly {
		update = bson.M{
			"$push": bson.M{
				"menu": bson.M{
					"$each": u.Menu,
				},
			},
		}
	} else {
		update = bson.M{
			"$set": bson.M{
				"menu": u.Menu,
			},
		}
	}

	if result, err := collection.UpdateMany(ctx.Context(), filter, update); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "Updated Meals", result, nil))
	}
}
