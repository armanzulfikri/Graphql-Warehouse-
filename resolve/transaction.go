package resolve

import (
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/graphql-go/graphql"
)

//GetTransaction func
func GetTransaction(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	var transactions []models.Transactions

	dbPG.Find(&transactions)
	return transactions, nil

}

//CreateTransaction func
func CreateTransaction(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var transactions models.Transactions

	transactions.ID = uint(rand.Intn(100000))
	transactions.ProductID = uint(params.Args["product_id"].(int))
	transactions.UserID = uint(params.Args["user_id"].(int))
	transactions.RackID = uint(params.Args["rack_id"].(int))
	transactions.ProductStock = params.Args["product_stock"].(int)

	db.Create(&transactions)
	return transactions, nil
}

//UpdateTransaction func
func UpdateTransaction(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	productID := uint(params.Args["product_id"].(int))
	userID := uint(params.Args["user_id"].(int))
	rackID := uint(params.Args["rack_id"].(int))
	productStock := params.Args["product_stock"].(int)

	var transactions models.Transactions
	db.Model(&transactions).Where("id =  ?", id).Updates(models.Transactions{
		ProductID:    productID,
		UserID:       userID,
		RackID:       rackID,
		ProductStock: productStock,
	})

	return transactions, nil
}

//DeleteTransaction func
func DeleteTransaction(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var transactions models.Transactions

	db.Where("id = ?", id).Delete(&transactions, id)

	return transactions, nil
}
