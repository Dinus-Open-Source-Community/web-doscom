package model

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FileUpload struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           uint      `gorm:"not null" json:"user_id"`
	Category         string    `gorm:"type:varchar(50);not null" json:"category"`
	OriginalFilename string    `gorm:"type:varchar(255);not null" json:"original_filename"`
	StoredFilename   string    `gorm:"type:varchar(255);not null" json:"stored_filename"`
	FileSize         int64     `gorm:"not null" json:"file_size"`
	ContentType      string    `gorm:"type:varchar(100);not null" json:"content_type"`
	FileURL          string    `gorm:"type:text;not null" json:"file_url"`
	UploadedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"uploaded_at"`
	UpdatedAt        time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type FileUploadResponse struct {
	ID               uint   `json:"id"`
	UserID           uint   `json:"user_id"`
	Category         string `json:"category"`
	OriginalFilename string `json:"original_filename"`
	StoredFilename   string `json:"stored_filename"`
	FileSize         int64  `json:"file_size"`
	ContentType      string `json:"content_type"`
	FileURL          string `json:"file_url"`
}

type UploadFileRequest struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
	Folder     string
	UserID     uint
}

type FileUploadModel struct {
	DB *gorm.DB
}

// Create saves a new file upload record
func (m *FileUploadModel) CreateMetaData(fileUpload *FileUpload) (*FileUploadResponse, error) {
	if err := m.DB.Create(fileUpload).Error; err != nil {
		return nil, err
	}
	return &FileUploadResponse{
		ID:               fileUpload.ID,
		UserID:           fileUpload.UserID,
		Category:         fileUpload.Category,
		OriginalFilename: fileUpload.OriginalFilename,
		StoredFilename:   fileUpload.StoredFilename,
		FileSize:         fileUpload.FileSize,
		ContentType:      fileUpload.ContentType,
		FileURL:          fileUpload.FileURL,
	}, nil
}

// save a new file upload record
func (m *FileUploadModel) CreateMetaDataMultiple(ctx context.Context, fileUpload []*FileUpload) ([]*FileUploadResponse, error) {
	if err := m.DB.WithContext(ctx).Create(&fileUpload).Error; err != nil {
		return nil, err
	}

	response := make([]*FileUploadResponse, len(fileUpload))
	for i, fileData := range fileUpload {
		response[i] = &FileUploadResponse{
			ID:               fileData.ID,
			UserID:           fileData.UserID,
			Category:         fileData.Category,
			OriginalFilename: fileData.OriginalFilename,
			StoredFilename:   fileData.StoredFilename,
			FileSize:         fileData.FileSize,
			ContentType:      fileData.ContentType,
			FileURL:          fileData.FileURL,
		}
	}

	return response, nil
}

// GetByID retrieves a file upload by ID
func (m *FileUploadModel) GetByID(id uint) (*FileUpload, error) {
	var fileUpload FileUpload
	err := m.DB.Where("id = ?", id).First(&fileUpload).Error
	if err != nil {
		return nil, err
	}
	return &fileUpload, nil
}

// GetByUserID retrieves all file uploads for a user
func (m *FileUploadModel) GetByUserID(userID uint) ([]FileUpload, error) {
	var fileUploads []FileUpload
	err := m.DB.Where("user_id = ?", userID).Order("uploaded_at DESC").Find(&fileUploads).Error
	return fileUploads, err
}

// GetByCategory retrieves all file uploads in a category
func (m *FileUploadModel) GetByCategory(category string) ([]FileUpload, error) {
	var fileUploads []FileUpload
	err := m.DB.Where("category = ?", category).Order("uploaded_at DESC").Find(&fileUploads).Error
	return fileUploads, err
}

// GetByUserAndCategory retrieves file uploads for a user in a specific category
func (m *FileUploadModel) GetByUserAndCategory(userID uint, category string) ([]FileUpload, error) {
	var fileUploads []FileUpload
	err := m.DB.Where("user_id = ? AND category = ?", userID, category).Order("uploaded_at DESC").Find(&fileUploads).Error
	return fileUploads, err
}

func (m *FileUploadModel) UpdateFileUploadPartial(oldFileUrl string, data map[string]any) error {
	return m.DB.
		Model(&FileUpload{}).
		Where("file_url = ?", oldFileUrl).
		Updates(data).
		Error
}

func (m *FileUploadModel) UpdateFileUpload(userID int, fileUpload *FileUpload) (*FileUpload, error) {
	uploadData := &FileUpload{
		Category:         fileUpload.Category,
		OriginalFilename: fileUpload.OriginalFilename,
		StoredFilename:   fileUpload.StoredFilename,
		FileSize:         fileUpload.FileSize,
		ContentType:      fileUpload.ContentType,
		FileURL:          fileUpload.FileURL,
		UpdatedAt:        time.Now(),
	}

	var updatedMetadata FileUpload
	if err := m.DB.Model(&FileUpload{}).
		Clauses(clause.Returning{}).
		Where("user_id = ?", userID).
		Updates(uploadData).
		Scan(&updatedMetadata).
		Error; err != nil {
		return nil, err
	}

	return &updatedMetadata, nil
}

// Delete removes a file upload record
func (m *FileUploadModel) Delete(id uint) error {

	result := m.DB.Where("id = ?", id).Delete(&FileUpload{})

	if result.Error != nil {
		return fmt.Errorf("failed while deleted metadata file %w: ", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("failed, no data match to delete")
	}

	return nil
}

// DeleteByFilename removes a file upload record by stored filename
func (m *FileUploadModel) DeleteByFilename(filename string) error {
	return m.DB.Where("stored_filename = ?", filename).Delete(&FileUpload{}).Error
}
