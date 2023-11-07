package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/smooth-55/graphql-go/apps/user"
)

type RootGql struct {
	userQuery    user.Query
	userMutation user.Mutation
}

func NewRootGql(
	userQuery user.Query,
	userMutation user.Mutation,
) RootGql {
	return RootGql{
		userQuery:    userQuery,
		userMutation: userMutation,
	}
}

func (r RootGql) GetRootQuery() *graphql.Object {
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": r.userQuery.GetAllUserField(),
			"user":  r.userQuery.GetOneUserField(),
		},
	})
	return rootQuery

}

func (r RootGql) GetRootMutation() *graphql.Object {
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": r.userMutation.CreateUserField(),
		},
	})
	return rootQuery

}
