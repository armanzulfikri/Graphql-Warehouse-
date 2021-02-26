package resolve

import (
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/graphql-go/graphql"
)

//GetProduct func
func GetProduct(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	var product []models.Products

	dbPG.Find(&product)
	return product, nil

}

//CreateProduct func
func CreateProduct(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var product models.Products

	product.ID = uint(rand.Intn(100000))
	product.UserID = uint(params.Args["user_id"].(int))
	product.CategoryID = uint(params.Args["category_id"].(int))
	product.SupplierID = uint(params.Args["supplier_id"].(int))
	product.ProductName = params.Args["product_name"].(string)
	product.ProductImageURL = params.Args["product_image_url"].(string)
	product.Description = params.Args["description"].(string)

	db.Create(&product)
	return product, nil
}

//UpdateProduct func
func UpdateProduct(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	userID := uint(params.Args["user_id"].(int))
	categoryID := uint(params.Args["category_id"].(int))
	supplierID := uint(params.Args["supplier_id"].(int))
	productName := params.Args["product_name"].(string)
	productImageURL := params.Args["product_image_url"].(string)
	description := params.Args["description"].(string)

	var product models.Products
	db.Model(&product).Where("id =  ?", id).Updates(models.Products{
		UserID:          userID,
		CategoryID:      categoryID,
		SupplierID:      supplierID,
		ProductName:     productName,
		ProductImageURL: productImageURL,
		Description:     description,
	})

	return product, nil
}

//DeleteProduct func
func DeleteProduct(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var product models.Products

	db.Where("id = ?", id).Delete(&product, id)

	return product, nil
}
