package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *Handler) GetHostelStaffAllotmentWithUser(ctx *fiber.Ctx) error {
	params := struct {
		HostelID primitive.ObjectID `params:"hostelID" binding:"required"`
	}{}

	if err := ctx.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	staffAllotments, err := controller.GetHostelStaffAllotmentWithUser(h.MongikClient, params.HostelID)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"hostelStaffAllotments": staffAllotments,
	})
}
