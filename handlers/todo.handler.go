package handlers

import (

	"github.com/gofiber/fiber/v2"
)

func AddTodoHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Add todo",
	})

}

func GetTodosHandler(c *fiber.Ctx) error {
    todos := c.Locals("userTodos")
    return c.JSON(fiber.Map{
		"message": "Add todo",
        "todos": todos,
	})
}

func DeleteTodoHandler(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
		"message": "Add todo",
	})
}

func UpdateTodoHandler(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{
		"message": "Add todo",
	})
}
