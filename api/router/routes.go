package router

import (
	"database/sql"
	"log"
	"os"
	"todo-list/api/handler"
	"todo-list/config"
	"todo-list/internal/adapters/persistence"
	s3Adaptor "todo-list/internal/adapters/s3"
	"todo-list/internal/ports"
	"todo-list/internal/service"
	storage_service "todo-list/internal/service/storage"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB, cfg *config.Config) {

	persistence.MigrateUp(
		db,
		"./migrations/",
		cfg.DB.DBName,
	)
	storageS3Client := s3Adaptor.CreateAWSS3Client("http://localhost:9000", "minioadmin", "minioadmin")
	storageMinioClient, err := s3Adaptor.CreateMinioClient("localhost:9000", "minioadmin", "minioadmin", false)
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}

	RegisterStorageRepository(storageS3Client, storageMinioClient)

	// repository
	todoItemRepository := ports.NewTodoItem(db)
	storageRepository := RegisterStorageRepository(storageS3Client, storageMinioClient)

	//service
	todoItemService := service.NewTodoItemService(todoItemRepository)
	storageService := storage_service.NewStorageService(storageRepository)

	//handler
	todoItemController := handler.NewTodoItemHandler(todoItemService)
	storageController := handler.NewStorageController(storageService)
	//todo_item
	router.POST("/todo_items", todoItemController.Create)

	//upload/download
	router.POST("/files", storageController.Upload)
	router.GET("/files/:file_name", storageController.Download)
}

func RegisterStorageRepository(s3Client *s3.Client, minioClient *minio.Client) ports.Storage {
	storageType := os.Getenv("STORAGE_TYPE")
	switch storageType {
	case "minio":
		return ports.NewMinioStorage(minioClient)
	case "aws":
		return ports.NewS3Storage(s3Client)
	default:
		return ports.NewMinioStorage(minioClient)

	}
}
