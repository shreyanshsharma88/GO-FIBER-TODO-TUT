package handlers

import (
	"go-fiber/auth"
	"go-fiber/db"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func UserSignUpHandler(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	data, err := db.DBPool.Query(c.Context(), "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING *", user.Username, user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot insert user",
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "User created",
		"data":    data,
	})
}

func LoginHandler (c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	_, err := db.DBPool.Query( c.Context() , "SELECT * FROM users WHERE username = $1 AND password = $2", user.Username, user.Password)	
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot get user",
		})
	}
	token, err := auth.GenerateJwt(user.Username, user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot generate token",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
