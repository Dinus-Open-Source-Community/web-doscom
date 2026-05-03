package handler

import (
	"math"
	"mime/multipart"
	"net/http"
	"strconv"

	"web_doscom/internal/authorization"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"
	"web_doscom/internal/service"
	"web_doscom/internal/utils"

	"github.com/gin-gonic/gin"
)

type WorkHandler struct {
	Service *service.WorkService
}

func NewWorkHandler(s *service.WorkService) *WorkHandler {
	return &WorkHandler{Service: s}
}

func (h *WorkHandler) CreateWork(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)

	if err := authorization.CheckRolePermission(userRole,
		constants.RoleKoordinator,
		constants.RoleAdmin,
	); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "role is not allowed to access this resource",
		})
		return
	}

	var work model.CreateRequestWork
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "failed to read file",
		})
		return
	}

	files := form.File["files"]

	workInputDataResponse, err := h.Service.CreateWork(
		ctx,
		&work,
		files,
		userRole,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to insert data, something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Work created successfully",
		"data":    workInputDataResponse,
	})
}

func (h *WorkHandler) GetAllWorksOrByFilterTechnologies(c *gin.Context) {
	ctx := c.Request.Context()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	filterProjectType := c.DefaultQuery("projecttype", "")
	offset := (page - 1) * limit

	worksResponseData, totalData, err := h.Service.GetAllWorksAndByTechnologi(
		ctx,
		offset,
		limit,
		filterProjectType,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to fetch data, something went wrong",
		})
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	currentPage := (offset / limit) * 1

	c.JSON(http.StatusOK, gin.H{
		"message":     "Success fetching all works",
		"work data":   worksResponseData,
		"totalPage":   totalPage,
		"currentPage": currentPage,
	})
}

func (h *WorkHandler) GetWorkByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	workResponseData, err := h.Service.GetWorkByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "something went wrong while fetching data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success fetching work detail",
		"data":    workResponseData,
	})
}

func (h *WorkHandler) UpdateWorkByID(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleKoordinator,
		constants.RoleAdmin,
	); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "forbidden to access this resource",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format",
		})
		return
	}

	var workDataToUpdate model.WorkPatch
	if err := c.ShouldBindJSON(&workDataToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "invalid request body, budy",
		})
		return
	}

	form, _ := c.MultipartForm()
	var files []*multipart.FileHeader
	if form != nil {
		files = form.File["files"]
	}

	workUpdatedDataResponse, err := h.Service.UpdateWorkByID(
		ctx,
		id,
		&workDataToUpdate,
		files,
		userRole,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to update work, something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Work updated successfully",
		"data":    workUpdatedDataResponse,
	})
}

func (h *WorkHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	if err := authorization.CheckRolePermission(userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "forbidden to access this resource",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if err := h.Service.DeleteWork(ctx, id, userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to delete work, something went wrong",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "Work deleted successfully",
	})
}
