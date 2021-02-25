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
