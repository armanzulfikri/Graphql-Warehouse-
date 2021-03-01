package models

import (
	"time"

	"gorm.io/gorm"
)

// Districts model
type Districts struct {
	ID         uint   `gorm:"primarykey"`
	Name       string `json:"name"`
	ProvinceID uint   `json:"province_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
