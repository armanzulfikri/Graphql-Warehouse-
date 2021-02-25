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
			//========================= User mutation ============================
			"register": &graphql.Field{
				Type:        types.UserType(),
				Description: "Register New User",
				Args:        args.CreateUserArgs(),

				Resolve: resolve.CreateUser,
			},
			"updateuser": &graphql.Field{
				Type:        types.UserType(),
				Description: "Update User",
				Args:        args.UpdateUserArgs(),

				Resolve: resolve.UpdateUserArgs,
			},
			"deleteuser": &graphql.Field{
				Type:        types.UserType(),
				Description: "Delete User",
				Args:        args.DeleteUserArgs(),
				Resolve:     resolve.DeleteUserArgs,
			},
			//========================= supplier mutation ============================
			"createsupplier": &graphql.Field{
				Type:        types.SupplierType(),
				Description: "Register New Supplier",
				Args:        args.CreateSupplierArgs(),

				Resolve: resolve.CreateSuppliers,
			},
			"updatesupplier": &graphql.Field{
				Type:        types.SupplierType(),
				Description: "Update Supplier",
				Args:        args.UpdateSUpplierArgs(),

				Resolve: resolve.UpdateSupplier,
			},
			"deletesupplier": &graphql.Field{
				Type:        types.SupplierType(),
				Description: "Delete Supplier",
				Args:        args.DeleteSupplierArgs(),
				Resolve:     resolve.DeleteSupplier,
			},
		},
	},
)
