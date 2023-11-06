package gql

import "github.com/graphql-go/graphql"

type Todo struct {
	Id    string
	Title string
}

var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createTodo": &graphql.Field{
			Type: TodoType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"input": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewInputObject(graphql.InputObjectConfig{
						Name: "CreateTodoInput",
						Fields: graphql.InputObjectConfigFieldMap{
							"title": &graphql.InputObjectFieldConfig{
								Type: graphql.NewNonNull(graphql.String),
							},
							// Add more fields as needed
						},
					})), // Use the input type as the argument
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rq := p.Args["input"]
				var todo = Todo{
					Id:    p.Args["id"].(string),
					Title: rq.(map[string]interface{})["title"].(string),
				}
				// Implement your logic to create a new todo
				return todo, nil

			},
		},

		"deleteTodo": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Implement your logic to delete a todo
				return true, nil
			},
		},
	},
})
