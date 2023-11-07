package config

import (
	"github.com/smooth-55/graphql-go/apps/todo"
	"github.com/smooth-55/graphql-go/apps/user"
	"github.com/smooth-55/graphql-go/gql"
	"github.com/smooth-55/graphql-go/routes"
	"go.uber.org/fx"
)

// Module exports dependency
var InstalledModule = fx.Options(
	user.Module,
	todo.Module,
	gql.Module,
	routes.Module,
)
