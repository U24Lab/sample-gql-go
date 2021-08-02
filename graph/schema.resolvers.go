package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"

	"github.com/U24Lab/sample-gql-go/graph/generated"
	"github.com/U24Lab/sample-gql-go/graph/model"
	"github.com/U24Lab/sample-gql-go/mongoDBRepo"
)

var todoRepo mongoDBRepo.TodoRepo = mongoDBRepo.New()

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.NewTodo, error) {

	todo := &model.NewTodo{
		Text:   input.Text,
		UserID: strconv.Itoa(rand.Int()),
	}
	todoRepo.Save(todo)

	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.NewTodo, error) {

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

	return todoRepo.GetAll(), nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
