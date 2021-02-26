package args

import "github.com/graphql-go/graphql"

//CreateRacksArgs func
func CreateRacksArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"warehouse_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"category_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"rack_code": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"rack_capacity": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

//UpdateRacksArgs func
func UpdateRacksArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"warehouse_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"category_id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"rack_code": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"rack_capacity": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}

//DeleteRacksArgs func
func DeleteRacksArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
