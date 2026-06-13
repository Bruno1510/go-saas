package usecase

import (
	"context"

	"github.com/brunoferrero/go-saas/internal/domain"
	"github.com/brunoferrero/go-saas/internal/infrastructure/security"
	"github.com/brunoferrero/go-saas/internal/repository/interfaces"
	"github.com/brunoferrero/go-saas/internal/usecase/dto"

	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB

	tenantRepo interfaces.TenantRepository
	userRepo   interfaces.UserRepository
}

func NewAuthService(
	db *gorm.DB,
	tenantRepo interfaces.TenantRepository,
	userRepo interfaces.UserRepository,
) *AuthService {

	return &AuthService{
		db:         db,
		tenantRepo: tenantRepo,
		userRepo:   userRepo,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	req dto.RegisterRequest,
) error {

	return s.db.Transaction(func(tx *gorm.DB) error {

		hashedPassword, err :=
			security.HashPassword(req.Password)

		if err != nil {
			return err
		}

		tenant := &domain.Tenant{
			Name: req.CompanyName,
		}

		if err := s.tenantRepo.Create(
			ctx,
			tx,
			tenant,
		); err != nil {
			return err
		}

		user := &domain.User{
			TenantID: tenant.ID,
			Name:     req.Name,
			Email:    req.Email,
			Password: hashedPassword,
			Role:     domain.RoleAdmin,
		}

		if err := s.userRepo.Create(
			ctx,
			tx,
			user,
		); err != nil {
			return err
		}

		return nil
	})
}