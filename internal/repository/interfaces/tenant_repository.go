package interfaces

import (
	"context"

	"github.com/brunoferrero/go-saas/internal/domain"
	"gorm.io/gorm"
)

type TenantRepository interface {
	Create(
		ctx context.Context,
		tx *gorm.DB,
		tenant *domain.Tenant,
	) error
}