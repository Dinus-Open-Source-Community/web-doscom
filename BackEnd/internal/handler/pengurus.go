package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"web_doscom/internal/database/model"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

type PengurusHandler struct {
	// Model          *model.PengurusModel
	// GalleryService *service.GalleryService
	Service *service.PengurusService
}

func NewPengurusHandler(m *service.PengurusService) *PengurusHandler {
	return &PengurusHandler{Service: m}
}

// set position by role
func (PengurusHandler) SetPositionByrole(position string) (string, error) {
	// set Divisi
	var validPosition string
	if model.ValidPosition[position] {
		validPosition = position
	} else {
		return "", fmt.Errorf("Role not valid, koe sopo cok")
	}
	return validPosition, nil
}

// CreatePengurus godoc
// @Summary Create new pengurus
// @Description Create pengurus using multipart/form-data (JSON fields + file upload)
// @Tags Pengurus
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Profile Picture"
// @Param id_user formData int false "User ID (optional)"
// @Param email formData string false "Email (optional)"
// @Param divisi formData string true "Divisi (required)"
// @Param name formData string true "Name (required)"
// @Param position formData string true "Position (required)"
// @Param sosmed formData string false "Sosmed (optional)"
// @Param period formData string true "Period (YYYY-MM-DD)"
// @Success 201 {object} model.PengurusResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/v1/pengurus/ [post]
func (h *PengurusHandler) CreatePengurus(c *gin.Context) {
	user_role := c.MustGet("role").(string)
	user_ID := c.MustGet("user_id").(int)
	email_user := c.MustGet("email").(string)

	var input model.RegisterPengurusRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body: " + err.Error(),
		})
		return
	}
	// save photo to gallery and insert to database
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})
		return
	}
	profilePic, err := h.Service.GalleryModel.UploadInsertSingleImage(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Auto-assign assign position for kor and anggota
	position, err := h.SetPositionByrole(input.Position) // masih issue tapi normal tapi issue tapi normal
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// check user_id
	switch {
	case strings.HasPrefix(user_role, "_ang"):
		input.UserID = user_ID
		input.Email = email_user
	case strings.HasPrefix(user_role, "Kor_"), strings.HasPrefix(user_role, "BPH"):
		if input.UserID == 0 && input.Email == "" {
			input.UserID = user_ID
			input.Email = email_user
		}
	case strings.HasPrefix(user_role, "Super_Admin"):
		if input.UserID == 0 && input.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "User_id & Email must be given, datane sopo kii wok",
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user_id, koe sopo wok",
		})
		return
	}

	pengurus := &model.Pengurus{
		UserID:   input.UserID,
		URLAsset: profilePic.AssetUrl,
		Email:    input.Email,
		Divisi:   input.Divisi,
		Name:     input.Name,
		Position: position,
		Sosmed:   input.Sosmed,
		Period:   input.Period,
	}
	if err := h.Service.PengurusModel.InsertPengurus(pengurus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := model.PengurusResponse{
		ID:       pengurus.ID,
		URLAsset: pengurus.URLAsset,
		Email:    pengurus.Email,
		Divisi:   pengurus.Divisi,
		Name:     pengurus.Name,
		Position: pengurus.Position,
		Sosmed:   pengurus.Sosmed,
		Period:   pengurus.Period,
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":  "Pengurus created successfully",
		"pengurus": resp,
	})
}

// GetPengurus godoc
// @Summary Get Pengurus by ID
// @Description Mengambil data pengurus berdasarkan ID
// @Accept json
// @Produce json
// @Param id path int true "Pengurus ID"
// @Success 200 {object} model.PengurusResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Tags Pengurus
// @Router /api/v1/pengurus/{id} [get]
func (h *PengurusHandler) GetPengurus(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pengurus id"})
		return
	}
	pengurus, err := h.Service.PengurusModel.GetPengurusById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengurus Not Found"})
		return
	}
	resp := model.PengurusResponse{
		ID:       pengurus.ID,
		URLAsset: pengurus.URLAsset,
		Email:    pengurus.Email,
		Divisi:   pengurus.Divisi,
		Name:     pengurus.Name,
		Position: pengurus.Position,
		Sosmed:   pengurus.Sosmed,
		Period:   pengurus.Period,
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Get pengurus",
		"pengurus": resp,
	})
}

// GetAllPengurus godoc
// @Summary Get all Pengurus
// @Description Mengambil semua data penguru
// @Accept json
// @Produce json
// @Param divisi query string false "Filter by divisi (optional)"
// @Success 200 {object} model.PengurusResponse
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Tags Pengurus
// @Router /api/v1/pengurus/ [get]
func (h *PengurusHandler) GetAllPengurus(c *gin.Context) {
	// get role
	role := c.MustGet("role").(string)
	divisi := c.Query("divisi")
	var isValid = map[string]bool{
		"Super_Admin":  true,
		"Kor_Pemro":    true,
		"Kor_Jaringan": true,
		"Kor_Medcrev":  true,
		"Kor_Data":     true,
		"BPH":          true,
	}
	if !isValid[role] {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Heyy siapa kau, tidak boleh akses data ini",
		})
		return
	}

	pengurusList, err := h.Service.PengurusModel.GetAllPengurusByDivisi(divisi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch pengurus data",
		})
		return
	}
	var respList []model.PengurusResponse
	for _, p := range pengurusList {
		respList = append(respList, model.PengurusResponse{
			ID:       p.ID,
			URLAsset: p.URLAsset,
			Email:    p.Email,
			Divisi:   p.Divisi,
			Name:     p.Name,
			Position: p.Position,
			Sosmed:   p.Sosmed,
			Period:   p.Period,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "List of pengurus",
		"pengurus": respList,
	})
}

// UpdatePengurus godoc
// @Summary Update Pengurus
// @Description Update data pengurus berdasarkan ID
// @Accept json
// @Produce json
// @Param id path int true "Pengurus ID"
// @Param pengurus body model.PengurusPatch true "Pengurus info"
// @Success 200 {object} model.PengurusResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Tags Pengurus
// @Router /api/v1/pengurus/{id} [put]
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
	updatePengurus, err := h.Service.PengurusModel.UpdatePengurus(id, patch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update data pengurus",
		})
		return
	}
	resp := model.PengurusResponse{
		ID:       updatePengurus.ID,
		URLAsset: updatePengurus.URLAsset,
		Email:    updatePengurus.Email,
		Divisi:   updatePengurus.Divisi,
		Name:     updatePengurus.Name,
		Position: updatePengurus.Position,
		Sosmed:   updatePengurus.Sosmed,
		Period:   updatePengurus.Period,
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully update pengurus data",
		"pengurus": resp,
	})
}

// DeletePengurus godoc
// @Summary Delete Pengurus
// @Description Delete data pengurus berdasarkan ID
// @Accept json
// @Produce json
// @Param id path int true "Pengurus ID"
// @Success 200 {object} model.PengurusResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Tags Pengurus
// @Router /api/v1/pengurus/{id} [delete]
func (h *PengurusHandler) DeletePengurus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	if err := h.Service.PengurusModel.DeletePengurus(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "pengurus not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pengurus deleted",
	})
}
