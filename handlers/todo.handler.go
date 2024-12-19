package handlers

import (
	"fmt"
	"go-fiber/db"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func AddTodoHandler(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	err := db.DBPool.QueryRow(c.Context(),"INSERT INTO todo (description , priority , userid) VALUES ($1, $2, $3) RETURNING *", todo.TodoDescription, todo.TodoPriority, todo.TodoUserId).Scan(&todo.TodoID, &todo.TodoDescription, &todo.TodoPriority, &todo.TodoUserId)	
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Todo added successfully",
		"todo": todo,
	})

}

func GetTodosHandler(c *fiber.Ctx) error {
    todos := c.Locals("userTodos")
    return c.JSON(fiber.Map{
		"message": "Get todo",
        "todos": todos,
	})
}

func DeleteTodoHandler(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
		"message": "Delete todo",
	})
}

func UpdateTodoHandler(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
		"message": "Add todo",
	})
}
