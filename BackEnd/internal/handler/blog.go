package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	Service *service.BlogService
}

func NewBlogHandler(m *service.BlogService) *BlogHandler {
	return &BlogHandler{Service: m}
}

// CreateBlog godoc
// @Summary Create a new blog
// @Description Create blog dengan upload gambar baru dan/atau memilih gambar yang sudah ada
// @Tags Blog
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Judul Blog"
// @Param slug formData string true "Slug Blog"
// @Param content formData string true "Konten Blog"
// @Param kategori formData string true "Kategori Blog"
// @Param published_at formData string false "Tanggal publish (format RFC3339)"
// @Param is_published formData bool false "Status publish"
// @Param id_work formData int true "ID Work"
// @Param id_pengurus formData int true "ID Pengurus"
// @Success 200 {object} map[string]interface{} "Blog created successfully"
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/v1/blogs/ [post]
func (h *BlogHandler) CreateBlog(c *gin.Context) {
	user_role := c.MustGet("role").(string)

	existingID := c.PostFormArray("existingID_image")
	existingIDS := make([]int, 0, len(existingID))
	for _, v := range existingID {
		id, err := strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "harus integer",
			})
			return
		}

		existingIDS = append(existingIDS, id)
	}

	if user_role != "Kor_Medcrev" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized, koe rak dhion wok",
		})
		return
	}

	var input model.RegisterBlog
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// insert foto -> multiple file
	blog := &model.Blog{
		Title:       input.Title,
		Slug:        input.Slug,
		Content:     input.Content,
		Kategori:    input.Kategori,
		PublishedAt: input.PublishedAt,
		IsPublished: input.IsPublished,
		WorkID:      input.WorkID,
		PengurusID:  input.PengurusID,
	}

	if err := h.Service.BlogModel.InsertBlog(blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert data",
		})
		return
	}

	// insert gallery and blog_gallery
	blogGallery, err := h.Service.CreateBlogImage(blog.ID, existingIDS, files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to insert data",
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog input is valid",
		"data":    blogGallery,
	})
}

// Create Blog (validation only, no DB insert)
func (h *BlogHandler) Create(c *gin.Context) {
	var blog model.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Example: add custom validation if needed
	if blog.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	if blog.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content is required"})
		return
	}

	// If all validation passed
	c.JSON(http.StatusOK, gin.H{"message": "Blog input is valid", "blog": blog})
}

// List all Blogs
func (h *BlogHandler) List(c *gin.Context) {
	var blogs []model.Blog
	if err := h.Service.Order("created_at DESC").Find(&blogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// Get Blog by ID
func (h *BlogHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var blog model.Blog
	if err := h.DB.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// Update Blog
func (h *BlogHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input model.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var blog model.Blog
	if err := h.DB.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}
	blog.Title = input.Title
	blog.Slug = input.Slug
	blog.Content = input.Content
	blog.PublishedAt = input.PublishedAt
	blog.IsPublished = input.IsPublished
	blog.Kategori = input.Kategori
	blog.WorkID = input.WorkID
	blog.PengurusID = input.PengurusID
	if err := h.DB.Save(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// UpdateKategori updates the kategori of a blog by id
func (h *BlogHandler) UpdateKategori(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req struct {
		Kategori string `json:"kategori" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var blog model.Blog
	if err := h.DB.First(&blog, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}
	blog.Kategori = req.Kategori
	if err := h.DB.Save(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kategori updated", "blog": blog})
}

// ListByKategori returns blogs filtered by kategori
func (h *BlogHandler) ListByKategori(c *gin.Context) {
	kategori := c.Param("kategori")
	var blogs []model.Blog
	if err := h.DB.Where("kategori = ?", kategori).Order("created_at DESC").Find(&blogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "List blogs by kategori",
		"kategori": kategori,
		"blogs":    blogs,
	})
}

// Delete Blog
func (h *BlogHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.DB.Delete(&model.Blog{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
