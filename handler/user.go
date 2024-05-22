package handler

import (
	"net/http"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) CreateUser(ctx *fiber.Ctx) error {
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
