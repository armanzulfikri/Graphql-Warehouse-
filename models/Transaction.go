package models

import (
	"time"

	"gorm.io/gorm"
)

// Transactions model / product in warehouse
type Transactions struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProductID    uint `json:"product_id"`
	UserID       uint `json:"user_id"`
	RackID       uint `json:"rack_id"`
	ProductStock int  `json:"product_stock"`
}
