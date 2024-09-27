package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func accessTokenMiddleware(c *fiber.Ctx, db *sql.DB) error {
	authToken := c.Get("Authorization")
	if authToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		}) // Return an error
	}

	refreshToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	_, ok := refreshToken.Claims.(jwt.MapClaims)

	if ok && refreshToken.Valid {
		return c.Next()
	}

	return c.Status(401).JSON(fiber.Map{
		"message": "Unauthorized",
	}) // Return an error
}
