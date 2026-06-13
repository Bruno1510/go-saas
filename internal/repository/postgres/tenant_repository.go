package postgres

import (
	"context"

	"github.com/brunoferrero/go-saas/internal/domain"
	"gorm.io/gorm"
)

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(
	db *gorm.DB,
) *tenantRepository {
	return &tenantRepository{
		db: db,
	}
}

func (r *tenantRepository) Create(
	ctx context.Context,
	tx *gorm.DB,
	tenant *domain.Tenant,
) error {

	return tx.
		WithContext(ctx).
		Create(tenant).
		Error
}