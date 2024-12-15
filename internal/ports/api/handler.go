package ports

import "github.com/gin-gonic/gin"

type TodoItemHandler interface {
	Create(c *gin.Context)
}

type StorageController interface {
	Upload(ctx *gin.Context)
	Download(ctx *gin.Context)
}
