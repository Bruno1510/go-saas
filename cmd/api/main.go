package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/brunoferrero/go-saas/internal/config"
	"github.com/brunoferrero/go-saas/internal/container"
	"github.com/brunoferrero/go-saas/internal/infrastructure/database"
	"github.com/brunoferrero/go-saas/internal/routes"
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

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	c := container.Build(db)

	routes.Register(
		app,
		c,
	)

	log.Println("Database connected")

	log.Fatal(
		app.Listen(":" + cfg.AppPort),
	)
}