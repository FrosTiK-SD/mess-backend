package handler

import (
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) RegisterStudent(ctx *fiber.Ctx) error {
	ctx.JSON(interfaces.GetGenericResponse(true, "", nil, nil))
	return nil
}
