package handler

import (
	"github.com/FrosTiK-SD/mess-backend/constants"
	"github.com/gofiber/fiber/v2"
)

// File for functionalities specific to students
type StudentFilter struct {
	Department []constants.Department `json:"department,omitempty"`
	Course     []constants.Course     `json:"course,omitempty"`
	StartYear  []int                  `json:"startYear,omitempty"`
	EndYear    []int                  `json:"endYear,omitempty"`
	RollNo     []int                  `json:"rollNo,omitempty"`
}

func (handler *Handler) GetFilteredStudents(ctx *fiber.Ctx) error {
	sf := new(StudentFilter)
	ctx.BodyParser(sf)

	return nil
}
