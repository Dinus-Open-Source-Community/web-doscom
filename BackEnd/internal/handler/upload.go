package handler

import (
	"net/http"
	"strings"

	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	storageService *service.StorageService
}

func NewStorageHandler(storageService *service.StorageService) *StorageHandler {
	return &StorageHandler{
		storageService: storageService,
	}
}

// UploadImageRequest represents the upload request
type UploadImageRequest struct {
	Category string `form:"category" binding:"required,oneof=gallery blog work pengurus"`
}

// UploadImageResponse represents the upload response
type UploadImageResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	FileURL  string `json:"file_url,omitempty"`
	FileName string `json:"file_name,omitempty"`
}

// UploadImage handles image upload
// @Summary Upload an image
// @Description Upload an image file to MinIO storage
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file"
// @Param category formData string true "Category (gallery, blog, work, pengurus)"
// @Success 200 {object} UploadImageResponse
// @Failure 400 {object} UploadImageResponse
// @Failure 500 {object} UploadImageResponse
// @Security ApiKeyAuth
// @Router /api/v1/upload/image [post]
// UploadImage godoc
// @Summary      Upload image to MinIO
// @Description  Upload an image file to MinIO object storage with category
// @Tags         Upload
// @Accept       multipart/form-data
// @Produce      json
// @Param        file      formData  file    true  "Image file to upload"
// @Param        category  formData  string  true  "Category (gallery, blog, work, pengurus)"
// @Success      200  {object}  map[string]interface{}  "File uploaded successfully"
// @Failure      400  {object}  map[string]string       "Invalid request"
// @Failure      401  {object}  map[string]string       "Unauthorized"
// @Failure      500  {object}  map[string]string       "Upload failed"
// @Security     BearerAuth
// @Router       /api/v1/upload/image [post]
// func (h *StorageHandler) UploadImage(c *gin.Context) {
// 	var req UploadImageRequest
// 	ctx := c.Request.Context()
//
// 	// Bind form data
// 	if err := c.ShouldBind(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, UploadImageResponse{
// 			Success: false,
// 			Message: "Invalid request: " + err.Error(),
// 		})
// 		return
// 	}
//
// 	// Get user_id from JWT context
// 	userIDInterface, exists := c.Get("user_id")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, UploadImageResponse{
// 			Success: false,
// 			Message: "User not authenticated",
// 		})
// 		return
// 	}
//
// 	userID, ok := userIDInterface.(int)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, UploadImageResponse{
// 			Success: false,
// 			Message: "Invalid user ID type context",
// 		})
// 		return
// 	}
//
// 	// Get file from form
// 	fileHeader, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, UploadImageResponse{
// 			Success: false,
// 			Message: "File is required",
// 		})
// 		return
// 	}
//
// 	// Open file
// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, UploadImageResponse{
// 			Success: false,
// 			Message: "Failed to open file",
// 		})
// 		return
// 	}
// 	defer file.Close()
//
// 	// Upload file to MinIO with user ID
// 	fileURL, err := h.storageService.UploadFile(
// 		ctx,
// 		file,
// 		fileHeader,
// 		"gallery",
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, UploadImageResponse{
// 			Success: false,
// 			Message: "Failed to upload file: " + err.Error(),
// 		})
// 		return
// 	}
//
// 	// upload file to minio and database
// 	file
//
// 	c.JSON(http.StatusOK, UploadImageResponse{
// 		Success:  true,
// 		Message:  "File uploaded successfully",
// 		FileURL:  fileURL,
// 		FileName: filepath.Base(fileURL),
// 	})
// }

// DeleteFileRequest represents the delete request
type DeleteFileRequest struct {
	FileName string `json:"file_name" binding:"required"`
}

// DeleteFileResponse represents the delete response
type DeleteFileResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// DeleteFile handles file deletion
// @Summary Delete a file
// @Description Delete a file from MinIO storage
// @Tags upload
// @Accept json
// @Produce json
// @Param request body DeleteFileRequest true "Delete request"
// @Success 200 {object} DeleteFileResponse
// @Failure 400 {object} DeleteFileResponse
// @Failure 500 {object} DeleteFileResponse
// @Security ApiKeyAuth
// @Router /api/v1/upload/file [delete]
// DeleteFile godoc
// @Summary      Delete file from MinIO
// @Description  Delete an uploaded file from MinIO storage
// @Tags         Upload
// @Accept       json
// @Produce      json
// @Param        request  body      map[string]string  true  "File name to delete"
// @Success      200  {object}  map[string]string  "File deleted successfully"
// @Failure      400  {object}  map[string]string  "Invalid request"
// @Failure      401  {object}  map[string]string  "Unauthorized"
// @Failure      500  {object}  map[string]string  "Delete failed"
// @Security     BearerAuth
// @Router       /api/v1/upload/file [delete]
func (h *StorageHandler) DeleteFile(c *gin.Context) {
	var req DeleteFileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, DeleteFileResponse{
			Success: false,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	// Delete file from MinIO
	err := h.storageService.DeleteFile(c.Request.Context(), req.FileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DeleteFileResponse{
			Success: false,
			Message: "Failed to delete file: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, DeleteFileResponse{
		Success: true,
		Message: "File deleted successfully",
	})
}

// ListFilesRequest represents the list files request
type ListFilesRequest struct {
	Category string `form:"category" binding:"required,oneof=gallery blog work pengurus"`
}

// ListFilesResponse represents the list files response
type ListFilesResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Files   []string `json:"files,omitempty"`
	Count   int      `json:"count"`
}

// ListFiles handles listing files in a category
// @Summary List files in a category
// @Description List all files in a specific category folder
// @Tags upload
// @Accept json
// @Produce json
// @Param category query string true "Category (gallery, blog, work, pengurus)"
// @Success 200 {object} ListFilesResponse
// @Failure 400 {object} ListFilesResponse
// @Failure 500 {object} ListFilesResponse
// @Security ApiKeyAuth
// @Router /api/v1/upload/files [get]
// ListFiles godoc
// @Summary      List files from MinIO
// @Description  List all files in a specific category from MinIO
// @Tags         Upload
// @Accept       json
// @Produce      json
// @Param        category  query     string  false  "Category to filter (gallery, blog, work, pengurus)"
// @Success      200  {object}  map[string]interface{}  "Files retrieved successfully"
// @Failure      401  {object}  map[string]string       "Unauthorized"
// @Failure      500  {object}  map[string]string       "Failed to list files"
// @Security     BearerAuth
// @Router       /api/v1/upload/files [get]
func (h *StorageHandler) ListFiles(c *gin.Context) {
	var req ListFilesRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, ListFilesResponse{
			Success: false,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	// List files from MinIO
	files, err := h.storageService.ListFiles(c.Request.Context(), req.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ListFilesResponse{
			Success: false,
			Message: "Failed to list files: " + err.Error(),
		})
		return
	}

	// Convert full paths to URLs
	var fileURLs []string
	for _, file := range files {
		fileURL := strings.Replace(file, req.Category+"/", "", 1)
		fileURLs = append(fileURLs, fileURL)
	}

	c.JSON(http.StatusOK, ListFilesResponse{
		Success: true,
		Message: "Files retrieved successfully",
		Files:   fileURLs,
		Count:   len(fileURLs),
	})
}
