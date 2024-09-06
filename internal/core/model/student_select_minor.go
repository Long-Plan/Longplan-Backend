package model

import (
	"time"

	"gorm.io/gorm"
)

type StudentSelectMinor struct {
	ID           int             `gorm:"primary_key"`
	StudentCode  int             `gorm:"not null"`
	CurriculumID int             `gorm:"unique;not null"`
	CourseNo     int             `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at"`
}
