package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/smooth-55/graphql-go/apps/user"
)

type RootMutation struct {
	userMutation user.Mutation
}

func NewRootMutation(
	userMutation user.Mutation,

) RootMutation {
	return RootMutation{
		userMutation: userMutation,
	}
}

func (r RootMutation) GetMutation() *graphql.Object {
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": r.userMutation.CreateUserField(),
			// add other mutations
		},
	})
	return rootQuery

}
