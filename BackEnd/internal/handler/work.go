package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
)

type WorkHandler struct {
	Model *model.WorkModel
}

func NewWorkHandler(db *model.WorkModel) *WorkHandler {
	return &WorkHandler{Model: db}

}

func (h *WorkHandler) Create(c *gin.Context) {

	var work model.Work
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Model.InsertWork(&work).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, work)

}

func (h *WorkHandler) List(c *gin.Context) {
	work, err := h.Model.GetAllWorks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success gel all works, nasi padang satu bungkus",
		"works":   work,
	})
}

func (h *WorkHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	// var work model.Work
	work, err := h.Model.GetWorkById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success get work, NASI PADANG SATU BUNGKUS",
		"work":    work,
	})
}

func (h *WorkHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// req body data new user
	var patchData map[string]any
	if err := c.ShouldBindJSON(&patchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// update data
	updateDataWork, err := h.Model.UpdateWork(id, patchData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success update work, bwang nasi padangnya mana",
		"work":    updateDataWork,
	})

}

func (h *WorkHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Model.DeleteWork(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete data: " + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success delete data, info nasi padang bang",
	})
}
