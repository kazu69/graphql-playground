package resolver

import (
	"context"
	"testing"

	models "github.com/kazu69/todos/app/models"
)

func TestTodo(t *testing.T) {
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
		},
	}

	qr := queryResolver{&r}
	ctx := context.Background()

	todo, err := qr.Todo(ctx, "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if todo.ID != "1" {
		t.Fatalf("error: query not correct filtering")
	}
}

func TestUser(t *testing.T) {
	r := Resolver{
		users: []models.User{
			{
				ID:   "123",
				Name: "Bob",
				Todos: []models.Todo{
					{
						ID:   "1",
						Text: "test",
						Done: false,
						User: models.User{
							ID:   "123",
							Name: "Bob",
						},
					},
				},
			},
		},
	}

	qr := queryResolver{&r}
	ctx := context.Background()

	user, err := qr.User(ctx, "123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if user.ID != "123" {
		t.Fatalf("error: query not correct filtering")
	}
}

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

	_, err := qr.Todos(ctx)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUsers(t *testing.T) {
	r := Resolver{
		users: []models.User{
			{
				ID:   "123",
				Name: "Bob",
				Todos: []models.Todo{
					{
						ID:   "1",
						Text: "test",
						Done: false,
						User: models.User{
							ID:   "123",
							Name: "Bob",
						},
					},
				},
			},
		},
	}

	qr := queryResolver{&r}
	ctx := context.Background()

	_, err := qr.Users(ctx)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
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

	if result == false {
		t.Fatalf("error: todo not deleted")
	}
}

func TestCreateUser(t *testing.T) {
	t.Skip()
}

func TestUpdateUser(t *testing.T) {
	t.Skip()
}

func TestDeleteUser(t *testing.T) {
	t.Skip()
}
