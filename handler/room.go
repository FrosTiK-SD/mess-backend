package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetHostelRooms(ctx *fiber.Ctx) error {
	hostelID, err := primitive.ObjectIDFromHex(ctx.Params("hostelID"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	rooms, err := controller.GetRoomsByHostelId(h.MongikClient, hostelID)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"rooms": rooms,
	})
}

func (h *Handler) BatchCreateHostelRooms(ctx *fiber.Ctx) error {

	var reqBody interfaces.BatchCreateHostelRoomsRequest

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	hostelID, err := primitive.ObjectIDFromHex(ctx.Params("hostelID"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = controller.BatchCreateRooms(h.MongikClient, hostelID, &reqBody)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated)

}
