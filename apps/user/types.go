package user

import "github.com/graphql-go/graphql"

type UserTypes struct{}

func NewUserTypes() UserTypes { return UserTypes{} }

func (u UserTypes) GetAllUserType(typeName string) *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"full_name": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

}

func (u UserTypes) UserResponseType() *graphql.Object {

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "UserType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"full_name": &graphql.Field{
				Type: graphql.String,
			},
			"phone": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func (u UserTypes) CreateUserInputType() *graphql.InputObject {
	obj := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateTodoInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"username": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"full_name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"phone": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
	return obj
}
