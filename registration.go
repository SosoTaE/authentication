package main

import (
	"database/sql"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationRoute(c *fiber.Ctx, db *sql.DB) error {
	p := new(User)

	if err := c.BodyParser(p); err != nil {
		return c.JSON(fiber.Map{
			"message": "Error while parsing body",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error while hashing password",
		})
	}

	_, err = db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", p.Email, string(hashedPassword))
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "User registration failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registration successful",
	})

}
