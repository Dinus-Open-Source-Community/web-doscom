package handler

import (
	"net/http"
	"strconv"

	"web_doscom/internal/database/model"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

type PengurusHandler struct {
<<<<<<< HEAD
=======
	// Model          *model.PengurusModel
	// GalleryService *service.GalleryService
>>>>>>> master
	Service        *service.PengurusService
	StorageService *service.StorageService
}

func NewPengurusHandler(pengurusService *service.PengurusService, storageService *service.StorageService) *PengurusHandler {
	return &PengurusHandler{
		Service:        pengurusService,
		StorageService: storageService,
<<<<<<< HEAD
=======
	}
}

// set position by role
func (PengurusHandler) SetPositionByrole(position string) (string, error) {
	// set Divisi
	var validPosition string
	if model.ValidPosition[position] {
		validPosition = position
	} else {
		return "", fmt.Errorf("Role not valid, koe sopo cok")
>>>>>>> master
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
<<<<<<< HEAD
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
=======
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer file.Close()

	// Upload to MinIO with pengurus category
	fileURL, err := h.StorageService.UploadFile(
		c.Request.Context(),
		file,
		fileHeader,
		"pengurus",
		uint(user_ID),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upload file: " + err.Error(),
		})
		return
	}

	// Auto-assign position for kor and anggota
	position, err := h.SetPositionByrole(input.Position)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
>>>>>>> master
		})
		return
	}

<<<<<<< HEAD
=======
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
				"error": "User_id & Email must be given",
			})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user_id",
		})
		return
	}

	// Handle sosmed field - database constraint requires specific values
	sosmedValue := input.Sosmed
	if sosmedValue == "" {
		sosmedValue = "instagram" // Default value to satisfy database constraint
	}

	pengurus := &model.Pengurus{
		UserID:   input.UserID,
		URLAsset: fileURL,
		Email:    input.Email,
		Divisi:   input.Divisi,
		Name:     input.Name,
		Position: position,
		Sosmed:   sosmedValue,
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

>>>>>>> master
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
<<<<<<< HEAD
			PhotoURL: p.PhotoURL,
=======
			URLAsset: p.URLAsset,
>>>>>>> master
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
