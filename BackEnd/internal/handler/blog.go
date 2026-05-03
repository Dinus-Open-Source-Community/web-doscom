package handler

import (
	"fmt"
	"math"
	"mime/multipart"
	"net/http"
	"strconv"

	"web_doscom/internal/constants"
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

func (h *BlogHandler) checkRolePermission(userRole string) error {

	_, err := constants.GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("role not valid %w", err)
	}

	allowedRole := map[string]bool{
		constants.RoleKeySuperAdmin:  true,
		constants.RoleKeyKoorMedcrev: true,
	}

	if !allowedRole[userRole] {
		return fmt.Errorf("role have no permission")
	}

	return nil
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
	ctx := c.Request.Context()
	user_role := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)

	if err := h.checkRolePermission(user_role); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "forbiddennnnn",
		})
		return
	}

	var input model.CreateRequestBlog
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

	blogInput := &model.BlogPayload{
		AuthorID:    userID,
		Content:     input.Content,
		ExistingID:  input.ExistingID,
		Kategori:    input.Kategori,
		PublishedAt: input.PublishedAt,
		Slug:        input.Slug,
		Status:      input.Status,
		Title:       input.Title,
	}

	// service insert blog
	blogResponse, err := h.Service.CreateBlogImage(
		ctx,
		blogInput,
		files,
		user_role,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to insert data",
		})
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully create blog",
		"data":    blogResponse,
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
func (h *BlogHandler) GetAllBlogs(c *gin.Context) {
	// role := c.MustGet("role").(string)
	ctx := c.Request.Context()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// service blog
	blogs, totalData, err := h.Service.GetAllBlogs(ctx, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to fetch blog data",
		})
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	currentPage := (offset / limit) + 1

	c.JSON(http.StatusOK, gin.H{
		"message":     "Succsess get all blogs",
		"data":        blogs,
		"totalPage":   totalPage,
		"currentPage": currentPage,
	})

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
func (h *BlogHandler) GetBlogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	blog, err := h.Service.GetBlogByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "blog not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "blog found successfully",
		"blog":    blog,
	})
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
	userID := c.MustGet("user_id").(int)
	userRole := c.MustGet("role").(string)
	if err := h.checkRolePermission(userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "forbiddennnnn",
		})
		return
	}

	ctx := c.Request.Context()
	idBlog, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id blog"})
		return
	}

	var dataPatch model.BlogPatch
	if err := c.ShouldBind(&dataPatch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, _ := c.MultipartForm()
	var files []*multipart.FileHeader
	if form != nil {
		files = form.File["files"]
	}

	// call service update blog
	blogResponse, err := h.Service.UpdateBlog(
		ctx,
		idBlog,
		userID,
		&dataPatch,
		files,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to update blog, something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully update blog",
		"data":    blogResponse,
	})
}

// get all blogs and get blog by kategory
func (h *BlogHandler) ListByKategori(c *gin.Context) {
	ctx := c.Request.Context()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	kategoriArray, exists := c.GetQueryArray("kategori")
	if exists && len(kategoriArray) > 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "max 3 kategory allowed",
		})
		return
	}

	var (
		blogs     []model.BlogThumbnail
		totalData int
		err       error
	)

	if exists {
		blogs, totalData, err = h.Service.GetBlogByKategori(
			ctx,
			kategoriArray,
			limit,
			offset,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed to fetch data",
			})
			return
		}
	} else {
		blogs, totalData, err = h.Service.GetAllBlogs(
			ctx,
			page,
			limit,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "failed to fetch data",
			})
			return
		}
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	currentPage := (offset / limit) + 1
	c.JSON(http.StatusOK, gin.H{
		"message":     "successfully fetch data",
		"blogs":       blogs,
		"totalPage":   totalPage,
		"currentPage": currentPage,
	})
}

// admin handler
func (h *BlogHandler) ListBlogs(c *gin.Context) {
	ctx := c.Request.Context()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("page", "10"))
	offset := (page - 1) * limit

	userRole := c.MustGet("role").(string)
	if err := h.checkRolePermission(userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "forbiddennnnn",
		})
		return
	}
	kategoriArray, exists := c.GetQueryArray("kategory")
	if exists && len(kategoriArray) > 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "max 3 kategory allowed",
		})
		return
	}

	// call service
	blogs, totalData, err := h.Service.GetAllBlogsForAdmin(
		ctx,
		kategoriArray,
		offset,
		limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "terjadi kesalahan ketika mengambil data",
		})
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	currentPage := (offset / limit) + 1

	c.JSON(http.StatusOK, gin.H{
		"message":     "successfully fetch data",
		"blogs":       blogs,
		"totalPage":   totalPage,
		"currentPage": currentPage,
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
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	if err := h.checkRolePermission(userRole); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "forbiddennnnn",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// call service deleteblogid
	if err := h.Service.DeleteBlogByID(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "terjadi kesalahan ketika menghapus data :)",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "successfully delete data",
	})
}
