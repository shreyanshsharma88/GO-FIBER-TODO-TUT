package routes

import (
	"go-fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(api fiber.Router) {
	api.Post("/signup", handlers.UserSignUpHandler)	
	api.Post("/login" , handlers.LoginHandler)

}
