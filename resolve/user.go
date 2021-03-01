package resolve

import (
	"log"
	"math/rand"
	"os"
	"time"
	"warehouse/config"
	"warehouse/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
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

//LoginUser func
func LoginUser(params graphql.ResolveParams) (interface{}, error) {
	var (
		user   models.User
		result interface{}
	)
	db := config.Connect()
	email := params.Args["email"].(string)
	password := params.Args["password"].(string)
	db.Where("email = ?", email).First(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Email", user.Email, "Password Salah")
		result = gin.H{
			"message": "email atau password anda salah",
		}
	} else {
		type authCustomClaims struct {
			Email    string `json:"email"`
			Role     string `json:"role"`
			FullName string `json:"full_name"`
			jwt.StandardClaims
		}
		Claims := &authCustomClaims{
			user.Email,
			user.Role,
			user.FullName,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
		}
		if err := godotenv.Load(".env"); err != nil {
			log.Println("SECRET TOKEN not found, message", err.Error())
			result = gin.H{
				"message": "gagal create token",
				"token":   nil,
			}
		} else {
			sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), Claims)

			if token, err := sign.SignedString([]byte(os.Getenv("SECRET_TOKEN"))); err != nil {
				log.Println("Gagal Create Token, message", err.Error())
				result = gin.H{
					"message": "anda tidak berhasil login",
					"token":   nil,
				}
			} else {
				log.Println("Email", user.Email, "telah login")
				result = gin.H{
					"message": "anda berhasil login",
					"email":   user.Email,
					"token":   token,
				}
			}
		}
	}
	return result, nil
}
