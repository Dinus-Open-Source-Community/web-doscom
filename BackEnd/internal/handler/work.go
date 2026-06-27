package handler

import (
	"log"
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
	userID := c.MustGet("user_id").(int)

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
	if err := c.ShouldBind(&work); err != nil {
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
		userID,
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

func (h *WorkHandler) GetWorkTypes(c *gin.Context) {
	ctx := c.Request.Context()

	workTypes, err := h.Service.GetWorkTypes(ctx)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to fetch data, someting went wong",
			err,
		)
	}

	response.Success(
		c,
		"Successfully fetch work types",
		http.StatusOK,
		workTypes,
	)

}

func (h *WorkHandler) GetAllWorks(c *gin.Context) {
	ctx := c.Request.Context()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	filterProjectType := c.DefaultQuery("projecttype", "")

	worksResponseData, totalData, err := h.Service.GetAllWorksAndByProjectType(
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

	log.Printf("[handler] userRole: %s", userRole)
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

func (h *WorkHandler) GetWorkByIDPublic(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid id format",
			err,
		)
	}

	workResponseData, err := h.Service.GetWorkByID(ctx, "", id, true)
	if err != nil {
		response.Error(
			c,
			http.StatusNotFound,
			"something went wrong while fetching data",
			err,
		)
		return
	}

	responseData := &dto.WorkResponseClient{
		ID:           workResponseData.ID,
		Title:        workResponseData.Title,
		Tagline:      workResponseData.Tagline,
		Description:  workResponseData.Description,
		Slug:         workResponseData.Slug,
		ProjectType:  workResponseData.ProjectType,
		Technologies: workResponseData.Technologies,
		ProjectDate:  workResponseData.ProjectDate,
		ImageURL:     workResponseData.ImageURL,
		Gallery:      workResponseData.Gallery,
	}
	response.Success(
		c,
		"Success fetching work detail",
		http.StatusOK,
		responseData,
	)
}

func (h *WorkHandler) GetWorkByIDAdmin(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
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

	workResponseData, err := h.Service.GetWorkByID(ctx, userRole, id, false)
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

// update status work
func (h *WorkHandler) UpdateStatusWork(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)

	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleAdmin,
		constants.RoleKeyBPH,
	); err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"forbidden to access this resource",
			err,
		)
		return
	}

	idWork, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid request param, budy",
			err,
		)
		return
	}

	var workStatusUpdate dto.WorkUpdateStatusRequest
	if err := c.ShouldBindJSON(&workStatusUpdate); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid request body, budy",
			err,
		)
		return
	}

	workUpdatedDataResponse, err := h.Service.UpdateStatusWork(ctx, idWork, userID, workStatusUpdate.Status, userRole)
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

func (h *WorkHandler) UpdateWork(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)
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
			"invalid request param, budy",
			err,
		)
		return
	}

	var workDataToUpdate dto.WorkPatch
	if err := c.ShouldBind(&workDataToUpdate); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid request body, budy",
			err,
		)
		return
	}
	log.Printf("[handler] workDataToUpdate: %v", workDataToUpdate)

	form, _ := c.MultipartForm()
	var files []*multipart.FileHeader
	if form != nil {
		files = form.File["files"]
	}

	workUpdatedDataResponse, err := h.Service.UpdateWorkByID(
		ctx,
		id, userID,
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
		http.StatusOK,
		nil,
	)
}
