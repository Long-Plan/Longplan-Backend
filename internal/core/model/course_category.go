package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseCategory struct {
	ID               int          `gorm:"primary_key"`
	NameTh           string       `gorm:"not null"`
	NameEn           string       `gorm:"not null"`
	Credit           int          `gorm:"not null;default:0"`
	Kind             CategoryKind `gorm:"not null"`
	SelfCategorizeID *int
	Note             string
	ParentID         *int
	CrossCategoryID  *int
	CreatedAt        time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt        *gorm.DeletedAt `json:"deleted_at"`
}
