package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Migrations func
func Migrations(db *gorm.DB) {
	var checkTableProvinces, checkTableDistricts, checkTableUserss,
		checkTabelSuppliers, checkTableCategory, checkTableRack,
		checkTableProducts, checkTableWarehouse bool

	db.Migrator().DropTable(&Provinces{})
	db.Migrator().DropTable(&Districts{})
	db.Migrator().DropTable(&User{})
	db.Migrator().DropTable(&Suppliers{})
	db.Migrator().DropTable(&Categories{})
	db.Migrator().DropTable(&Racks{})
	db.Migrator().DropTable(&Products{})
	db.Migrator().DropTable(&Warehouses{})

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

	checkTableCategory = db.Migrator().HasTable(&Categories{})
	if !checkTableCategory {
		db.Migrator().CreateTable(&Categories{})
		fmt.Println("Create Table Category")
	}

	checkTableRack = db.Migrator().HasTable(&Racks{})
	if !checkTableRack {
		db.Migrator().CreateTable(&Racks{})
		fmt.Println("Create Table Racks")
	}

	checkTableProducts = db.Migrator().HasTable(&Products{})
	if !checkTableProducts {
		db.Migrator().CreateTable(&Products{})
		fmt.Println("Create Table Product")
	}

	checkTableWarehouse = db.Migrator().HasTable(&Warehouses{})
	if !checkTableWarehouse {
		db.Migrator().CreateTable(&Warehouses{})
		fmt.Println("Create Table Warehouse")
	}

}
