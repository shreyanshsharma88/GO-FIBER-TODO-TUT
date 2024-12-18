package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber/db"
	"go-fiber/routes"
)

func main() {
	app := fiber.New()

	db.ConnectDB()

	app.Get("/health", func(c *fiber.Ctx) error {
		user, err := db.DBPool.Query(c.Context(), "SELECT * FROM users")
		if err != nil {
			return c.Status(500).SendString("Cannot get users")
		}
		return c.JSON(fiber.Map{
			"users": user, 
		})
	})

	userRouter := app.Group("/api/user")
	routes.RegisterUserRoutes(userRouter)

	todoRouter := app.Group("/api/todo")
	routes.RegisterTodoRoutes(todoRouter)

	fmt.Println("Server is running on port :8080")
	app.Listen(":8080")



}