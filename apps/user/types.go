package user

import "github.com/graphql-go/graphql"

type Types struct{}

func NewTypes() Types { return Types{} }

func (u Types) GetAllUserType(typeName string) *graphql.Object {

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

func (u Types) UserResponseType() *graphql.Object {

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

func (u Types) CreateUserInputType() *graphql.InputObject {
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
