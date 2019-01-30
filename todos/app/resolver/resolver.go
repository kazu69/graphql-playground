package resolver

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"unsafe"

	graph "github.com/kazu69/todos/app/graph"
	models "github.com/kazu69/todos/app/models"
)

type Resolver struct {
	todos []models.Todo
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

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*bool, error) {
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
		return &deleted, nil
	}

	return &deleted, errors.New("Todo is not found")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todo(ctx context.Context, id *string) (*models.Todo, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id *string) (*models.User, error) {
	panic("not implemented")
}

func (r *queryResolver) Todos(ctx context.Context, userFilter *models.UserFilter, doneFilter *models.DoneFilter) ([]models.Todo, error) {
	filterdTodo := r.todos

	if userFilter != nil {
		userId := *(*string)(unsafe.Pointer(userFilter))
		tmp := []models.Todo{}
		for _, todo := range filterdTodo {
			if todo.User.ID == userId {
				tmp = append(tmp, todo)
			}
		}
		filterdTodo = tmp
	}

	if doneFilter != nil {
		state := *(*bool)(unsafe.Pointer(doneFilter))
		tmp := []models.Todo{}
		for _, todo := range filterdTodo {
			if todo.Done == state {
				tmp = append(tmp, todo)
			}
		}
		filterdTodo = tmp
	}

	return filterdTodo, nil
}

func (r *queryResolver) Users(ctx context.Context, userNameFilter *models.UserNameFilter) ([]models.User, error) {
	panic("not implemented")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (models.User, error) {
	panic("not implemented")
}
