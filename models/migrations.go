package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var checkTableProvinces, checkTableDistricts, checkTableUserss,
		checkTabelSuppliers bool

	db.Migrator().DropTable(&Provinces{})
	db.Migrator().DropTable(&Districts{})
	db.Migrator().DropTable(&User{})
	db.Migrator().DropTable(&Suppliers{})

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

	checkTableUserss = db.Migrator().HasTable(&User{})
	if !checkTableUserss {
		db.Migrator().CreateTable(&User{})
		fmt.Println("Create Table User ")
	}

	checkTabelSuppliers = db.Migrator().HasTable(&Suppliers{})
	if !checkTabelSuppliers {
		db.Migrator().CreateTable(&Suppliers{})
		fmt.Println("Create Table Suppliers")
	}

}
