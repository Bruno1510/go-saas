package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/brunoferrero/go-saas/internal/usecase"
	"github.com/brunoferrero/go-saas/internal/usecase/dto"
)

type AuthHandler struct {
	authService *usecase.AuthService
}

func NewAuthHandler(
	authService *usecase.AuthService,
) *AuthHandler {

	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(
	c *fiber.Ctx,
) error {

	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "invalid body",
			},
		)
	}

	if err := h.authService.Register(
		c.Context(),
		req,
	); err != nil {

		return c.Status(500).JSON(
			fiber.Map{
				"message": err.Error(),
			},
		)
	}

	return c.Status(201).JSON(
		fiber.Map{
			"message": "user created",
		},
	)
}