package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogHandler struct {
	Model *model.BlogModel
}

func NewBlogHandler(db *model.BlogModel) *BlogHandler {
	return &BlogHandler{Model: db}
}

// Create Blog
func (h *BlogHandler) CreateBlog(c *gin.Context) {
	var input model.RegisterBlog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog := &model.Blog{
		Title:       input.Title,
		GalleryID:   input.GalleryID,
		Slug:        input.Slug,
		Content:     input.Content,
		Kategori:    input.Kategori,
		PublishedAt: input.PublishedAt,
		IsPublished: input.IsPublished,
		WorkID:      input.WorkID,
		ActivityID:  input.ActivityID,
		PengurusID:  input.PengurusID,
	}

	if err := h.DB.Create(blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
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
	if err := h.Model.Order("created_at DESC").Find(&blogs).Error; err != nil {
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
