package container

import (
	httpHandler "github.com/brunoferrero/go-saas/internal/handler/http"
	"github.com/brunoferrero/go-saas/internal/usecase"
	"github.com/brunoferrero/go-saas/internal/repository/postgres"

	"gorm.io/gorm"
)

type Container struct {
	AuthHandler *httpHandler.AuthHandler
}

func Build(
	db *gorm.DB,
) *Container {

	tenantRepo := postgres.NewTenantRepository(db)

	userRepo := postgres.NewUserRepository(db)

	authService := usecase.NewAuthService(
		db,
		tenantRepo,
		userRepo,
	)

	authHandler := httpHandler.NewAuthHandler(
		authService,
	)

	return &Container{
		AuthHandler: authHandler,
	}
}