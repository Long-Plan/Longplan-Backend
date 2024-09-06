package model

import (
	"time"

	"gorm.io/gorm"
)

type Plan struct {
	ID          int             `gorm:"primary_key"`
	Name        string          `gorm:"not null"`
	StudentCode int             `gorm:"not null"`
	CreatedAt   time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`
}
