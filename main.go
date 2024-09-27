package main

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	SecretKey, err := ReadSecretKeys("public.pem", "private.pem")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	tokenSecretKey = SecretKey

	key, err := publicKeyToString(tokenSecretKey.Public)

	if err != nil {
		panic(err)
	}

	publicKey = key

	databaseConfiguration, err := ReadENV()
	if err != nil {
		panic(err)
	}

	db, err := InitDatabase(databaseConfiguration)

	if err != nil {
		panic(err)
	}

	configuration, err := ReadConfig(db)

	if err != nil {
		panic(err)
	}

	config = configuration

	app := fiber.New()

	app.Post("/api/refresh", func(ctx *fiber.Ctx) error {
		return RefreshRoute(ctx, db)
	})

	app.Post("/api/registration", func(ctx *fiber.Ctx) error {
		return RegistrationRoute(ctx, db)
	})

	app.Post("/api/login", func(ctx *fiber.Ctx) error {
		return LoginRoute(ctx, db)
	})

	app.Post("/api/publicKey", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"publicKey": publicKey,
		})
	})

	log.Fatal(app.Listen(":3000"))

}
