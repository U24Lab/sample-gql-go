package main

import (
	"os"

	"github.com/U24Lab/sample-gql-go/httpCustom"
	"github.com/U24Lab/sample-gql-go/middleware"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	server := gin.Default()
	server.Use(middleware.BasicAuth())
	server.POST("/query", httpCustom.GraphQLHandler())
	server.GET("/", httpCustom.PlaygroundHandler())
	server.Run(port)
}
