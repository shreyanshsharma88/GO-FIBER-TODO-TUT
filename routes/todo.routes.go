package routes

import "github.com/gofiber/fiber/v2"

func RegisterTodoRoutes (api fiber.Router){
	api.Post("/")
	api.Get("/")
	api.Delete("/:id")
	api.Put("/:id")
}