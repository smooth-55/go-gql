package todo

import "github.com/graphql-go/graphql"

type Types struct{}

func NewTypes() Types { return Types{} }

func (u Types) TodoType(typeName string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: typeName,
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"is_completed": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	})

}
