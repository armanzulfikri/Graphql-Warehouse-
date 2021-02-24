package service

import (
	"warehouse/args"
	"warehouse/resolve"
	"warehouse/types"

	"github.com/graphql-go/graphql"
)

var mutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			//======================== province mutation =======================
			"createprovince": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Create new province",
				Args:        args.CreateProvinceArgs(),

				Resolve: resolve.CreateProvince,
			},
			"updateprovince": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Update Product",
				Args:        args.UpdateProvinceArgs(),

				Resolve: resolve.UpdateProvince,
			},
			"deleteprovince": &graphql.Field{
				Type:        types.ProvinceType(),
				Description: "Delete Province",
				Args:        args.DeleteProvinceArgs(),
				Resolve:     resolve.DeleteProvince,
			},
			//========================= district mutation ============================
			"createdistrict": &graphql.Field{
				Type:        types.DistrictType(),
				Description: "Create new district",
				Args:        args.CreateDistrictArgs(),

				Resolve: resolve.CreateDistrict,
			},
			"updatedistrict": &graphql.Field{
				Type:        types.DistrictType(),
				Description: "Update District",
				Args:        args.UpdateDistrictsArgs(),

				Resolve: resolve.UpdateDistricts,
			},
			"deletedistrict": &graphql.Field{
				Type:        types.DistrictType(),
				Description: "Delete District",
				Args:        args.DeleteDistrictsArgs(),
				Resolve:     resolve.DeleteDistricts,
			},
		},
	},
)
