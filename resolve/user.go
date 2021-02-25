package resolve

import (
	"log"
	"math/rand"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

//GetUser func
func GetUser(params graphql.ResolveParams) (interface{}, error) {
	dbPG := config.Connect()

	var user []models.User

	dbPG.Find(&user)
	return user, nil

}

//CreateUser func
func CreateUser(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	rand.Seed(time.Now().UnixNano())
	var user models.User

	user.ID = uint(rand.Intn(100000))
	user.FullName = params.Args["full_name"].(string)
	user.Gender = params.Args["gender"].(string)
	user.BirthPlace = params.Args["birth_place"].(string)
	user.BirthDate = params.Args["birth_date"].(string)
	user.Role = params.Args["role"].(string)
	user.Email = params.Args["email"].(string)
	user.Password = params.Args["password"].(string)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error -> ", err.Error())
	}
	user.Password = string(hash)
	db.Create(&user)
	return user, nil
}

//UpdateUserArgs func
func UpdateUserArgs(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))
	fullName := params.Args["full_name"].(string)
	gender := params.Args["full_name"].(string)
	birthDate := params.Args["full_name"].(string)
	birthPlace := params.Args["full_name"].(string)
	role := params.Args["full_name"].(string)
	email := params.Args["full_name"].(string)

	var user models.User
	db.Model(&user).Where("id =  ?", id).Updates(models.User{
		FullName:   fullName,
		Gender:     gender,
		BirthDate:  birthDate,
		BirthPlace: birthPlace,
		Role:       role,
		Email:      email,
	})

	return user, nil
}

//DeleteUserArgs func
func DeleteUserArgs(params graphql.ResolveParams) (interface{}, error) {
	db := config.Connect()
	id := uint(params.Args["id"].(int))

	var user models.User

	db.Where("id = ?", id).Delete(&user, id)

	return user, nil
}
