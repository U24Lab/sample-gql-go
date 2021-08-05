package httpCustom

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/U24Lab/sample-gql-go/graph"
	"github.com/U24Lab/sample-gql-go/graph/generated"
	"github.com/gin-gonic/gin"
)

func GraphQLHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	fmt.Println("Calling")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

func PlaygroundHandler() gin.HandlerFunc {

	h := playground.Handler("GraphQL playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
