package routes

import (
	"go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterTodoRoutes (api fiber.Router){
	api.Get("/", handlers.GetTodosHandler)
	api.Post("/", handlers.AddTodoHandler)
	api.Delete("/:id", handlers.DeleteTodoHandler)
	api.Put("/:id", handlers.UpdateTodoHandler)
}