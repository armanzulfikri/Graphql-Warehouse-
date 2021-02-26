package resolve

import (
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/graphql-go/graphql"
)

//GetWarehouse func
func GetWarehouse(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	var warehouse []models.Warehouses

	dbPG.Find(&warehouse)
	return warehouse, nil

}

//CreateWarehouse func
func CreateWarehouse(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var warehouse models.Warehouses

	warehouse.ID = uint(rand.Intn(100000))
	warehouse.DistrictsID = uint(params.Args["district_id"].(int))
	warehouse.WarehousesName = params.Args["warehouse_name"].(string)
	warehouse.WarehousesCapacity = params.Args["warehouse_capacity"].(int)
	warehouse.WarehousesType = params.Args["warehouse_type"].(string)
	warehouse.WarehousesLocation = params.Args["warehouse_location"].(string)

	db.Create(&warehouse)
	return warehouse, nil
}

//UpdateWarehouse func
func UpdateWarehouse(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	districtID := uint(params.Args["district_id"].(int))
	warehouseName := params.Args["warehouse_name"].(string)
	warehousesCapacity := params.Args["warehouse_capacity"].(int)
	warehouseType := params.Args["warehouse_type"].(string)
	warehousesLocation := params.Args["warehouse_location"].(string)

	var warehouse models.Warehouses
	db.Model(&warehouse).Where("id =  ?", id).Updates(models.Warehouses{
		DistrictsID:        districtID,
		WarehousesName:     warehouseName,
		WarehousesCapacity: warehousesCapacity,
		WarehousesType:     warehouseType,
		WarehousesLocation: warehousesLocation,
	})

	return warehouse, nil
}

//DeleteWarehouse func
func DeleteWarehouse(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var warehouse models.Warehouses

	db.Where("id = ?", id).Delete(&warehouse, id)

	return warehouse, nil
}
