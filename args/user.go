package args

import "github.com/graphql-go/graphql"

//CreateUserArgs func
func CreateUserArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"full_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"gender": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_place": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"role": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//UpdateUserArgs func
func UpdateUserArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"full_name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"gender": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_date": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"birth_place": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"role": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	}
}

//DeleteUserArgs func
func DeleteUserArgs() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	}
}
