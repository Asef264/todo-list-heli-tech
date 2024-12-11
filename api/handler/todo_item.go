package handler

import (
	"log"
	"net/http"
	"todo-list/internal/service"
	"todo-list/internal/service/dto"

	"github.com/gin-gonic/gin"
)

type TodoItemHandler interface {
	Create(c *gin.Context)
}

type todoItem struct {
	todoItemService service.TodoItemService
}

func NewTodoItemHandler(tis service.TodoItemService) TodoItemHandler {
	return &todoItem{
		todoItemService: tis,
	}
}

func (ti todoItem) Create(c *gin.Context) {
	var request dto.CreateTodoItemRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	res, err := ti.todoItemService.Create(c.Request.Context(), request)
	if err != nil {
		log.Printf("Error creating todo item: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create TodoItem"})
		return
	}

	// Return the created TodoItem as a response
	c.JSON(http.StatusCreated, res)
}
