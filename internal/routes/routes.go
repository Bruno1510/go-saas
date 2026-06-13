package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/brunoferrero/go-saas/internal/container"
)

func Register(
	app *fiber.App,
	c *container.Container,
) {

	app.Get("/health", func(ctx *fiber.Ctx) error {

		return ctx.JSON(
			fiber.Map{
				"status": "ok",
			},
		)
	})

	auth := app.Group("/auth")

	auth.Post(
		"/register",
		c.AuthHandler.Register,
	)
}