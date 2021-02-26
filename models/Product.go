package models

import (
	"time"

	"gorm.io/gorm"
)

// Products true table
type Products struct {
	ID              uint   `gorm:"primarykey"`
	UserID          uint   `json:"user_id"`
	CategoryID      uint   `json:"category_id"`
	SupplierID      uint   `json:"supplier_id"`
	ProductName     string `json:"product_name"`
	ProductImageURL string `json:"product_image_url"`
	Description     string `json:"description"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
