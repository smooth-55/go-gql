package user

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewQuery),
	fx.Provide(NewTypes),
	fx.Provide(NewMutation),
)
