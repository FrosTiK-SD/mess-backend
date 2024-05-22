package handler

import (
	"github.com/FrosTiK-SD/mess-backend/interfaces"
	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Hello(ctx *fiber.Ctx) error {
	ctx.JSON(interfaces.GetGenericResponse(true, "hello", nil, nil))
	return nil
}
