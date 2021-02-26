package args

import "github.com/graphql-go/graphql"

//CreateWarehouseArgs func
func CreateWarehouseArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"district_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"warehouse_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"warehouse_capacity": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"warehouse_type": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"warehouse_location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//UpdateWarehouseArgs func
func UpdateWarehouseArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"district_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"warehouse_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"warehouse_capacity": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"warehouse_type": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"warehouse_location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//DeleteWarehouseArgs func
func DeleteWarehouseArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
