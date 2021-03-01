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
			"login": &graphql.Field{
				Type:        types.Login(),
				Description: "Login User",
				Args:        args.LoginUserArgs(),
				Resolve:     resolve.LoginUser,
			},
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
				Description: "Create New Supplier",
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
			//========================= category mutation ============================
			"createcategory": &graphql.Field{
				Type:        types.CategoriesType(),
				Description: "Create Category",
				Args:        args.CreateCategoryArgs(),

				Resolve: resolve.CreateCategory,
			},
			"updatecategory": &graphql.Field{
				Type:        types.CategoriesType(),
				Description: "Update Category",
				Args:        args.UpdateCategoryArgs(),
				Resolve:     resolve.UpdateCategory,
			},
			"deletecategory": &graphql.Field{
				Type:        types.CategoriesType(),
				Description: "Delete Category",
				Args:        args.DeleteCategoryArgs(),
				Resolve:     resolve.DeleteCategory,
			},
			//========================= racks mutation ============================
			"createracks": &graphql.Field{
				Type:        types.RacksType(),
				Description: "Create New Rack",
				Args:        args.CreateRacksArgs(),
				Resolve:     resolve.CreateRacks,
			},
			"updateracks": &graphql.Field{
				Type:        types.RacksType(),
				Description: "Update Rack",
				Args:        args.UpdateRacksArgs(),
				Resolve:     resolve.UpadateRack,
			},
			"deleteracks": &graphql.Field{
				Type:        types.RacksType(),
				Description: "Delete Rack",
				Args:        args.DeleteRacksArgs(),
				Resolve:     resolve.DeleteRaks,
			},
			//========================= racks mutation ============================
			"createproducts": &graphql.Field{
				Type:        types.ProductType(),
				Description: "Create New Product",
				Args:        args.CreateProducts(),
				Resolve:     resolve.CreateProduct,
			},
			"updateproducts": &graphql.Field{
				Type:        types.ProductType(),
				Description: "Update Products",
				Args:        args.UpdateProductArgs(),
				Resolve:     resolve.UpdateProduct,
			},
			"deleteproduct": &graphql.Field{
				Type:        types.ProductType(),
				Description: "Delete Product",
				Args:        args.DeleteProductArgs(),
				Resolve:     resolve.DeleteProduct,
			},
			//========================= racks mutation ============================
			"createwarehouse": &graphql.Field{
				Type:        types.WarehouseType(),
				Description: "Create New Warehouse",
				Args:        args.CreateWarehouseArgs(),
				Resolve:     resolve.CreateWarehouse,
			},
			"updatewarehouse": &graphql.Field{
				Type:        types.WarehouseType(),
				Description: "Update Warehouse",
				Args:        args.UpdateWarehouseArgs(),
				Resolve:     resolve.UpdateWarehouse,
			},
			"deletewarehouse": &graphql.Field{
				Type:        types.WarehouseType(),
				Description: "Delete Warehouse",
				Args:        args.DeleteWarehouseArgs(),
				Resolve:     resolve.DeleteWarehouse,
			},
		},
	},
)
