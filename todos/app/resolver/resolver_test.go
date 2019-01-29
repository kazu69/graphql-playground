package resolver

import (
	"context"
	"testing"

	models "github.com/kazu69/todos/app/models"
)

func TestTodos(t *testing.T) {
	r := Resolver{
		todos: []models.Todo{
			{
				ID:   "1",
				Text: "test",
				Done: false,
				User: models.User{
					ID:   "abc",
					Name: "test_user",
				},
			},
			{
				ID:   "2",
				Text: "test",
				Done: true,
				User: models.User{
					ID:   "abc",
					Name: "test_user",
				},
			},
		},
	}

	qr := queryResolver{&r}
	ctx := context.Background()

	stateFilter := models.DoneFilter{Done: false}
	userFilter := models.UserFilter{ID: "abc"}
	filterdTodo, err := qr.Todos(ctx, &userFilter, &stateFilter)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(filterdTodo) > 1 || len(filterdTodo) == 0 {
		t.Fatalf("error: query todos not exist")
	}

	if filterdTodo[0].Done != false {
		t.Fatalf("error: query not correct filtering")
	}

	stateFilter = models.DoneFilter{Done: true}
	userFilter = models.UserFilter{ID: "abc"}
	filterdTodo, err = qr.Todos(ctx, &userFilter, &stateFilter)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(filterdTodo) > 1 || len(filterdTodo) == 0 {
		t.Fatalf("error: query todos not exist")
	}

	if filterdTodo[0].Done != true {
		t.Fatalf("error: query not correct filtering")
	}
}

func TestCreateTodo(t *testing.T) {
	r := Resolver{
		todos: []models.Todo{},
	}

	mr := mutationResolver{&r}

	input := models.NewTodo{
		Text:   "test",
		UserID: "123",
	}

	ctx := context.Background()
	_, err := mr.CreateTodo(ctx, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(mr.todos) <= 0 {
		t.Fatalf("error: todo not created")
	}
}

func TestUpdateTodo(t *testing.T) {
	r := Resolver{
		todos: []models.Todo{
			{
				ID:   "123",
				Text: "test",
				Done: false,
				User: models.User{
					ID:   "abc",
					Name: "test_user",
				},
			},
		},
	}

	mr := mutationResolver{&r}
	ctx := context.Background()

	input := models.UpdateTodo{
		ID:   "123",
		Done: true,
	}

	todo, err := mr.UpdateTodo(ctx, input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !todo.Done || mr.todos[0].Done != true {
		t.Fatalf("error: todo not updated")
	}
}

func TestDeleteTodo(t *testing.T) {
	r := Resolver{
		todos: []models.Todo{
			{
				ID:   "123",
				Text: "test",
				Done: false,
				User: models.User{
					ID:   "abc",
					Name: "test_user",
				},
			},
		},
	}

	mr := mutationResolver{&r}
	ctx := context.Background()

	id := "123"

	result, err := mr.DeleteTodo(ctx, id)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result == nil {
		t.Fatalf("error: todo not deleted")
	}
}
