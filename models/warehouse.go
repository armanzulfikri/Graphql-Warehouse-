package models

import (
	"time"

	"gorm.io/gorm"
)

//Warehouses entity
type Warehouses struct {
	ID                 uint `gorm:"primarykey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	DistrictsID        uint           `json:"district_id"`
	WarehousesName     string         `json:"warehouse_name"`
	WarehousesCapacity int            `json:"warehouse_capacity"`
	WarehousesType     string         `json:"warehouse_type"`
	WarehousesLocation string         `json:"warehouse_location"`
}
