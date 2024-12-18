package routes

import (
	"go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterTodoRoutes (api fiber.Router){
	api.Get("/", handlers.GetTodosHandler)
	// api.Post("/")
	// api.Delete("/:id")
	// api.Put("/:id")
}