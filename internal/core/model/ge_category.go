package model

import (
	"time"
	
	"gorm.io/gorm"
)

type GeCategory struct {
	ID        int             `gorm:"primary_key;auto_increment" json:"id"`
	NameTh    string          `gorm:"not null" json:"name_th"`
	NameEn    string          `gorm:"not null" json:"name_en"`
	ParentID  int             `json:"parent_id"`
	Kind      CategoryKind    `gorm:"not null" json:"kind"`
	CreatedAt time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
}
