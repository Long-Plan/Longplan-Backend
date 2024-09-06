package model

import (
	"time"

	"gorm.io/gorm"
)

type Curriculum struct {
	ID           int             `gorm:"primary_key"`
	Code         int             `gorm:"unique"`
	NameTh       string          `gorm:"unique;not null"`
	NameEn       string          `gorm:"unique"`
	FSCategoryID int             `gorm:"unique;not null"`
	GETemplateID int             `gorm:"not null"`
	GECategoryID int             `gorm:"unique;not null"`
	FECategoryID int             `gorm:"unique;not null"`
	CreatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at"`
}
