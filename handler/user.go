package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kr/pretty"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) []map[string]interface{} {
	var res map[string]interface{}
	var todos []map[string]interface{}
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	pretty.Println()
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	return todos
}

func (handler *Handler) RegisterStudent(ctx *fiber.Ctx) error {
	test := indexHandler(ctx, handler.DB)
	return ctx.JSON(fiber.Map{"message": test})
	// return nil
}
