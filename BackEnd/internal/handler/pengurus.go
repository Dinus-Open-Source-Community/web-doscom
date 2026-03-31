package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

type PengurusHandler struct {
	Service        *service.PengurusService
	StorageService *service.StorageService
}

func NewPengurusHandler(pengurusService *service.PengurusService, storageService *service.StorageService) *PengurusHandler {
	return &PengurusHandler{
		Service:        pengurusService,
		StorageService: storageService,
	}
}

// CreatePengurus godoc
// @Summary Create new pengurus
// @Description Create pengurus with profile picture upload to MinIO. All files uploaded to MinIO bucket.
// @Tags Pengurus
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Profile Picture (image file)"
// @Param id_user formData int false "User ID (optional, auto-filled for non-admin)"
// @Param email formData string false "Email (optional, auto-filled for non-admin)"
// @Param divisi formData string true "Divisi - Valid values: bph, pemro, jaringan, medcrev, data"
// @Param name formData string true "Name (2-150 characters)"
// @Param position formData string true "Position - Valid values: ketum, sdm, pr, pm, pm_ang, sekum, bendum, sek_ang, ben_ang, kor_pemro, kor_jaringan, kor_medcrev, kor_data, anggota, pemro_ang, jaringan_ang, medcrev_ang, data_ang"
// @Param sosmed formData string false "Social Media Platform (optional) - Valid values: instagram, linkedin, github. Leave empty for default 'instagram'"
// @Param period formData string true "Period (YYYY-MM-DD format)"
// @Success 201 {object} model.PengurusResponse
// @Failure 400 {object} map[string]string "Validation error - check divisi, position, or sosmed values"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /api/v1/pengurus/ [post]
func (h *PengurusHandler) CreatePengurus(c *gin.Context) {
	user_role := c.MustGet("role").(string)
	user_ID := c.MustGet("user_id").(int)
	// email_user := c.MustGet("email").(string)
	ctx := c.Request.Context()

	var input model.RegisterPengurusRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body: " + err.Error(),
		})
		return
	}

	// Get file from form
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer file.Close()

	uploadFile := &model.UploadFileRequest{
		FileHeader: fileHeader,
		File:       file,
		Folder:     "pengurus",
		UserID:     uint(user_ID),
	}

	// insert data
	pengurusDataResponse, err := h.Service.CreatePengurus(
		ctx,
		user_ID,
		user_role,
		&input,
		uploadFile,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to create data pengurus, server error",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Pengurus created successfully",
		"pengurus": pengurusDataResponse,
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
func (h *PengurusHandler) GetPengurusByID(c *gin.Context) {
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
		PhotoURL: pengurus.PhotoURL,
		Email:    pengurus.Email,
		Divisi:   pengurus.Divisi,
		Name:     pengurus.Name,
		Position: pengurus.Position,
		Sosmed:   pengurus.Sosmed,
		Period:   pengurus.Period,
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully get data",
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
	divisi := c.Param("divisi")

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
			PhotoURL: p.PhotoURL,
			Email:    "",
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

func (h *PengurusHandler) GetAllPengurusByDivision(c *gin.Context) {
	divisi := c.Query("divisi")
	ctx := c.Request.Context()
	userRole := c.MustGet("user_role").(string)

	// call service GetAllPengurusbaseonddivision
	pengurusResponse, err := h.Service.GetAllPengurusBaseOnDivision(
		ctx,
		userRole,
		divisi,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to get data somting wong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully get data",
		"data":    pengurusResponse,
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
	ctx := c.Request.Context()
	idCurrentUser := c.MustGet("id").(int)
	currentUserRole := c.MustGet("role").(string)

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

	fileheader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	file, err := fileheader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open file",
		})
		return
	}

	defer file.Close()

	// file request to upload to MinIO
	fileUpload := &model.UploadFileRequest{
		FileHeader: fileheader,
		File:       file,
		Folder:     "pengurus",
		UserID:     uint(id),
	}
	// update data pengurus
	updatedPengurus, err := h.Service.UpdateDataPengurus(
		ctx,
		id,
		idCurrentUser,
		currentUserRole,
		&patch,
		fileUpload,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to update data pengurus, server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully update pengurus data",
		"pengurus": updatedPengurus,
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
	userRole := c.MustGet("user_role").(string)
	ctx := c.Request.Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	if err := h.Service.DeletePengurusById(ctx, id, userRole); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "pengurus not Found",
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "pengurus deleted",
	})
}
