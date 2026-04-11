package handler

import (
	"log"
	"net/http"
	"strconv"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

// handler for uploading image or video
type GalleryHandler struct {
	GalleryService *service.GalleryService
	StorageService *service.StorageService
}

func NewUploadHandler(galleryService *service.GalleryService, storageService *service.StorageService) *GalleryHandler {
	return &GalleryHandler{
		GalleryService: galleryService,
		StorageService: storageService,
	}
}

// InsertGallery godoc
// @Summary Upload gallery images/videos
// @Description Upload multiple images or videos to MinIO storage with metadata. All files stored in MinIO bucket 'doscom-uploads/gallery'.
// @Accept multipart/form-data
// @Produce json
// @Param gallery_type formData string true "Type of gallery (fun, proker, achievment, work, activity, blog, pengurus, etc)"
// @Param description formData string true "Description of gallery"
// @Param event_date formData string true "Event date (YYYY-MM-DD)"
// @Param files formData []file true "Upload multiple image/video files"
// @Success 200 {object} model.GalleryResponse "Successfully uploaded to MinIO"
// @Failure 400 {object} map[string]string "Bad request - missing fields or invalid files"
// @Failure 401 {object} map[string]string "Unauthorized - authentication required"
// @Failure 500 {object} map[string]string "Server error - upload failed"
// @Security BearerAuth
// @Tags Gallery
// @Router /api/v1/gallery/ [post]
func (m *GalleryHandler) InsertGallery(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("user_id").(int)
	userRole := c.MustGet("role").(string)

	var input model.CreateGallery
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed: " + err.Error(),
		})
		return
	}

	allowedRole := map[string]bool{
		constants.RoleKeyKoorMedcrev: true,
		constants.RoleKeySuperAdmin:  true,
	}

	if !allowedRole[userRole] {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "You're not allowed broo",
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No files uploaded",
		})
		return
	}
	if len(files) > 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Max upload 5 file",
		})
		return
	}

	galleryData := &model.GalleryInsert{
		IDUsers:     userID,
		GalleryName: input.GalleryName,
		GalleryType: input.GalleryType,
		Description: input.Description,
		EventDate:   input.EventDate,
	}
	fileUploadData := make([]*model.UploadFileRequest, len(files))
	for i, fileHeader := range files {
		fileContent, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed while opening file upload",
			})
			return
		}
		defer fileContent.Close()

		fileUploadData[i] = &model.UploadFileRequest{
			FileHeader: fileHeader,
			File:       fileContent,
			Folder:     "gallery",
			UserID:     uint(userID),
		}
	}

	galleryResponse, err := m.GalleryService.UploadAndInsertGalleryMultiple(
		ctx,
		galleryData,
		fileUploadData,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to insert and upload gallery",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully insert data",
		"data":    galleryResponse,
	})
}

func (m *GalleryHandler) GetAllGalleryAndByYear(c *gin.Context) {
	ctx := c.Request.Context()

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		log.Println("invalid page param: ", pageStr)
		page = 1
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		log.Println("invalid limit param: ", limitStr)
		limit = 10
	}
	offset := (page - 1) * limit

	// call service
	startYear := c.Query("start_year")
	endYear := c.Query("end_year")

	galleryList, totalPages, currentPage, err := m.GalleryService.GetAllGalleryAndByDate(
		ctx,
		startYear, endYear,
		limit, offset,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to get data gallery, some issue at backend :))",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Successfully get data",
		"totalPages":  totalPages,
		"currentPage": currentPage,
		"gallery":     galleryList,
	})
}

// delete gallery by id
// DeleteGallery godoc
// @Summary Delete gallery by ID
// @Description Delete gallery data dan file nya berdasarkan ID
// @Tags Gallery
// @Accept json
// @Produce json
// @Param id path int true "Gallery ID"
// @Success 200 {object} map[string]interface{} "Successfully delete Gallery"
// @Failure 400 {object} map[string]interface{} "Invalid ID"
// @Failure 404 {object} map[string]interface{} "Gallery not found"
// @Security BearerAuth
// @Router /api/v1/gallery/{id} [delete]
func (m *GalleryHandler) DeleteGallery(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	_, err := constants.GetRoleInfo(userRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "you cannot access this resource",
		})
		return
	}

	allowedRole := map[string]bool{
		constants.RoleKeyKoorMedcrev: true,
		constants.RoleKeySuperAdmin:  true,
	}

	if !allowedRole[userRole] {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "request cannot be processed",
			"message": "role not valid or have no permission",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if err := m.GalleryService.DeleteGallery(ctx, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Gallery not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete Gallery, bang nasi padang satu bungkus bang",
	})
}

// insert photo profile
// func (m *GalleryHandler) InsertProfilePic(c *gin.Context) {
// 	// Get user ID from context
// 	userIDVal, exists := c.Get("user_id")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "User not authenticated",
// 		})
// 		return
// 	}
//
// 	userID, ok := userIDVal.(int)
// 	if !ok {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid user ID type",
// 		})
// 		return
// 	}
//
// 	fileHeader, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to read file",
// 		})
// 		return
// 	}
//
// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Failed to open file",
// 		})
// 		return
// 	}
// 	defer file.Close()
//
// 	// Upload to MinIO with pengurus category
// 	fileURL, err := m.StorageService.UploadFile(
// 		c.Request.Context(),
// 		file,
// 		fileHeader,
// 		"pengurus",
// 		uint(userID),
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
//
// 	// Insert to database
// 	fileUpload := &model.Gallery{
// 		GalleryName: fileHeader.Filename,
// 		GalleryType: "pengurus",
// 		Description: "foto Profile",
// 		EventDate:   time.Now().Format("2006-01-02"),
// 		FileSize:    fileHeader.Size,
// 		MimeType:    fileHeader.Header.Get("Content-Type"),
// 		AssetUrl:    fileURL,
// 		Kategori:    "image",
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}
//
// 	upload, err := m.GalleryService.InsertGallery(fileUpload)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to insert data",
// 		})
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Successfully insert data",
// 		"data":    upload,
// 	})
// }
