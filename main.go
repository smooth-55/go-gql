package main

import (
	"github.com/smooth-55/graphql-go/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
