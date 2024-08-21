package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) CreateRoomAllotment(ctx *fiber.Ctx) error {
	var roomAllotment models.RoomAllotment

	if err := ctx.BodyParser(&roomAllotment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := controller.CreateRoomAllotment(h.MongikClient, &roomAllotment); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"roomAllotment": roomAllotment,
	})

}

func (h *Handler) GetSemesterRoomAllotmentsWithUser(ctx *fiber.Ctx) error {
	params := struct {
		SemesterID primitive.ObjectID `params:"semesterID" binding:"required"`
		RoomID     primitive.ObjectID `params:"roomID" binding:"required"`
	}{}

	if err := ctx.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	roomAllotments, err := controller.GetSemesterRoomAllotmentsWithUser(h.MongikClient, params.SemesterID, params.RoomID)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"roomAllotments": roomAllotments,
	})
}
