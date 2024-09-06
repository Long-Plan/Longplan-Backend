package model

import (
	"time"

	"gorm.io/gorm"
)

type AccountType struct {
	ID        int            `gorm:"primary_key"`
	NameTh    string          `gorm:"unique;not null"`
	NameEn    string          `gorm:"unique;not null"`
	CreatedAt time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
