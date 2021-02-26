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
			"get_category": &graphql.Field{
				Type:        graphql.NewList(types.CategoriesType()),
				Description: "Get Category List",
				Resolve:     resolve.GetCategory,
			},
			"get_rack": &graphql.Field{
				Type:        graphql.NewList(types.RacksType()),
				Description: "Get Rack List",
				Resolve:     resolve.GetRack,
			},
			"get_product": &graphql.Field{
				Type:        graphql.NewList(types.ProductType()),
				Description: "Get Product List",
				Resolve:     resolve.GetProduct,
			},
			"get_warehouse": &graphql.Field{
				Type:        graphql.NewList(types.WarehouseType()),
				Description: "Get Warehouse List",
				Resolve:     resolve.GetWarehouse,
			},
		},
	},
)
