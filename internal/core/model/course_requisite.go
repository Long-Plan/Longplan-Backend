package model

import (
	"time"

	"gorm.io/gorm"
)

type CourseRequisite struct {
	ID                 int             `gorm:"primary_key"`
	CurriculumCourseID int             `gorm:"not null"`
	RelatedCourseNo    int             `gorm:"not null"`
	RequisiteType      RequisiteType   `gorm:"not null"`
	CreatedAt          time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt          *gorm.DeletedAt `json:"deleted_at"`
}
