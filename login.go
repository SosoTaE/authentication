package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func LoginRoute(c *fiber.Ctx, db *sql.DB) error {
	p := new(User)

	if err := c.BodyParser(p); err != nil {
		return c.JSON(fiber.Map{
			"message": "Error while parsing body",
		})
	}

	row, err := db.Query("SELECT * FROM users WHERE email = $1", p.Email)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Error while querying user",
		})
	}

	var u User
	row.Next()
	err = row.Scan(&u.id, &u.Email, &u.Password, &u.RefreshToken)
	fmt.Println(err)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Error while querying user",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	tokens, err := generateTokens(p, tokenSecretKey, config)
	fmt.Println(err)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Error while generating tokens",
		})
	}

	_, err = db.Exec("UPDATE users SET \"refreshToken\" = $1 WHERE email = $2", tokens.RefreshToken, u.Email)
	fmt.Println(err)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Error while updating tokens",
		})
	}

	return c.JSON(tokens)
}
