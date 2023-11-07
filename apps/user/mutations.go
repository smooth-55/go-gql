package user

import "github.com/graphql-go/graphql"

type Mutation struct {
	types Types
}

func NewMutation(
	types Types,
) Mutation {
	return Mutation{
		types: types,
	}
}

func (m Mutation) CreateUserResolver(p graphql.ResolveParams) (interface{}, error) {
	inputs := p.Args["input"].(map[string]interface{})
	user := User{
		Id:       1,
		Username: inputs["username"].(string),
		FullName: inputs["full_name"].(string),
		Phone:    inputs["phone"].(string),
	}
	return user, nil
}

func (m Mutation) CreateUserField() *graphql.Field {
	return &graphql.Field{
		Type:    m.types.UserResponseType(),
		Resolve: m.CreateUserResolver,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type:        m.types.CreateUserInputType(),
				Description: "this is des",
			},
		},
		Description: "This mutation is used to create a new user",
	}
}
