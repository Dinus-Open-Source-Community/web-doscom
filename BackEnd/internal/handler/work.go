package handler

import (
	"math"
	"mime/multipart"
	"net/http"
	"strconv"

	"web_doscom/internal/authorization"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/response"

	// "web_doscom/internal/database/model"
	"web_doscom/internal/service"

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
		response.Error(
			c,
			http.StatusForbidden,
			"role is not allowed to access this resource",
			err,
		)
		return
	}

	var work dto.CreateRequestWork
	if err := c.ShouldBindJSON(&work); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			err,
		)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"failed to read file",
			err,
		)
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
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to insert data, something went wrong",
			err,
		)
		return
	}

	response.Success(
		c,
		"Work created successfully",
		http.StatusCreated,
		workInputDataResponse,
	)
}

func (h *WorkHandler) GetAllWorks(c *gin.Context) {
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
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to fetch data, something went wrong",
			err,
		)
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	currentPage := (offset / limit) * 1

	response.Success(
		c,
		"Success fetching all works",
		http.StatusOK,
		gin.H{
			"work data":   worksResponseData,
			"totalPage":   totalPage,
			"currentPage": currentPage,
		},
	)
}

func (h *WorkHandler) GetAllWorksByDivision(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role is not allowed to access this resource",
			err,
		)
		return
	}

	worksResponseData, totalData, err := h.Service.GetWorksByDivision(
		ctx,
		userRole,
		limit,
		offset,
	)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to fetch data, something went wrong",
			err,
		)
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	currentPage := (offset / limit) * 1

	response.Success(
		c,
		"Success fetching all works",
		http.StatusOK,
		gin.H{
			"worksData":   worksResponseData,
			"totalPage":   totalPage,
			"currentPage": currentPage,
		},
	)

}

func (h *WorkHandler) GetWorkByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid id format",
			err,
		)
		return
	}

	workResponseData, err := h.Service.GetWorkByID(ctx, id)
	if err != nil {
		response.Error(
			c,
			http.StatusNotFound,
			"something went wrong while fetching data",
			err,
		)
		return
	}

	response.Success(
		c,
		"Success fetching work detail",
		http.StatusOK,
		workResponseData,
	)
}

func (h *WorkHandler) UpdateWork(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleKoordinator,
		constants.RoleAdmin,
	); err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"forbidden to access this resource",
			err,
		)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid request body, budy",
			err,
		)
		return
	}

	var workDataToUpdate dto.WorkPatch
	if err := c.ShouldBindJSON(&workDataToUpdate); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid request body, budy",
			err,
		)
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
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to update work, something went wrong",
			err,
		)
		return
	}

	response.Success(
		c,
		"Work updated successfully",
		http.StatusOK,
		workUpdatedDataResponse,
	)
}

func (h *WorkHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	if err := authorization.CheckRolePermission(userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"forbidden to access this resource",
			err,
		)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid id format",
			err,
		)
		return
	}

	if err := h.Service.DeleteWork(ctx, id, userRole); err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to delete work, something went wrong",
			err,
		)
		return
	}

	response.Success(
		c,
		"Work deleted successfully",
		http.StatusNoContent,
		nil,
	)
}
