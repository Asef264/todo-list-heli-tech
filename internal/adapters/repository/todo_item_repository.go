package repository

import (
	"context"
	"database/sql"
	"todo-list/internal/domain"
	ports "todo-list/internal/ports/repository"
)

type todoItem struct {
	db *sql.DB
}

func NewTodoItem(db *sql.DB) ports.TodoItem {
	return &todoItem{db: db}
}

func (ti todoItem) CreateTodoItem(ctx context.Context, entity domain.TodoItem) (*domain.TodoItem, error) {
	todoItem := domain.TodoItem{}
	if err := ti.db.QueryRow(CreateTodoItemQuery, entity.ID, entity.Description, entity.DueDate, entity.FileID).Scan(
		&todoItem.ID, &todoItem.Description, &todoItem.DueDate, &todoItem.FileID, &todoItem.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &todoItem, nil
}
