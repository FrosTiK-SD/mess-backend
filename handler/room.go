package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GenerateRooms(ctx *fiber.Ctx) error {
	type ReqBody struct {
		HostelID   primitive.ObjectID `json:"hostelId" binding:"required"`
		RangeStart int                `json:"rangeStart" binding:"required"`
		RangeEnd   int                `json:"rangeEnd" binding:"required"`
	}
	var reqBody ReqBody

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(400, err.Error())
	}

	err := controller.GenerateRooms(h.MongikClient, reqBody.HostelID, reqBody.RangeStart, reqBody.RangeEnd)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
