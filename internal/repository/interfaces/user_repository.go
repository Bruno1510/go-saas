package interfaces

import (
	"context"

	"github.com/brunoferrero/go-saas/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(
		ctx context.Context,
		tx *gorm.DB,
		user *domain.User,
	) error

	FindByEmail(
		ctx context.Context,
		email string,
	) (*domain.User, error)
}