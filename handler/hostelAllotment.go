package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) BatchCreateHostelAllotments(ctx *fiber.Ctx) error {
	var reqBody = struct {
		Semester primitive.ObjectID   `json:"semester" binding:"required"`
		Hostel   primitive.ObjectID   `json:"hostel" binding:"required"`
		Users    []primitive.ObjectID `json:"users" binding:"required"`
	}{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	err := controller.BatchCreateHostelAllotments(*h.MongikClient, reqBody.Semester, reqBody.Hostel, reqBody.Users)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
