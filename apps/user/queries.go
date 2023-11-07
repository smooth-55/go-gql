package user

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Query struct {
	types Types
}

func NewQuery(
	types Types,

) Query {
	return Query{
		types: types,
	}
}

func (q Query) GetAllUsersResolver(p graphql.ResolveParams) (interface{}, error) {
	users := []User{
		{
			Id:       1,
			Username: "username",
			FullName: "full name",
			Phone:    "98xxxxxx",
		},
		{
			Id:       2,
			Username: "username2",
			FullName: "full name2",
			Phone:    "98xxxxxx",
		},
	}
	return users, nil
}

func (q Query) GetAllUserField() *graphql.Field {
	return &graphql.Field{
		Type:    graphql.NewList(q.types.GetAllUserType("GetAllUser")),
		Args:    nil,
		Resolve: q.GetAllUsersResolver,
	}

}

func (q Query) GetOneUserResolver(p graphql.ResolveParams) (interface{}, error) {
	Id := p.Args["id"].(string)

	users := []User{
		{
			Id:       1,
			Username: "username",
			FullName: "full name",
			Phone:    "98xxxxxx",
		},
		{
			Id:       2,
			Username: "username2",
			FullName: "full name2",
			Phone:    "98xxxxxx",
		},
	}
	for _, user := range users {
		if fmt.Sprintf("%v", user.Id) == Id {
			return user, nil
		}
	}
	return nil, errors.New("User doesn't exists")
}

func (q Query) GetOneUserField() *graphql.Field {
	return &graphql.Field{
		Type: q.types.GetAllUserType("GetOneUser"),
		Args: graphql.FieldConfigArgument{

			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: q.GetOneUserResolver,
	}

}
