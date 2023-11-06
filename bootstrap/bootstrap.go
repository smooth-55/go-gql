package bootstrap

import (
	"context"
	"fmt"

	"github.com/smooth-55/graphql-go/infrastructure"
	"github.com/smooth-55/graphql-go/routes"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(bootstrap),
	infrastructure.Module,
	routes.Module,
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Printf("Starting Application")
			fmt.Printf("------------------------")
			fmt.Printf("------ Boilerplate 📺 ------")
			fmt.Printf("------------------------")

			go func() {
				fmt.Printf("🛠️ setting up routes")
				routes.Setup()
				// handler.Gin.SetMode(gin.ReleaseMode)
				handler.Gin.LoadHTMLGlob("templates/*")
				_ = handler.Gin.Run(":8000")
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("App stopped")
			return nil
		},
	})
}
