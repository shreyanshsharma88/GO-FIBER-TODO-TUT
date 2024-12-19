package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber/db"
	"go-fiber/models"
)

func AddTodoHandler(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	err := db.DBPool.QueryRow(c.Context(), "INSERT INTO todo (description , priority , userid) VALUES ($1, $2, $3) RETURNING *", todo.TodoDescription, todo.TodoPriority, todo.TodoUserId).Scan(&todo.TodoID, &todo.TodoDescription, &todo.TodoPriority, &todo.TodoUserId)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "Todo added successfully",
		"todo":    todo,
	})

}

func GetTodosHandler(c *fiber.Ctx) error {
	todos := c.Locals("userTodos")
	return c.JSON(fiber.Map{
		"message": "Get todo",
		"todos":   todos,
	})
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	todoID := c.Params("id")
	todos := c.Locals("userTodos").([]models.Todo)
	for i, todo := range todos {
		if todo.TodoID == todoID {
			break
		}
		if i == len(todos)-1 {
			return c.Status(404).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
	}
	_, err := db.DBPool.Exec(c.Context(), "DELETE FROM todo WHERE id = $1", todoID)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Deleted todo",
	})
}

func UpdateTodoHandler(c *fiber.Ctx) error {
	todoId := c.Params("id")
	userTodos := c.Locals("userTodos").([]models.Todo)
	todo := new(models.Todo)

	for i, todo := range userTodos {
		if todo.TodoID == todoId {
			break
		}
		if i == len(userTodos)-1 {
			return c.Status(404).JSON(fiber.Map{
				"message": "Todo not found",
			})
		}
	}
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	_, err := db.DBPool.Exec(c.Context(), "UPDATE todo SET description = $1, priority = $2 WHERE id = $3", todo.TodoDescription, todo.TodoPriority, todoId)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Add todo",
	})
}
