package todo

import (
	"github.com/graphql-go/graphql"
)

type Query struct {
	types Types
}

func NewQuery(types Types) Query {
	return Query{types: types}
}

func (q Query) GetAllTodoResolver(p graphql.ResolveParams) (interface{}, error) {
	todos := []Todo{
		{
			Id:          1,
			Title:       "todo 1",
			Description: "todo 1 desc",
			IsCompleted: true,
		},
		{
			Id:          2,
			Title:       "todo 2",
			Description: "todo 2 desc",
			IsCompleted: false,
		},
	}
	return todos, nil
}

func (q Query) GetAllTodoField() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(q.types.TodoType("AllTodo")),
		// // we can add arguments if needed
		// Args: graphql.FieldConfigArgument{
		// 	"title": &graphql.ArgumentConfig{
		// 		Type: graphql.String,
		// 	},
		// },
		Resolve: q.GetAllTodoResolver,
	}
}
