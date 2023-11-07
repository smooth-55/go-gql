package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/smooth-55/graphql-go/gql"
	"github.com/smooth-55/graphql-go/infrastructure"
)

type Handler struct {
	query    gql.RootQuery
	mutation gql.RootMutation
}

func NewHandler(
	query gql.RootQuery,
	mutation gql.RootMutation,
) Handler {
	return Handler{
		query:    query,
		mutation: mutation,
	}
}

func (h Handler) GetSchema() *graphql.Schema {
	var Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    h.query.GetQuery(),
		Mutation: h.mutation.GetMutation(),
	})
	if err != nil {
		panic(err)
	}
	return &Schema
}

func (h Handler) RequestHandler(c *gin.Context) {
	// Parse the incoming JSON data
	fmt.Println("inside handlerssok")
	var requestData map[string]interface{}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}

	// Extract the GraphQL query from the request
	query, _ := requestData["query"].(string)

	result := graphql.Do(graphql.Params{
		Schema:        *h.GetSchema(),
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
	router  infrastructure.Router
	handler Handler
}

// NewUserRoutes creates new user controller
func NewHandlerRoutes(
	router infrastructure.Router,
	handler Handler,

) HandlerRoutes {
	return HandlerRoutes{
		router:  router,
		handler: handler,
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
		gql.POST("/gql", i.handler.RequestHandler)
	}
}
