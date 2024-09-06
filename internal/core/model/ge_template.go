package model

import (
	"time"

	"gorm.io/gorm"
)

type GeTemplate struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string          `gorm:"type:varchar(255);not null" json:"name"`
	RootID    int             `gorm:"type:int;not null;unique" json:"root_id"`
	CreatedAt time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
