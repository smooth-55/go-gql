package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/smooth-55/graphql-go/gql"
	"github.com/smooth-55/graphql-go/infrastructure"
)

func Handler(c *gin.Context) {
	// Parse the incoming JSON data
	var requestData map[string]interface{}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	// Extract the GraphQL query from the request
	query, _ := requestData["query"].(string)

	result := graphql.Do(graphql.Params{
		Schema:        gql.Schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		c.JSON(400, gin.H{"errors": result.Errors})
		return
	}

	// Return the JSON response
	c.JSON(200, gin.H{"data": result.Data})
}

// UserRoutes struct
type HandlerRoutes struct {
	router infrastructure.Router
}

// NewUserRoutes creates new user controller
func NewHandlerRoutes(
	router infrastructure.Router,

) HandlerRoutes {
	return HandlerRoutes{
		router: router,
	}
}

// Setup user routes
func (i HandlerRoutes) Setup() {
	gql := i.router.Gin.Group("")
	{
		gql.GET("/playground", func(c *gin.Context) {
			c.HTML(200, "playground.html", gin.H{
				"message": "Graphql playground",
			})
		})
		gql.POST("/gql", Handler)
	}
}
