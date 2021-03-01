package types

import (
	"github.com/graphql-go/graphql"
)

//ProvinceType func
func ProvinceType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//DistrictType func
func DistrictType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"province_id": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}

//UserType func
func UserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "District",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"full_name": &graphql.Field{
					Type: graphql.String,
				},
				"gender": &graphql.Field{
					Type: graphql.String,
				},
				"birth_date": &graphql.Field{
					Type: graphql.String,
				},
				"birth_place": &graphql.Field{
					Type: graphql.String,
				},
				"role": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//SupplierType func
func SupplierType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Suppliers",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"supplier_name": &graphql.Field{
					Type: graphql.String,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"telepon": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//CategoriesType func
func CategoriesType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Categories",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"category_name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//RacksType func
func RacksType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Raks",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"warehouse_id": &graphql.Field{
					Type: graphql.Int,
				},
				"category_id": &graphql.Field{
					Type: graphql.Int,
				},
				"rack_code": &graphql.Field{
					Type: graphql.String,
				},
				"rack_capacity": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}

//ProductType func
func ProductType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Products",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"user_id": &graphql.Field{
					Type: graphql.Int,
				},
				"category_id": &graphql.Field{
					Type: graphql.Int,
				},
				"supplier_id": &graphql.Field{
					Type: graphql.Int,
				},
				"product_name": &graphql.Field{
					Type: graphql.String,
				},
				"product_image_url": &graphql.Field{
					Type: graphql.String,
				},
				"description": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//WarehouseType func
func WarehouseType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Warehouses",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"district_id": &graphql.Field{
					Type: graphql.Int,
				},
				"warehouse_name": &graphql.Field{
					Type: graphql.String,
				},
				"warehouse_capacity": &graphql.Field{
					Type: graphql.Int,
				},
				"warehouse_type": &graphql.Field{
					Type: graphql.String,
				},
				"warehouse_location": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//Login func
func Login() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Login",
			Fields: graphql.Fields{
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"token": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}

//TransactionsType func
func TransactionsType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Transaction",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"product_id": &graphql.Field{
					Type: graphql.Int,
				},
				"user_id": &graphql.Field{
					Type: graphql.Int,
				},
				"rack_id": &graphql.Field{
					Type: graphql.Int,
				},
				"product_stock": &graphql.Field{
					Type: graphql.Int,
				},
			},
		},
	)
}
