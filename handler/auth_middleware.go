package handler

import (
	"fmt"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) FiberAuthenticateUser(ctx *fiber.Ctx) error {
	idToken := ctx.Get("token", "")
	fmt.Println(idToken)
	noCache := ctx.Get(constants.CACHE_CONTROL_HEADER, "") == constants.NO_CACHE

	email, _, errStr := utils.VerifyToken(h.MongikClient.CacheClient, idToken, h.JwkSet, noCache)
	fmt.Println(*email)
	if errStr != nil {
		return fiber.NewError(fiber.StatusForbidden, *errStr)
	}

	user, err := controller.GetUserByEmail(h.MongikClient, email, noCache)
	fmt.Println(user)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, err.Error())
	}

	ctx.Locals(constants.SESSION, user)

	return ctx.Next()

}
