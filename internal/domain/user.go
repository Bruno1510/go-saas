package domain

import "gorm.io/gorm"

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleAdmin UserRole = "ADMIN"
	RoleUser  UserRole = "USER"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`

	TenantID uuid.UUID
	Tenant   Tenant

	Name     string
	Email    string `gorm:"unique"`
	Password string

	Role UserRole

	CreatedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}