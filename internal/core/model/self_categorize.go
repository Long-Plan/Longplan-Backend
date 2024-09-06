package model

import (
	"time"

	"gorm.io/gorm"
)

type SelfCategorize struct {
	ID           int             `gorm:"primary_key"`
	CurriculumID int             `gorm:"not null"`
	NameTh       string          `gorm:"not null"`
	NameEn       string          `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at"`
}
