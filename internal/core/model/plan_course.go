package model

import (
	"time"

	"gorm.io/gorm"
)

type PlanCourse struct {
	ID         int `gorm:"primary_key"`
	PlanID     int `gorm:"not null"`
	Semester   int `gorm:"not null"`
	Year       int `gorm:"not null"`
	CourseNo   *int
	CourseType PlanCourseType  `gorm:"not null;default:'Normal'"`
	CreatedAt  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at"`
}
