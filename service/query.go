package service

import (
	"warehouse/resolve"
	"warehouse/types"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"list_province": &graphql.Field{
				Type:        graphql.NewList(types.ProvinceType()),
				Description: "Get Province List",
				Resolve:     resolve.GetProvince,
			},

			"list_district": &graphql.Field{
				Type:        graphql.NewList(types.DistrictType()),
				Description: "Get District List",
				Resolve:     resolve.GetDistrict,
			},

			"get_province_district": &graphql.Field{
				Type:        graphql.NewList(types.ProvinceType()),
				Description: "Get province and district list",
				Resolve:     resolve.GetProvinceDistrict,
			},
			"get_user": &graphql.Field{
				Type:        graphql.NewList(types.UserType()),
				Description: "Get User List",
				Resolve:     resolve.GetUser,
			},
			"get_supllier": &graphql.Field{
				Type:        graphql.NewList(types.SupplierType()),
				Description: "Get Supplier List",
				Resolve:     resolve.GetSupplier,
			},
		},
	},
)
