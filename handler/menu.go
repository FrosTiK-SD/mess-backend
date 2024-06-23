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

func (handler *Handler) CreateMenuItem(ctx *fiber.Ctx) error {
	var MenuItem models.MenuItem

	if err := ctx.BodyParser(&MenuItem); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	MenuItem.ID = primitive.NewObjectID()

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MENU_ITEMS)
	if result, err := collection.InsertOne(ctx.Context(), MenuItem); err != nil {
		return err
	} else {
		return ctx.JSON(interfaces.GetGenericResponse(true, "MenuItem Created", result, nil))
	}
}

func (handler *Handler) GetMenuItem(ctx *fiber.Ctx) error {
	var MenuItem models.MenuItem
	menuItemID, errObjID := primitive.ObjectIDFromHex(ctx.Get("menuItemID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MENU_ITEMS)
	if errFind := collection.FindOne(ctx.Context(), bson.M{"_id": menuItemID}).Decode(&MenuItem); errFind != nil {
		return errFind
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found MenuItem with the given ID", MenuItem, nil))
}

func (handler *Handler) GetAllMenuItemsOfAMess(ctx *fiber.Ctx) error {
	var MenuItems []models.MenuItem
	messID, errObjID := primitive.ObjectIDFromHex(ctx.Get("messID"))
	if errObjID != nil {
		return errObjID
	}

	collection := handler.MongikClient.MongoClient.Database(constants.DB).Collection(constants.COLLECTION_MENU_ITEMS)
	if cursor, errFind := collection.Find(ctx.Context(), bson.M{"mess": messID}); errFind != nil {
		return errFind
	} else {
		if errDecode := cursor.All(ctx.Context(), &MenuItems); errDecode != nil {
			return errDecode
		}
	}

	return ctx.JSON(interfaces.GetGenericResponse(true, "Found all MenuItems of the given mess", MenuItems, nil))
}
