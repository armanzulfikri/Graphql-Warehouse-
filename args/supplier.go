package args

import "github.com/graphql-go/graphql"

//CreateSupplierArgs func
func CreateSupplierArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"supplier_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"address": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"telepon": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//UpdateSUpplierArgs func
func UpdateSUpplierArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"supplier_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"address": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"telepon": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//DeleteSupplierArgs func
func DeleteSupplierArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
