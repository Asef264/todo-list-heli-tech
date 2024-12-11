package ports

import (
	"context"
	"database/sql"
	"todo-list/internal/domain"
)

type TodoItem interface {
	CreateTodoItem(ctx context.Context, entity domain.TodoList) (*domain.TodoList, error)
}

type todoItem struct {
	db *sql.DB
}

func NewTodoItem(db *sql.DB) TodoItem {
	return &todoItem{db: db}
}

func (ti todoItem) CreateTodoItem(ctx context.Context, entity domain.TodoList) (*domain.TodoList, error) {
	todoItem := domain.TodoList{}
	if err := ti.db.QueryRowContext(ctx, CreateTodoItemQuery, entity.ID, entity.Description, entity.DueDate, entity.FileID).Scan(&todoItem); err != nil {
		return nil, err
	}
	return &todoItem, nil
}
