package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
)

type PengurusHandler struct {
	Model *model.PengurusModel
}

func NewPengurusHandler(m *model.PengurusModel) *PengurusHandler {
	return &PengurusHandler{Model: m}
}

// CreatePengurus godoc
func (h *PengurusHandler) CreatePengurus(c *gin.Context) {
	var input model.RegisterPengurusRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read req body"})
		return
	}

	// Auto-assign kor position for Admin
	var role string
	row := h.Model.DB.Raw("SELECT role FROM users WHERE id = ?", input.UserID).Row()
	row.Scan(&role)
	if role == "Admin" {
		switch input.Divisi {
		case "pemro":
			input.Position = "kor_pemro"
		case "jaringan":
			input.Position = "kor_jaringan"
		case "medcrev":
			input.Position = "kor_medcrev"
		case "data":
			input.Position = "kor_data"
		case "bph":
			input.Position = "sekum,bendum,pm,sdm,pr"
		}
	}

	pengurus := &model.Pengurus{
		UserID:   input.UserID,
		URLAsset: input.URLAsset,
		Email:    input.Email,
		Divisi:   input.Divisi,
		Name:     input.Name,
		Position: input.Position,
		Sosmed:   input.Sosmed,
		Period:   input.Period,
	}
	if err := h.Model.InsertPengurus(pengurus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := model.PengurusResponse{
		ID:       pengurus.ID,
		UserID:   pengurus.UserID,
		URLAsset: pengurus.URLAsset,
		Email:    pengurus.Email,
		Divisi:   pengurus.Divisi,
		Name:     pengurus.Name,
		Position: pengurus.Position,
		Sosmed:   pengurus.Sosmed,
		Period:   pengurus.Period,
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Pengurus created successfully", "pengurus": resp})
}

// GetPengurus godoc
func (h *PengurusHandler) GetPengurus(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pengurus id"})
		return
	}
	pengurus, err := h.Model.GetPengurusById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengurus Not Found"})
		return
	}
	resp := model.PengurusResponse{
		ID:       pengurus.ID,
		UserID:   pengurus.UserID,
		URLAsset: pengurus.URLAsset,
		Email:    pengurus.Email,
		Divisi:   pengurus.Divisi,
		Name:     pengurus.Name,
		Position: pengurus.Position,
		Sosmed:   pengurus.Sosmed,
		Period:   pengurus.Period,
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get pengurus", "pengurus": resp})
}

// GetAllPengurus godoc
func (h *PengurusHandler) GetAllPengurus(c *gin.Context) {
	pengurusList, err := h.Model.GetAllPengurus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pengurus data"})
		return
	}
	var respList []model.PengurusResponse
	for _, p := range pengurusList {
		respList = append(respList, model.PengurusResponse{
			ID:       p.ID,
			UserID:   p.UserID,
			URLAsset: p.URLAsset,
			Email:    p.Email,
			Divisi:   p.Divisi,
			Name:     p.Name,
			Position: p.Position,
			Sosmed:   p.Sosmed,
			Period:   p.Period,
		})
	}
	c.JSON(http.StatusOK, gin.H{"message": "List of pengurus", "pengurus": respList})
}

// UpdatePengurus godoc
func (h *PengurusHandler) UpdatePengurus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var patch model.PengurusPatch
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read req body"})
		return
	}
	updatePengurus, err := h.Model.UpdatePengurus(id, patch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update data pengurus"})
		return
	}
	resp := model.PengurusResponse{
		ID:       updatePengurus.ID,
		UserID:   updatePengurus.UserID,
		URLAsset: updatePengurus.URLAsset,
		Email:    updatePengurus.Email,
		Divisi:   updatePengurus.Divisi,
		Name:     updatePengurus.Name,
		Position: updatePengurus.Position,
		Sosmed:   updatePengurus.Sosmed,
		Period:   updatePengurus.Period,
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully update pengurus data", "pengurus": resp})
}

// DeletePengurus godoc
func (h *PengurusHandler) DeletePengurus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.Model.DeletePengurus(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pengurus not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "pengurus deleted"})
}
