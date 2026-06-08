package handler

import (
	"log"
	"net/http"
	"strconv"

	"web_doscom/internal/authorization"
	"web_doscom/internal/database/model/dto"
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
// function untuk create data pengurus lain
func (h *PengurusHandler) CreatePengurusProfile(c *gin.Context) {
	ctx := c.Request.Context()
	user_role := c.MustGet("role").(string)
	user_ID := c.MustGet("user_id").(int)

	var input dto.RegisterPengurusRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body: " + err.Error(),
		})
		return
	}

	// Get file from form
	fileHeader, err := c.FormFile("file")
	uploadFileData := new(dto.UploadFileRequest)
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to open file",
			})
			return
		}
		defer file.Close()

		uploadFileData = &dto.UploadFileRequest{
			FileHeader: fileHeader,
			File:       file,
			Folder:     "pengurus",
			UserID:     uint(user_ID),
		}
	}

	// insert data
	log.Printf("[InsertPengurus] terpanggil, user_id=%d", input.UserID)
	pengurusDataResponse, err := h.Service.CreatePengurus(
		ctx,
		user_ID,
		user_role,
		&input,
		uploadFileData,
	)
	if err != nil {
		// Handle specific conflict error
		if err.Error() == "email sudah terdaftar di pengurus" {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"message": "Gagal mendaftarkan pengurus: email sudah digunakan",
			})
			return
		}
		// Handle user not found or validation errors
		if err.Error() == "user_id tidak ditemukan" || err.Error() == "role not valid" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "Gagal mendaftarkan pengurus: input tidak valid",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to create data pengurus, server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Pengurus created successfully",
		"pengurus": pengurusDataResponse,
	})
}

// pengurus -> create pengurus profile
func (h *PengurusHandler) CreateMyPengurusProfile(c *gin.Context) {

	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)

	var input dto.RegisterPengurusRequest
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	// rewrite userID to current userID
	input.UserID = userID

	// Get file from form
	fileHeader, err := c.FormFile("file")
	uploadFileData := new(dto.UploadFileRequest)
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to open file",
			})
			return
		}

		defer file.Close()

		uploadFileData = &dto.UploadFileRequest{
			FileHeader: fileHeader,
			File:       file,
			Folder:     "pengurus",
			UserID:     uint(userID),
		}
	}

	pengurusDataResponse, err := h.Service.CreatePengurus(
		ctx,
		userID,
		userRole,
		&input,
		uploadFileData,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to create data pengurus, server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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
func (h *PengurusHandler) GetPengurusByUserID(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pengurus id"})
		return
	}
	log.Printf("[handler] id=%d", id)

	pengurusData, err := h.Service.GetPengurusByUserID(ctx, userRole, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "error while getting the data or data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Successfully get data",
		"pengurus data": pengurusData,
	})
}

func (h *PengurusHandler) GetPengurusByID(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pengurus id",
		})
		return
	}

	pengurusData, err := h.Service.GetPengurusByID(ctx, userRole, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "error while getting the data or data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Successfully get data",
		"pengurus data": pengurusData,
	})
}

func (h *PengurusHandler) GetPengurusProfile(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)

	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   err.Error(),
			"message": "failed to get data",
		})
		return
	}
	pengurusData, err := h.Service.GetPengurusByUserID(ctx, userRole, userID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":   err.Error(),
			"message": "failed to get data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"pengurus": pengurusData,
	})

}

func (h *PengurusHandler) GetAllPengurus(c *gin.Context) {
	ctx := c.Request.Context()
	divisi := c.Param("division")

	pengurusList, err := h.Service.GetAllPengurusByDivision(ctx, divisi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to fetch pengurus data",
		})
		return
	}

	pengurusDataResponse := make([]dto.PengurusPublicResponse, 0, len(pengurusList))
	for _, p := range pengurusList {
		pengurusDataResponse = append(pengurusDataResponse, dto.PengurusPublicResponse{
			ID:               p.ID,
			PhotoURL:         p.PhotoURL,
			Divisi:           p.Divisi,
			Name:             p.Name,
			Position:         p.Position,
			Sosmed:           p.Sosmed,
			StartPeriodeYear: p.StartPeriodeYear,
			EndPeriodeYear:   p.EndPeriodeYear,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "List of pengurus",
		"pengurus data": pengurusDataResponse,
	})
}

func (h *PengurusHandler) GetAllPengurusByDivision(c *gin.Context) {
	divisi := c.Query("divisi")
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)

	// call service GetAllPengurusbaseonddivision
	pengurusResponse, err := h.Service.GetPengurusBaseOnDivision(
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
func (h *PengurusHandler) UpdatePengurusByID(c *gin.Context) {
	ctx := c.Request.Context()
	idCurrentUser := c.MustGet("user_id").(int)
	currentUserRole := c.MustGet("role").(string)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var patch dto.PengurusPatch
	if err := c.ShouldBind(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind data, please use form-data for updates"})
		return
	}

	fileheader, err := c.FormFile("file")
	var fileUpload *dto.UploadFileRequest
	if err == nil {
		file, err := fileheader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to open file",
			})
			return
		}
		defer file.Close()

		// file request to upload to MinIO
		fileUpload = &dto.UploadFileRequest{
			FileHeader: fileheader,
			File:       file,
			Folder:     "pengurus",
			UserID:     uint(id),
		}
	}
	// update data pengurus
	updatedPengurus, err := h.Service.UpdateDataPengurus(
		ctx,
		id,
		idCurrentUser,
		currentUserRole,
		&patch,
		fileUpload,
		false,
	)
	if err != nil {
		// Handle permission issues
		if err.Error() == "You are not allowed to update this data" || err.Error() == "koordinator tidak dapat memperbarui foto pengurus" {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   err.Error(),
				"message": "Akses ditolak untuk memperbarui data ini",
			})
			return
		}

		// Handle record not found
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Pengurus data not found",
			})
			return
		}

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

func (h *PengurusHandler) UpdateMyPengurusProfile(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("user_id").(int)
	UserRole := c.MustGet("role").(string)

	_, err := authorization.GetRoleInfo(UserRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "role not valid",
		})
		return
	}

	var input dto.PengurusPatch
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	var fileUpload *dto.UploadFileRequest
	fileHeader, err := c.FormFile("file")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read file",
			})
			return
		}
		defer file.Close()

		fileUpload = &dto.UploadFileRequest{
			FileHeader: fileHeader,
			File:       file,
			Folder:     "pengurus",
			UserID:     uint(userID),
		}
	}

	updatedPengurus, err := h.Service.UpdateDataPengurus(
		ctx,
		userID,
		userID,
		UserRole,
		&input,
		fileUpload,
		true,
	)
	if err != nil {

		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Pengurus data not found",
			})
			return
		}
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
	userRole := c.MustGet("role").(string)
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
