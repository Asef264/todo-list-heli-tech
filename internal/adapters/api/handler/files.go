package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"
	ports "todo-list/internal/ports/api"
	storage_service "todo-list/internal/service/storage"

	"github.com/gin-gonic/gin"
)

type storageController struct {
	storageService storage_service.StorageService
}

func NewStorageController(storageService storage_service.StorageService) ports.StorageController {
	return &storageController{
		storageService: storageService,
	}
}

func (sc storageController) Upload(c *gin.Context) {
	flag := c.Query("is_mock")
	isMock, err := strconv.ParseBool(flag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid boolean value for 'flag'"})
		return
	}

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
	err = sc.storageService.Upload(c.Request.Context(), data, file.Filename, isMock)
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
	flag := c.Query("is_mock")
	isMock, err := strconv.ParseBool(flag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid boolean value for 'flag'"})
		return
	}
	fileName := c.Param("file_name")
	object, err := sc.storageService.Download(c.Request.Context(), fileName, isMock)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", object)
}
