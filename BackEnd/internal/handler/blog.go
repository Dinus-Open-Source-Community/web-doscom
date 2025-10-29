package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogHandler struct {
	DB *gorm.DB
}

func NewBlogHandler(db *gorm.DB) *BlogHandler {
	return &BlogHandler{DB: db}
}

// Create Blog
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
	if err := h.DB.Order("created_at DESC").Find(&blogs).Error; err != nil {
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
	blog.ActivityID = input.ActivityID
	blog.PengurusID = input.PengurusID
	if err := h.DB.Save(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) UpdateKategori(c *gin.Context) {
	// TODO: implement the logic for updating kategori
	c.JSON(200, gin.H{"message": "UpdateKategori not implemented"})
}
func (h *BlogHandler) ListByKategori(c *gin.Context) {
	kategori := c.Param("kategori")
	// TODO: Implement logic to list blogs by kategori
	c.JSON(200, gin.H{
		"message":  "List blogs by kategori",
		"kategori": kategori,
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
