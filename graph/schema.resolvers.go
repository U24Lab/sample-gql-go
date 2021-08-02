package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/U24Lab/sample-gql-go/graph/generated"
	"github.com/U24Lab/sample-gql-go/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {

	todo1 := &model.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID: input.UserID,
		},
	}
	r.todo = append(r.todo, todo1)

	return todo1, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {

	// r.todo = make([]*model.Todo, 0)
	// todo1 := &model.Todo{
	// 	ID:   "ABCO",
	// 	Text: "Demo",
	// 	Done: true,
	// 	User: &model.User{
	// 		ID:   "1212",
	// 		Name: "demoName",
	// 	},
	// }
	// r.todo = append(r.todo, todo1)

	return r.todo, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
