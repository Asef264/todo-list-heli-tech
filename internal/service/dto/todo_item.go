package dto

import (
	"context"
	"time"
	"todo-list/pkg/validator"
)

type CreateTodoItemRequest struct {
	Description string    `json:"description" validate:"required,lte=300"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	FileID      string    `json:"file_id" validate:"omitempty,uuid"`
}

func (ctlr CreateTodoItemRequest) Validate(ctx context.Context) error {
	return validator.Validate(ctx, ctlr)
}

type TodoItem struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	FileID      string    `json:"file_id"`
	CreatedAt   time.Time `json:"created_at"`
}
