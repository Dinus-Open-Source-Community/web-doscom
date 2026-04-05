package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUploadRoutes(rg *gin.RouterGroup, storageHandler *handler.StorageHandler) {
	upload := rg.Group("/upload")
	upload.Use(auth.AuthMiddleware())
	{
		// upload.POST("/image", storageHandler.UploadImage)
		upload.DELETE("/file", storageHandler.DeleteFile)
		upload.GET("/files", storageHandler.ListFiles)
	}
}
