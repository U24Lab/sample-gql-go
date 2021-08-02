package graph

import "github.com/U24Lab/sample-gql-go/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	todo []*model.Todo
}
