package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/brunoferrero/go-saas/internal/config"
	"github.com/brunoferrero/go-saas/internal/infrastructure/database"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)

	if err != nil {
		log.Fatal(err)
	}

	_ = db

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	log.Fatal(
		app.Listen(":" + cfg.AppPort),
	)
}