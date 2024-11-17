package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) BatchCreateMessAllotment(ctx *fiber.Ctx) error {
	reqBody := struct {
		Semester primitive.ObjectID   `json:"semester" binding:"required"`
		Mess     primitive.ObjectID   `json:"mess" binding:"required"`
		Users    []primitive.ObjectID `json:"users" binding:"required"`
	}{}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := controller.BatchCreateMessAllotment(h.MongikClient, reqBody.Mess, reqBody.Semester, reqBody.Users)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
