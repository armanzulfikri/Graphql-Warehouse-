package resolve

import (
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/graphql-go/graphql"
)

//GetSupplier func
func GetSupplier(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	var supplier []models.Suppliers

	dbPG.Find(&supplier)
	return supplier, nil

}

//CreateSuppliers func
func CreateSuppliers(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var supplier models.Suppliers

	supplier.ID = uint(rand.Intn(100000))
	supplier.SupplierName = params.Args["supplier_name"].(string)
	supplier.Address = params.Args["address"].(string)
	supplier.Telepon = params.Args["telepon"].(string)

	db.Create(&supplier)
	return supplier, nil
}

//UpdateSupplier func
func UpdateSupplier(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	supplierName := params.Args["supplier_name"].(string)
	address := params.Args["address"].(string)
	telepon := params.Args["telepon"].(string)

	var supplier models.Suppliers
	db.Model(&supplier).Where("id =  ?", id).Updates(models.Suppliers{
		SupplierName: supplierName,
		Address:      address,
		Telepon:      telepon,
	})

	return supplier, nil
}

//DeleteSupplier func
func DeleteSupplier(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var supplier models.Suppliers

	db.Where("id = ?", id).Delete(&supplier, id)

	return supplier, nil
}
