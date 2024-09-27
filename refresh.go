package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func RefreshRoute(c *fiber.Ctx, db *sql.DB) error {
	receivedRefreshToken := c.FormValue("refreshToken")
	refreshToken, err := jwt.Parse(receivedRefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenSecretKey.Public, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)

	if ok && refreshToken.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid refresh token"})
		}

		expirationTime := time.Unix(int64(exp), 0)
		if time.Now().After(expirationTime) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Time is expired"})
		}

		email, ok := claims["sub"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email"})
		}

		row, err := db.Query("SELECT * FROM users WHERE email = $1", email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email"})
		}

		user := new(User)

		row.Next()
		err = row.Scan(&user.id, &user.Email, &user.Password, &user.RefreshToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": "Internal server error",
				})
		}

		if email == user.Email {
			accessToken, err := generateToken(user, tokenSecretKey.Private, config.AccessTokenTime)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"error": "Internal server error",
					})
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"accessToken": accessToken,
			})
		}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid refresh token"})
}
