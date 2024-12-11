package cast

import (
	"todo-list/internal/domain"
	"todo-list/internal/service/dto"
)

func ToTodoItemModel(in dto.CreateTodoItemRequest, id string) domain.TodoList {
	return domain.TodoList{
		ID:          id,
		Description: in.Description,
		DueDate:     in.DueDate,
		FileID:      in.FileID,
	}
}

func ToTodoItemResponse(in domain.TodoList) dto.TodoItem {
	return dto.TodoItem{
		ID:          in.ID,
		Description: in.Description,
		DueDate:     in.DueDate,
		FileID:      in.FileID,
		CreatedAt:   in.CreatedAt,
	}
}