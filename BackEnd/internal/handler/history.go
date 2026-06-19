package handler

import (
	"math"
	"net/http"
	"strconv"
	"web_doscom/internal/response"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	Service *service.HistoryService
}

func NewHistoryHandler(m *service.HistoryService) *HistoryHandler {
	return &HistoryHandler{Service: m}
}

func (h *HistoryHandler) GetHistoryTimelineByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid id",
			err,
		)
		return
	}

	history, err := h.Service.GetHistoryTimelineByID(ctx, id)
	if err != nil {
		response.Error(
			c,
			http.StatusNotFound,
			"history not found",
			err,
		)
	}

	response.Success(
		c,
		"history found successfully",
		http.StatusOK,
		history,
	)
}

func (h *HistoryHandler) GetAllHistoryTimeline(c *gin.Context) {
	ctx := c.Request.Context()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	history, totalData, err := h.Service.GetAllHistoryTimeline(ctx, offset, limit)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to fetch data",
			err,
		)
		return
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))
	currentPage := (offset / limit) + 1
	response.Success(
		c,
		"successfully fetch data",
		http.StatusOK,
		gin.H{
			"message":     "successfully fetch data",
			"history":     history,
			"totalPage":   totalPage,
			"currentPage": currentPage,
		},
	)
}
