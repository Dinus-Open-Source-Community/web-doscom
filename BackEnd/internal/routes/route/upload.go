package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUploadRoutes registers all upload-related routes
func RegisterUploadRoutes(rg *gin.RouterGroup, storageHandler *handler.StorageHandler) {
	upload := rg.Group("/upload")
	upload.Use(auth.AuthMiddleware()) // Protect all upload endpoints with JWT authentication
	{
		// Upload image endpoint
<<<<<<< HEAD
		// upload.POST("/image", storageHandler.UploadImage)
		// Delete file endpoint
		upload.DELETE("/file", storageHandler.DeleteFile)
=======
		upload.POST("/image", storageHandler.UploadImage)

		// Delete file endpoint
		upload.DELETE("/file", storageHandler.DeleteFile)

>>>>>>> master
		// List files endpoint
		upload.GET("/files", storageHandler.ListFiles)
	}
}
