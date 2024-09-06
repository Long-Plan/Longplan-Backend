package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	CMUITAccount string `gorm:"primary_key"`
	Prename      string
	Firstname    string          `gorm:"not null"`
	Lastname     string          `gorm:"not null"`
	AccountType  int             `gorm:"not null"`
	Organization int             `gorm:"not null"`
	CreatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at"`
}
