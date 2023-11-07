package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/smooth-55/graphql-go/apps/todo"
	"github.com/smooth-55/graphql-go/apps/user"
)

type RootQuery struct {
	userQuery    user.Query
	todoQuery    todo.Query
	userMutation user.Mutation
}

func NewRootQuery(
	userQuery user.Query,
	userMutation user.Mutation,

	todoQuery todo.Query,

) RootQuery {
	return RootQuery{
		userQuery:    userQuery,
		userMutation: userMutation,
		todoQuery:    todoQuery,
	}
}

func (r RootQuery) GetQuery() *graphql.Object {
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": r.userQuery.GetAllUserField(),
			"user":  r.userQuery.GetOneUserField(),
			"todos": r.todoQuery.GetAllTodoField(),
			// add other queries
		},
	})
	return rootQuery

}
