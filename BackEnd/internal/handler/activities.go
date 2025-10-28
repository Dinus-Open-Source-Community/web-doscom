package handler

import (
	"net/http"
	"strconv"
	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
)

type ActivitiesHandler struct {
	Model *model.ActivityModel
}

func NewActivitiesHandler(m *model.ActivityModel) *ActivitiesHandler {
	return &ActivitiesHandler{Model: m}
}

// CreateActivities godoc
// @Summary      Create new activities
// @Description  Membuat activities baru
// @Tags         Activities
// @Accept       json
// @Produce      json
// @Param        activities  body  model.Activities  true  "Activities info"
// @Success      201  {object}  model.Activities
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/activities [post]
func (h *ActivitiesHandler) CreateActivities(c *gin.Context) {
	var input model.Activities
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Model.InsertActivity(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create activities"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Activities created successfully", "activities": input})
}

// GetActivities godoc
// @Summary      Get activities by ID
// @Description  Mendapatkan activities berdasarkan ID
// @Tags         Activities
// @Produce      json
// @Param        id   path      int  true  "Activities ID"
// @Success      200  {object}  model.Activities
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Router      /api/v1/activities/{id} [get]
func (h *ActivitiesHandler) GetActivities(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid activities id"})
		return
	}
	act, err := h.Model.GetActivitiesById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activities Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get activities", "activities": act})
}

// GetAllActivities godoc
// @Summary      Get all activities
// @Description  Mendapatkan daftar semua activities
// @Tags         Activities
// @Produce      json
// @Success      200  {array}   model.Activities
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/activities [get]
func (h *ActivitiesHandler) GetAllActivities(c *gin.Context) {
	acts, err := h.Model.GetAllActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch activities data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "List of activities", "activities": acts})
}

// UpdateActivities godoc
// @Summary      Update activities
// @Description  Mengupdate data activities berdasarkan ID
// @Tags         Activities
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "Activities ID"
// @Param        activities  body      model.Activities  true  "Activities info"
// @Success      200   {object}  model.Activities
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/activities/{id} [put]
func (h *ActivitiesHandler) UpdateActivities(c *gin.Context) {
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
	act, err := h.Model.UpdateActivities(id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data activities"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully update activities data", "activities": act})
}

// DeleteActivities godoc
// @Summary      Delete activities
// @Description  Menghapus activities berdasarkan ID
// @Tags         Activities
// @Produce      json
// @Param        id   path  int  true  "Activities ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/activities/{id} [delete]
func (h *ActivitiesHandler) DeleteActivities(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Model.DeleteActivities(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activities not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "activities deleted"})
}
