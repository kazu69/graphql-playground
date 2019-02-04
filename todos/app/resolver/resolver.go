package resolver

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	graph "github.com/kazu69/todos/app/graph"
	models "github.com/kazu69/todos/app/models"
)

type Resolver struct {
	todo  models.Todo
	user  models.User
	todos []models.Todo
	users []models.User
}

func (r *Resolver) Mutation() graph.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() graph.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (models.Todo, error) {
	todo := models.Todo{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Text: input.Text,
		Done: false,
	}

	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input models.UpdateTodo) (models.Todo, error) {
	affected := models.Todo{}
	for i, todo := range r.todos {
		if todo.ID == input.ID {
			r.todos[i].Done = input.Done
			affected = r.todos[i]
			break
		}
	}

	if affected.ID == "" {
		affected.ID = "-1"
		return affected, errors.New("Todo is not found")
	}

	return affected, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (bool, error) {
	newTodos := []models.Todo{}
	var deleted bool
	for _, todo := range r.todos {
		if todo.ID != id {
			newTodos = append(newTodos, todo)
		} else {
			deleted = true
		}
	}

	if deleted {
		r.todos = newTodos
		return deleted, nil
	}

	return deleted, errors.New("Todo is not found")
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (models.User, error) {
	panic("Not implement")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UpdateUser) (models.User, error) {
	panic("Not implement")
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic("Not implement")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, id string) (*models.Todo, error) {
	var response models.Todo
	for _, todo := range r.todos {
		if todo.ID == id {
			response = todo
		}
	}

	return &response, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	var response models.User
	for _, user := range r.users {
		if user.ID == id {
			response = user
		}
	}
	return &response, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]models.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	return r.users, nil
}
