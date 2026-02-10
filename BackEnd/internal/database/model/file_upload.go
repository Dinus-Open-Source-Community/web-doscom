package model

import (
	"time"

	"gorm.io/gorm"
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

type FileUploadModel struct {
	DB *gorm.DB
}

// Create saves a new file upload record
func (m *FileUploadModel) Create(fileUpload *FileUpload) error {
	return m.DB.Create(fileUpload).Error
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

// Delete removes a file upload record
func (m *FileUploadModel) Delete(id uint) error {
	return m.DB.Where("id = ?", id).Delete(&FileUpload{}).Error
}

// DeleteByFilename removes a file upload record by stored filename
func (m *FileUploadModel) DeleteByFilename(filename string) error {
	return m.DB.Where("stored_filename = ?", filename).Delete(&FileUpload{}).Error
}
