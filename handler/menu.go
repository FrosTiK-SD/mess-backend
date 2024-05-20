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

func (handler *Handler) CreateMenuItem(ctx *fiber.Ctx) error {
	var MenuItem models.MenuItem

	if err := ctx.BodyParser(&MenuItem); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	MenuItem.ID = primitive.NewObjectIDFromTimestamp(time.Now())

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MENUITEMS)
	if result, err := collection.InsertOne(ctx.Context(), MenuItem); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "Menu Item Created", result, nil))
	}
}

func (handler *Handler) GetMenuItem(ctx *fiber.Ctx) error {
	var MenuItem models.MenuItem
	menuItemID, errObjID := primitive.ObjectIDFromHex(ctx.GetReqHeaders()["menuItemID"][0])
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MENUITEMS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": menuItemID}).Decode(MenuItem); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found Menu Item with the given ID", MenuItem, nil))
}
