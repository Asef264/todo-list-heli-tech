package service

import (
	"context"
	"fmt"

	portsRepo "todo-list/internal/ports/repository"
	ports "todo-list/internal/ports/sqs"
	"todo-list/internal/service/cast"
	"todo-list/internal/service/dto"
	"todo-list/pkg/validator"

	"github.com/google/uuid"
)

type TodoItemService interface {
	Create(ctx context.Context, in dto.CreateTodoItemRequest) (dto.TodoItem, error)
}

type todoItem struct {
	todoItemRepository portsRepo.TodoItem
	sqsQueue           ports.SQS
}

func NewTodoItemService(dir portsRepo.TodoItem, sqsQueue ports.SQS) TodoItemService {
	return &todoItem{
		todoItemRepository: dir,
		sqsQueue:           sqsQueue,
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
	go func() {
		ti.sqsQueue.SendMessage(ctx, fmt.Sprint("Todo item created by ID:", id))
	}()
	return cast.ToTodoItemResponse(*res), nil
}
