package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Router Gin Router
type Router struct {
	Gin *gin.Engine
}

// NewRouter : all the routes are defined here
func NewRouter() Router {

	httpRouter := gin.Default()

	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "graphql 📺 API Up and Running"})
	})

	return Router{
		Gin: httpRouter,
	}
}

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRouter),
)
