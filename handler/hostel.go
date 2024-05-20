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

func (handler *Handler) CreateHostel(ctx *fiber.Ctx) error {
	var Hostel models.Hostel

	if err := ctx.BodyParser(&Hostel); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Hostel.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	if result, err := collection.InsertOne(ctx.Context(), Hostel); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "Menu Item Created", result, nil))
	}
}

func (handler *Handler) GetHostel(ctx *fiber.Ctx) error {
	var Hostel models.Hostel
	hostelID, errObjID := primitive.ObjectIDFromHex(ctx.Get("hostelID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_HOSTELS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": hostelID}).Decode(&Hostel); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found Hostel with the given ID", Hostel, nil))
}
