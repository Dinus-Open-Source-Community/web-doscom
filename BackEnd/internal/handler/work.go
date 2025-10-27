package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WorkHandler struct {
	DB *gorm.DB
}

func NewWorkHandler(db *gorm.DB) *WorkHandler {
	return &WorkHandler{DB: db}
}

func (h *WorkHandler) Create(c *gin.Context) {

	var work model.Work
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&work).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, work)

}

func (h *WorkHandler) List(c *gin.Context) {
	var works []model.Work
	if err := h.DB.Order("created_at DESC").Find(&works).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, works)
}

func (h *WorkHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var work model.Work
	if err := h.DB.First(&work, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "work not found"})
		return
	}
	c.JSON(http.StatusOK, work)
}

func (h *WorkHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var input model.Work
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var work model.Work
	if err := h.DB.First(&work, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "work not found"})
		return
	}
	work.Title = input.Title
	work.Description = input.Description
	work.ProjectDate = input.ProjectDate
	if err := h.DB.Save(&work).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, work)
}

func (h *WorkHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.DB.Delete(&model.Work{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
