package router

import (
	"database/sql"
	"log"
	"os"

	"todo-list/config"
	"todo-list/internal/adapters/api/handler"
	"todo-list/internal/adapters/repository"
	sqsAdapters "todo-list/internal/adapters/sqs"
	adapters "todo-list/internal/adapters/storage"
	ports "todo-list/internal/ports/storage"
	"todo-list/internal/service"
	storage_service "todo-list/internal/service/storage"
	dbPkg "todo-list/pkg/db"
	sqsPkg "todo-list/pkg/sqs"
	storagePkg "todo-list/pkg/storage"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB, cfg *config.Config) {

	dbPkg.MigrateUp(
		db,
		"./migrations/",
		cfg.DB.DBName,
	)
	storageS3Client := storagePkg.CreateAWSS3Client(cfg.S3Config.Endpoint, cfg.S3Config.AccessKey, cfg.S3Config.SecretKey, cfg.S3Config.Bucket)
	storageMinioClient, err := storagePkg.CreateMinioClient(cfg.MinioConfig.Host, cfg.MinioConfig.AccessKey, cfg.MinioConfig.SecretKey, cfg.MinioConfig.TLS)
	sqsClient := sqsPkg.CreateNewSQSClient(cfg.SQSConfig.AccessKey, cfg.SQSConfig.SecretKey, cfg.SQSConfig.QueueURL, cfg.SQSConfig.Region)

	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}

	RegisterStorageRepository(storageS3Client, storageMinioClient)

	// repository
	todoItemRepository := repository.NewTodoItem(db)
	storageRepository := RegisterStorageRepository(storageS3Client, storageMinioClient)
	sqsRepository := sqsAdapters.NewSQSAdapter(sqsClient)

	//service
	todoItemService := service.NewTodoItemService(todoItemRepository, sqsRepository)
	storageService := storage_service.NewStorageService(storageRepository)

	//adapter
	todoItemController := handler.NewTodoItemHandler(todoItemService)
	storageController := handler.NewStorageController(storageService)
	//todo_item
	router.POST("/todo_items", todoItemController.Create)

	//upload/download
	router.POST("/files", storageController.Upload)
	router.GET("/files/:file_name", storageController.Download)
}

func RegisterStorageRepository(s3Client *s3.S3, minioClient *minio.Client) ports.Storage {
	mockClient := make(map[string][]byte)
	storageType := os.Getenv("STORAGE_TYPE")
	switch storageType {
	case "minio":
		return adapters.NewMinioStorage(minioClient, mockClient)
	case "s3":
		return adapters.NewS3Storage(s3Client, mockClient)
	default:
		return adapters.NewS3Storage(s3Client, mockClient)

	}
}
