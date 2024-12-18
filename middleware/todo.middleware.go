package middleware

import (
	"fmt"
	"go-fiber/auth"
	"go-fiber/db"

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

	if userId == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	userTodos , err := db.DBPool.Query(c.Context(), "SELECT * FROM todo WHERE user_id = $1", userId)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot get todos",
		})
	}
	c.Locals("userTodos", userTodos)
	return c.Next()

}
