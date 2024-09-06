package model

import (
	"time"

	"gorm.io/gorm"
)

type DetailCourse struct {
	CourseNo     int    			`gorm:"primary_key"`
	TitleLongTh  string 			`gorm:"unique;not null"`
	TitleLongEn  string 			`gorm:"unique;not null"`
	TitleShortEn string
	CourseDescTh string
	CourseDescEn string
	Credit       int 				`gorm:"not null;default:0"`
	Prerequisite string
	CreatedAt    time.Time       	`gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       	`gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    *gorm.DeletedAt 	`json:"deleted_at"`
}
