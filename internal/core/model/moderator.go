package model

import (
	"time"

	"gorm.io/gorm"
)

type Moderator struct {
	CMUITAccount string          `gorm:"primary_key"`
	Role         ModRoles        `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at"`
}
