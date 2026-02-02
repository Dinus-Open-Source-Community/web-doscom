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
	blogs, err := h.Service.GetAllBlogs()
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
	blog, err := h.Service.GetBlogByID(id)
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
	updatedBlog, err := h.Service.UpdateBlog(id, patch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedBlog)
}

// UpdateKategori updates the kategori of a blog by id
func (h *BlogHandler) UpdateKategori(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req struct {
		Kategori string `json:"kategori" binding:"required,kategori"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patch := model.BlogPatch{Kategori: &req.Kategori}
	updatedBlog, err := h.Service.UpdateBlog(id, patch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kategori updated", "blog": updatedBlog})
}

// ListByKategori returns blogs filtered by kategori
func (h *BlogHandler) ListByKategori(c *gin.Context) {
	kategori := c.Param("kategori")
	blogs, err := h.Service.GetBlogsByKategori(kategori)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	if err := h.Service.DeleteBlog(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
