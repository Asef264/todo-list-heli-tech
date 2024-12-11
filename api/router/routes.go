package router

import (
	"database/sql"
	"todo-list/api/handler"
	"todo-list/config"
	"todo-list/internal/adapters/persistence"
	"todo-list/internal/ports"
	"todo-list/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB, cfg *config.Config) {

	persistence.MigrateUp(
		db,
		"./migrations/",
		cfg.DB.DBName,
	)

	// repository
	todoItemRepository := ports.NewTodoItem(db)

	//service
	todoItemService := service.NewTodoItemService(todoItemRepository)

	//handler
	todoItemController := handler.NewTodoItemHandler(todoItemService)

	//todo_item
	router.POST("/todo_items", todoItemController.Create)

	//upload/download
	router.POST("/files")
}
