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

	var savedUser models.User

	 err := db.DBPool.QueryRow(c.Context(), "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING *", user.Username, user.Password).Scan(&savedUser.ID, &savedUser.Username, &savedUser.Password)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot insert user",
		})
	}

	token, _ := auth.GenerateJwt(user.Username, user.ID)
	return c.Status(201).JSON(fiber.Map{
		"message": "User created",
		"data":    savedUser,
		"token":   token,
	})
}

func LoginHandler(c *fiber.Ctx) error {
	user := new(models.User)

	// Parse the JSON request body
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	// Query the database for the user
	var savedUser models.User
	err := db.DBPool.QueryRow(
		c.Context(),
		"SELECT id, username, password FROM users WHERE username = $1 AND password = $2",
		user.Username,
		user.Password,
	).Scan(&savedUser.ID, &savedUser.Username, &savedUser.Password)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid username or password",
			"error":   err.Error(),
		})
	}

	// Generate JWT token for the user
	token, err := auth.GenerateJwt(savedUser.Username, savedUser.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Cannot generate token",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
