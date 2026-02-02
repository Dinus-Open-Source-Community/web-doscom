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

func NewBlogHandler(service *service.BlogService) *BlogHandler {
	return &BlogHandler{Service: service}
}

// Create Blog godoc
// @Summary Create new blog
// @Description Create a new blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param blog body model.RegisterBlog true "Blog data"
// @Success 201 {object} model.Blog
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/blogs [post]
func (h *BlogHandler) Create(c *gin.Context) {
	var input model.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert using service model
	if err := h.Service.Model.InsertBlog(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"blog":    input,
	})
}

// List all Blogs godoc
// @Summary List all blogs
// @Description Get all blog posts
// @Tags Blog
// @Produce json
// @Success 200 {array} model.Blog
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/blogs [get]
func (h *BlogHandler) List(c *gin.Context) {
	blogs, err := h.Service.Model.GetAllBlogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// Get Blog by ID godoc
// @Summary Get blog by ID
// @Description Get detailed information of a blog
// @Tags Blog
// @Produce json
// @Param id path int true "Blog ID"
// @Success 200 {object} model.Blog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/blogs/{id} [get]
func (h *BlogHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	blog, err := h.Service.Model.GetBlogById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// Update Blog godoc
// @Summary Update blog
// @Description Update existing blog post
// @Tags Blog
// @Accept json
// @Produce json
// @Param id path int true "Blog ID"
// @Param blog body model.BlogPatch true "Update data"
// @Success 200 {object} model.Blog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/blogs/{id} [put]
func (h *BlogHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var patch model.BlogPatch
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	blog, err := h.Service.Model.UpdateBlog(id, patch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// ListByKategori godoc
// @Summary List blogs by category
// @Description Get all blogs specific to a category
// @Tags Blog
// @Produce json
// @Param kategori path string true "Category name"
// @Success 200 {array} model.Blog
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/blogs/kategori/{kategori} [get]
func (h *BlogHandler) ListByKategori(c *gin.Context) {
	kategori := c.Param("kategori")
	blogs, err := h.Service.Model.GetBlogsByKategori(kategori)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs: " + err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":  "List blogs by kategori",
		"kategori": kategori,
		"blogs":    blogs,
	})
}

// Delete Blog godoc
// @Summary Delete blog
// @Description Delete a blog post
// @Tags Blog
// @Param id path int true "Blog ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/v1/blogs/{id} [delete]
func (h *BlogHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Service.Model.DeleteBlog(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
