package handler

import (
	"net/http"
	"time"
	"web_doscom/internal/database/model"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

// handler for uploading image or video
type GalleryHandler struct {
	GalleryService *service.GalleryService
}

func NewUploadHandler(m *service.GalleryService) *GalleryHandler {
	return &GalleryHandler{GalleryService: m}
}

// InsertGallery godoc
// @Summary Upload gallery (image only or mixed)
// @Description Upload multiple images/videos + metadata (gallery_type, description, event_date)
// @Accept multipart/form-data
// @Produce json
// @Param gallery_type formData string true "Type of gallery (fun, proker, achievment, work, activity, blog, pengurus, etc)"
// @Param description formData string true "Description of gallery"
// @Param event_date formData string true "Event date (YYYY-MM-DD)"
// @Param files formData []file true "Upload multiple files"
// @Success 200 {object} model.GalleryResponse "Successfully insert data"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Server error"
// @Security ApiKeyAuth
// @Tags Gallery
// @Router /api/v1/gallery/ [post]
func (m *GalleryHandler) InsertGallery(c *gin.Context) {

	var input model.CreateGallery
	if c.ShouldBind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fieldss: ",
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
			"error": "No files uploaded, ngantuk ta, po piye mas?",
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

	// insert to database woooooooooooooooooo
	result := []*model.Gallery{}
	for _, files := range uploadedFile {
		file_upload := &model.Gallery{
			GalleryName: files.GalleryName,
			GalleryType: input.GalleryType,
			Description: input.Description,
			EventDate:   input.EventDate,
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
		result = append(result, fileUpload)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully insert data",
		"data":    result,
	})

}

// insert photo profile
func (m *GalleryHandler) InsertProfilePic(c *gin.Context) {

	files, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	uploadedFile, err := m.GalleryService.UploadSingleImage(files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fileUpload := &model.Gallery{
		GalleryName: uploadedFile.GalleryName,
		GalleryType: "pengurus",
		Description: "foto Profile",
		EventDate:   time.Now().Format("2006-01-02"),
		FileSize:    uploadedFile.FileSize,
		MimeType:    uploadedFile.MimeType,
		AssetUrl:    uploadedFile.AssetUrl,
		Kategori:    uploadedFile.Kategori,
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
