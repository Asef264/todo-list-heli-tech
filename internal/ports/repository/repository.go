package ports

import (
	"context"
	"todo-list/internal/domain"
)

type TodoItem interface {
	CreateTodoItem(ctx context.Context, entity domain.TodoItem) (*domain.TodoItem, error)
}
