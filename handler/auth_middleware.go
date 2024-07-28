package handler

import (
	"errors"
	"fmt"
	"slices"

	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/interfaces"
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

type Predicate = func(*fiber.Ctx) error

func (h *Handler) GetAccessControlHandler(predicate Predicate) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		if err := predicate(ctx); err != nil {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}

		return ctx.Next()
	}
}

func HasRoles(roles ...string) Predicate {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals(constants.SESSION).(interfaces.UserPopulated)
		if !ok {
			return errors.New("Could not find user")
		}
		allRoles := utils.GetAllRoles(&user)
		for ridx := range roles {
			if !slices.Contains(allRoles, roles[ridx]) {
				return errors.New(fmt.Sprintf(`User Does not have Role %s`, roles[ridx]))
			}
		}

		return nil
	}
}

func And(predicates ...Predicate) Predicate {
	return func(c *fiber.Ctx) error {
		var err error

		for pidx := range predicates {
			err = predicates[pidx](c)
			if err != nil {
				return err
			}
		}

		return nil
	}
}
