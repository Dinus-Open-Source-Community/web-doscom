package handler

import (
	"net/http"
	"strconv"
	"time"

	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActivitiesHandler struct {
	DB *gorm.DB
}

func NewActivitiesHandler(db *gorm.DB) *ActivitiesHandler {
	return &ActivitiesHandler{DB: db}
}

func (h *ActivitiesHandler) Create(c *gin.Context) {
	var act model.Activities
	if err := c.ShouldBindJSON(&act); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Optional: handle empty ActivitiesDate
	if act.ActivitiesDate.IsZero() {
		act.ActivitiesDate = time.Now()
	}
	if err := h.DB.Create(&act).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, act)
}

func (h *ActivitiesHandler) List(c *gin.Context) {
	var acts []model.Activities
	if err := h.DB.Order("created_at DESC").Find(&acts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, acts)
}

func (h *ActivitiesHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var act model.Activities
	if err := h.DB.First(&act, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activities not found"})
		return
	}
	c.JSON(http.StatusOK, act)
}

func (h *ActivitiesHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input model.Activities
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var act model.Activities
	if err := h.DB.First(&act, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activities not found"})
		return
	}
	act.ActivitiesTitle = input.ActivitiesTitle
	act.ActivitiesDesc = input.ActivitiesDesc
	act.ActivitiesDate = input.ActivitiesDate
	if err := h.DB.Save(&act).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, act)
}

func (h *ActivitiesHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.DB.Delete(&model.Activities{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
