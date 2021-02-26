package resolve

import (
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/graphql-go/graphql"
)

//GetRack func
func GetRack(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	var rack []models.Racks

	dbPG.Find(&rack)
	return rack, nil

}

//CreateRacks func
func CreateRacks(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var rack models.Racks

	rack.ID = uint(rand.Intn(100000))
	rack.WarehouseID = uint(params.Args["warehouse_id"].(int))
	rack.CategoryID = uint(params.Args["category_id"].(int))
	rack.RackCode = params.Args["rack_code"].(string)
	rack.RackCapacity = params.Args["rack_capacity"].(int)

	db.Create(&rack)
	return rack, nil
}

//UpadateRack func
func UpadateRack(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	warehouseID := uint(params.Args["warehouse_id"].(int))
	categoryID := uint(params.Args["category_id"].(int))
	rackCode := params.Args["rack_code"].(string)
	rackCapacity := params.Args["rack_capacity"].(int)

	var rack models.Racks
	db.Model(&rack).Where("id =  ?", id).Updates(models.Racks{
		WarehouseID:  warehouseID,
		CategoryID:   categoryID,
		RackCode:     rackCode,
		RackCapacity: rackCapacity,
	})

	return rack, nil
}

//DeleteRaks func
func DeleteRaks(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var rack models.Racks

	db.Where("id = ?", id).Delete(&rack, id)

	return rack, nil
}
