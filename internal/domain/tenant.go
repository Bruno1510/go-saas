package domain

import "gorm.io/gorm"

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	CreatedAt time.Time
}

func (t *Tenant) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New()
	return nil
}