package gql

import "github.com/graphql-go/graphql"

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"todos": &graphql.Field{
			Type: graphql.NewList(TodoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				todos := []Todo{
					{Id: "1", Title: "title"},
				}
				return todos, nil
			},
		},
		"todo": &graphql.Field{
			Type: TodoType,
			Args: graphql.FieldConfigArgument{

				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var todo = Todo{
					Id:    p.Args["id"].(string),
					Title: "this is title",
				}
				// Implement your logic to fetch a todo by ID
				return todo, nil
			},
		},
	},
})
