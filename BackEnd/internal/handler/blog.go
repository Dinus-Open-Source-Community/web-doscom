package handler

import (
	"math"
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
<<<<<<< HEAD
// @Security ApiKeyAuth
// @Router /api/v1/blogs/ [post]
func (h *BlogHandler) CreateBlog(c *gin.Context) {
	user_role := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)
	ctx := c.Request.Context()

	if user_role != "KoorMedcrev" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are allowed but not allowed now",
		})
		return
	}

	var input model.RegisterBlog
	if err := c.ShouldBind(&input); err != nil {
=======
// @Security BearerAuth
// @Router /api/v1/blogs [post]
func (h *BlogHandler) Create(c *gin.Context) {
	var input model.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
>>>>>>> master
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert using service model
	if err := h.Service.Model.InsertBlog(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog: " + err.Error()})
		return
	}

<<<<<<< HEAD
	files := form.File["files"]

	blogInput := &model.RequestBlog{
		AuthorID:    userID,
		Title:       input.Title,
		Slug:        input.Slug,
		Content:     input.Content,
		Kategori:    input.Kategori,
		PublishedAt: input.PublishedAt,
		Status:      input.Status,
	}

	// service insert blog
	blogResponse, err := h.Service.CreateBlogImage(
		ctx,
		blogInput,
		input.ExistingID,
		files,
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
=======
	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created successfully",
		"blog":    input,
>>>>>>> master
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
<<<<<<< HEAD
func (h *BlogHandler) GetAllBlogs(c *gin.Context) {
	// role := c.MustGet("role").(string)
	ctx := c.Request.Context()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// service blog
	blogs, totalData, err := h.Service.GetAllBlogs(ctx, page, limit)
=======
func (h *BlogHandler) List(c *gin.Context) {
	blogs, err := h.Service.Model.GetAllBlogs()
>>>>>>> master
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
<<<<<<< HEAD
func (h *BlogHandler) GetBlogByID(c *gin.Context) {
=======
func (h *BlogHandler) Get(c *gin.Context) {
>>>>>>> master
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
<<<<<<< HEAD

	blog, err := h.Service.GetBlogByID(id)
=======
	blog, err := h.Service.Model.GetBlogById(id)
>>>>>>> master
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "blog not found",
		})
		return
	}

<<<<<<< HEAD
=======
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
	
>>>>>>> master
	c.JSON(http.StatusOK, gin.H{
		"message": "blog found successfully",
		"blog":    blog,
	})
}

<<<<<<< HEAD
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
	ctx := c.Request.Context()
	idBlog, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id blog"})
		return
	}

	userID := c.MustGet("user_id").(int)
	userRole := c.MustGet("role").(string)
	if userRole != "KoorMedcrev" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are not allowed to access this resource",
		})
		return
	}

	var dataPatch model.BlogPatch
	if err := c.ShouldBindJSON(&dataPatch); err != nil {
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

=======
>>>>>>> master
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
<<<<<<< HEAD

	allowedRoles := map[string]bool{
		"KoorMedcrev": true,
		"SuperAdmin":  true,
	}

	// check user
	if !allowedRoles[userRole] {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you're not allowed brotherr",
		})
=======
	if err := h.Service.Model.DeleteBlog(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
>>>>>>> master
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
