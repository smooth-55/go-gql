package gql

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRootQuery),
	fx.Provide(NewRootMutation),
)
