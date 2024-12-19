package middleware

import (
	"fmt"
	"go-fiber/auth"
	"go-fiber/db"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func TodoMiddleware(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	token := c.Get("token")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	userId := auth.VerifyJwtToken(token)
	fmt.Println(userId)

	if userId == nil || userId == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	var userTodos []models.Todo
	rows, err := db.DBPool.Query(c.Context(), "SELECT * FROM todo WHERE userid = $1", userId)

	if err != nil {
		if err.Error() == "no rows in result set" {
			// return c.Status(200).JSON(fiber.Map{
			// 	"todos": []models.Todo{},
			// })
			return c.Next()

		}
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot get todos",
		})
	}
	defer rows.Close()
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.TodoID, &todo.TodoDescription, &todo.TodoPriority, &todo.TodoUserId)
		if err != nil {
			fmt.Println(err)
		}
		userTodos = append(userTodos, todo)
	}
	c.Locals("userTodos", userTodos)
	return c.Next()

}
