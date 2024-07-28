package handler

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) FiberAuthenticateUser(ctx *fiber.Ctx) error {
	idToken := ctx.Get("token", "")
	noCache := ctx.Get(constants.CACHE_CONTROL_HEADER, "") == constants.NO_CACHE

	email, _, errStr := utils.VerifyToken(h.MongikClient.CacheClient, idToken, h.JwkSet, noCache)
	if errStr != nil {
		return fiber.NewError(fiber.StatusForbidden, *errStr)
	}

	user, err := controller.GetUserPopulatedByEmail(h.MongikClient, email, noCache)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	ctx.Locals(constants.SESSION, user)

	return ctx.Next()

}
