package postgres

import (
	"context"

	"github.com/brunoferrero/go-saas/internal/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	tx *gorm.DB,
	user *domain.User,
) error {

	return tx.
		WithContext(ctx).
		Create(user).
		Error
}

func (r *userRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {

	var user domain.User

	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	return &user, err
}