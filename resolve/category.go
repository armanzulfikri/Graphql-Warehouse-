package resolve

import (
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/graphql-go/graphql"
)

//GetCategory func
func GetCategory(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	var category []models.Categories

	dbPG.Find(&category)
	return category, nil

}

//CreateCategory func
func CreateCategory(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var category models.Categories

	category.ID = uint(rand.Intn(100000))
	category.CategoryName = params.Args["category_name"].(string)

	db.Create(&category)
	return category, nil
}

//UpdateCategory func
func UpdateCategory(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	categoryName := params.Args["category_name"].(string)

	var category models.Categories
	db.Model(&category).Where("id =  ?", id).Update("category_name", categoryName)

	return category, nil
}

//DeleteCategory func
func DeleteCategory(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var category models.Categories

	db.Where("id = ?", id).Delete(&category, id)

	return category, nil
}
