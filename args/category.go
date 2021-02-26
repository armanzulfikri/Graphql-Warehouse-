package args

import "github.com/graphql-go/graphql"

//CreateCategoryArgs func
func CreateCategoryArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"category_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//UpdateCategoryArgs func
func UpdateCategoryArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"category_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//DeleteCategoryArgs func
func DeleteCategoryArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
