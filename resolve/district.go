package resolve

import (
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/middlewares"
	"warehouse/models"

	"github.com/graphql-go/graphql"
)

//GetDistrict func
func GetDistrict(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	type Districts struct {
		Name       string `json:"name"`
		ID         uint   `gorm:"primarykey" json:"id"`
		ProvinceID uint   `gorm:"foreignkey" json:"province_id"`
	}

	var district []Districts
	dbPG.Find(&district)
	return district, nil
}

//CreateDistrict func
func CreateDistrict(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var districts models.Districts

	verifToken, err := middlewares.VerifyToken(middlewares.Token)

	if err != nil {
		return nil, err
	}
	if verifToken["role"] != "admin" && verifToken["role"] != "entry" {

		return nil, err
	}

	districts.ID = uint(rand.Intn(100000))
	districts.ProvinceID = uint(params.Args["province_id"].(int))
	districts.Name = params.Args["name"].(string)

	db.Create(&districts)
	return districts, nil
}

//UpdateDistricts func
func UpdateDistricts(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	name := params.Args["name"].(string)
	provinceID := uint(params.Args["province_id"].(int))

	var districts models.Districts

	db.Model(&districts).Where("id = ?", id).Updates(models.Districts{Name: name, ProvinceID: provinceID})
	return districts, nil
}

//DeleteDistricts func
func DeleteDistricts(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var districts models.Districts

	db.Delete(&districts, id)

	return districts, nil
}
