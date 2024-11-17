package handler

import (
	"github.com/FrosTiK-SD/mess-backend/controller"
	"github.com/FrosTiK-SD/mess-backend/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) GetAllSemesters(ctx *fiber.Ctx) error {

	semester, err := controller.GetAllSemesters(h.MongikClient)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"semesters": semester,
	})
}

func (h *Handler) GetSemesterById(ctx *fiber.Ctx) error {
	semesterID, err := primitive.ObjectIDFromHex(ctx.Params("semesterID"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	semester, err := controller.GetSemesterById(h.MongikClient, semesterID)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return ctx.JSON(fiber.Map{
		"semester": *semester,
	})
}

func (h *Handler) CreateSemester(ctx *fiber.Ctx) error {
	var semester models.Semester

	if err := ctx.BodyParser(&semester); err != nil {

		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := controller.CreateSemester(h.MongikClient, &semester)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{
		"semester": semester,
	})
}

func (h *Handler) UpdateSemester(ctx *fiber.Ctx) error {
	var semester models.Semester

	semesterID, err := primitive.ObjectIDFromHex(ctx.Params("semester"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := ctx.BodyParser(&semester); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = controller.UpdateSemester(h.MongikClient, semesterID, &semester)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return ctx.JSON(fiber.Map{"semester": semester})
}
