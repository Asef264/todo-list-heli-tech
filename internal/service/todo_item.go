package service

import (
	"context"

	ports "todo-list/internal/ports/repository"
	"todo-list/internal/service/cast"
	"todo-list/internal/service/dto"
	"todo-list/pkg/validator"

	"github.com/google/uuid"
)

type TodoItemService interface {
	Create(ctx context.Context, in dto.CreateTodoItemRequest) (dto.TodoItem, error)
}

type todoItem struct {
	todoItemRepository ports.TodoItem
}

func NewTodoItemService(dir ports.TodoItem) TodoItemService {
	return &todoItem{
		todoItemRepository: dir,
	}
}

func (ti todoItem) Create(ctx context.Context, in dto.CreateTodoItemRequest) (dto.TodoItem, error) {
	if err := validator.Validate(ctx, in); err != nil {
		return dto.TodoItem{}, ErrValidation
	}
	id := uuid.New().String()
	res, err := ti.todoItemRepository.CreateTodoItem(ctx, cast.ToTodoItemModel(in, id))
	if err != nil || res.ID == "" {
		return dto.TodoItem{}, ErrCreation
	}
	return cast.ToTodoItemResponse(*res), nil
}
