package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var checkTableProvinces, checkTableDistricts bool

	db.Migrator().DropTable(&Provinces{})
	db.Migrator().DropTable(&Districts{})

	checkTableProvinces = db.Migrator().HasTable(&Provinces{})
	if !checkTableProvinces {
		db.Migrator().CreateTable(&Provinces{})
		fmt.Println("Create Table Provinces")
	}

	checkTableDistricts = db.Migrator().HasTable(&Districts{})
	if !checkTableDistricts {
		db.Migrator().CreateTable(&Districts{})
		fmt.Println("Create Table Districts")
	}

}
