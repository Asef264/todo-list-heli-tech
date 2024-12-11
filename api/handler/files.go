package handler

import (
	"io/ioutil"
	"net/http"
	storage_service "todo-list/internal/service/storage"

	"github.com/gin-gonic/gin"
)

type StorageController interface {
	Upload(ctx *gin.Context)
	Download(ctx *gin.Context)
}
type storageController struct {
	storageService storage_service.StorageService
}

func NewStorageController(storageService storage_service.StorageService) StorageController {
	return &storageController{
		storageService: storageService,
	}
}

func (sc storageController) Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open file"})
		return
	}

	data, err := ioutil.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error on converting, system error"})
	}
	err = sc.storageService.Upload(c.Request.Context(), data, file.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to upload file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
		"file":    file.Filename,
	})
}

func (sc storageController) Download(c *gin.Context) {
	fileName := c.Param("file_name")
	object, err := sc.storageService.Download(c.Request.Context(), fileName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", object)
}
