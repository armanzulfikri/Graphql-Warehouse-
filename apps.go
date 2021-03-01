package main

import (
	"fmt"
	"net/http"
	"warehouse/config"
	"warehouse/middlewares"
	"warehouse/models"
	"warehouse/seeders"
	"warehouse/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//connect DB
	pgDB := config.Connect()

	//migration
	models.Migrations(pgDB)

	//seeding data
	seeders.SeedProvince(pgDB)
	seeders.SeedDisctrict(pgDB)
	seeders.SeedUser(pgDB)
	seeders.SeedSupplier(pgDB)
	seeders.SeedCategory(pgDB)
	seeders.SeedRack(pgDB)
	seeders.SeedProduct(pgDB)
	seeders.SeedWareHouse(pgDB)

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		type Query struct {
			Query string `json:"query"`
		}

		middlewares.Token = c.Request.Header.Get("Authorization")
		var query Query

		c.Bind(&query)

		result := service.ExecuteQuery(query.Query, service.Schema)
		fmt.Println(result)
		c.JSON(http.StatusAccepted, result)
	})
	r.Run()
}
