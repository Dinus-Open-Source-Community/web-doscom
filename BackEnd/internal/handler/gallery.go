package handler

import (
	"math"
	"net/http"
	"strconv"
	"time"
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

	// Ambil form value satu per satu
	galleryName := c.PostForm("gallery_name")
	galleryType := c.PostForm("gallery_type")
	description := c.PostForm("description")
	eventDate := c.PostForm("event_date")
	kategori := c.PostForm("kategori")

	missingFields := []string{}
	if galleryName == "" {
		missingFields = append(missingFields, "gallery_name")
	}
	if galleryType == "" {
		missingFields = append(missingFields, "gallery_type")
	}
	if description == "" {
		missingFields = append(missingFields, "description")
	}
	if eventDate == "" {
		missingFields = append(missingFields, "event_date")
	}
	if kategori == "" {
		missingFields = append(missingFields, "kategori")
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
		missingFields = append(missingFields, "files (at least 1 file)")
	}

	if len(missingFields) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing required fields",
			"details": missingFields,
		})
		return
	}

	// upload file
	uploadedFile, err := m.GalleryService.UploadImage(files)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// insert to database
	result := []*model.Gallery{}
	for _, files := range uploadedFile {
		file_upload := &model.Gallery{
			GalleryName: files.GalleryName,
			GalleryType: galleryType,
			Description: description,
			EventDate:   eventDate,
			FileSize:    files.FileSize,
			MimeType:    files.MimeType,
			AssetUrl:    files.AssetUrl,
			Kategori:    files.Kategori,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		fileUpload, err := m.GalleryService.InsertGallery(file_upload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to insert data",
			})
			return
		}
		result = append(result, uploadedGallery)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully insert data",
		"data":    result,
	})
}

// get gallery by type

// GetGalleryByType godoc
// @Summary Get Gallery By Type
// @Description Mengambil data gallery berdasarkan tipe tertentu
// @Tags Gallery
// @Accept json
// @Produce json
// @Param type query string true "Gallery type (fun, proker, achievment, work, activity, blog, pengurus, etc)"
// @Param page query int false "Page number"
// @Param limit query int false "Page limit"
// @Success 200 {object} map[string]interface{} "Successfully fetch gallery data"
// @Failure 401 {object} map[string]string "Unauthorized, role tidak valid"
// @Failure 500 {object} map[string]string "Failed to fetch gallery data"
// @Security ApiKeyAuth
// @Router /api/v1/gallery/ [get]
func (m *GalleryHandler) GetGalleryByType(c *gin.Context) {

	// filtered by type
	tipe := c.Query("type")

	// pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	// get role
	// role := c.MustGet("role").(string)

	// if role != "Kor_Medcrev" {
	// c.JSON(http.StatusUnauthorized, gin.H{
	// "error": "Unauthorized, koe sopo wok",
	// })
	// return
	// }

	// get all gallery by type
	GalleryList, count, err := m.GalleryService.GetAllGalleryByType(tipe, page, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch gallery data",
		})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetch gallery data",
		"data":    GalleryList,
		"metadata": gin.H{
			"page":        page,
			"limit":       limit,
			"total_items": count,
			"total_pages": totalPages,
		},
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if err := m.GalleryService.DeleteGallery(id); err != nil {
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
func (m *GalleryHandler) InsertProfilePic(c *gin.Context) {
	// Get user ID from context
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not authenticated",
		})
		return
	}

	userID, ok := userIDVal.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID type",
		})
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer file.Close()

	// Upload to MinIO with pengurus category
	fileURL, err := m.StorageService.UploadFile(
		c.Request.Context(),
		file,
		fileHeader,
		"pengurus",
		uint(userID),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Insert to database
	fileUpload := &model.Gallery{
		GalleryName: fileHeader.Filename,
		GalleryType: "pengurus",
		Description: "foto Profile",
		EventDate:   time.Now().Format("2006-01-02"),
		FileSize:    fileHeader.Size,
		MimeType:    fileHeader.Header.Get("Content-Type"),
		AssetUrl:    fileURL,
		Kategori:    "image",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	upload, err := m.GalleryService.InsertGallery(fileUpload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully insert data",
		"data":    upload,
	})
}

// get gallery by id

func (m *GalleryHandler) GetGalleryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid gallery ID",
			"details": err.Error(),
		})
		return
	}

	gallery, err := m.GalleryService.GetGalleryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch gallery data",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetch gallery data",
		"data":    gallery,
	})
}

// get all gallery

func (m *GalleryHandler) GetAllGallery(c *gin.Context) {

	// pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	galleryList, count, err := m.GalleryService.GetAllGallery(page, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch gallery data",
			"details": err.Error(),
		})
		return
	}

	totalPages := 1
	if limit > 0 {
		totalPages = int(math.Ceil(float64(count) / float64(limit)))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully fetch gallery data",
		"data":    galleryList,
		"metadata": gin.H{
			"page":        page,
			"limit":       limit,
			"total_items": count,
			"total_pages": totalPages,
		},
	})
}

// update gallery by id
func (m *GalleryHandler) UpdateGallery(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid gallery ID",
			"details": err.Error(),
		})
		return
	}

	var input model.GalleryUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	updateGallery, err := m.GalleryService.UpdateGallery(id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update gallery data",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated gallery data",
		"data":    updateGallery,
	})
}
