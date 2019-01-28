package resolver

import (
	"context"
	"testing"

	models "github.com/kazu69/todos/app/models"
)

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
