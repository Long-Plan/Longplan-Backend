package model

import (
	"time"

	"gorm.io/gorm"
)

type CurriculumCourse struct {
	ID         int             `gorm:"primary_key"`
	CategoryID int             `gorm:"not null"`
	CourseNo   int             `gorm:"not null"`
	Semester   int            
	Year       int             
	CreatedAt  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at"`
}
