package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateHostelStaffAllotment(ctx *fiber.Ctx) error {
	var hostel models.HostelStaffAllotment

	if err := ctx.BodyParser(&hostel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := controller.CreateHostelStaffAllotment(h.MongikClient, &hostel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"hostel": hostel,
	})

}
