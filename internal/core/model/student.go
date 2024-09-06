package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	StudentCode int             `gorm:"primary_key"`
	CreatedAt   time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`
}
