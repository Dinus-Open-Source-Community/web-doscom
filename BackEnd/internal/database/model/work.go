package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WorkModel struct {
	DB *gorm.DB
}

type Work struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	PengurusID   int       `gorm:"column:pengurus_id" json:"pengurus_id"`
	Title        string    `gorm:"column:title" json:"title"`
	Tagline      string    `gorm:"column:tagline" json:"tagline"`
	Description  string    `gorm:"column:description" json:"description"`
	Slug         string    `gorm:"column:slug" json:"slug"`
	ProjectType  string    `gorm:"column:project_type" json:"project_type"`
	Technologies []string  `gorm:"type:text[]" json:"technologies"`
	ProjectDate  time.Time `gorm:"column:project_date" json:"project_date"`
	ImageURL     string    `gorm:"column:image_url" json:"image_url"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type RegisterWork struct {
	ExistingID   []*int    `form:"existingID_image"`
	PengurusID   int       `form:"pengurus_id" binding:"required"`
	Title        string    `form:"title" binding:"required"`
	Tagline      string    `form:"tagline"`
	Description  string    `form:"description" binding:"required"`
	Slug         string    `form:"slug" binding:"required"`
	ProjectType  string    `form:"project_type" binding:"required"`
	Technologies []string  `form:"technologies[]" binding:"required"`
	ProjectDate  time.Time `form:"project_date" binding:"required"`
}

type WorkResponse struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Tagline      string    `json:"tagline"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ProjectType  string    `json:"project_type"`
	Technologies []string  `json:"technologies"`
	ProjectDate  time.Time `json:"project_date"`
	ImageURL     string    `json:"image_url"`
}
type WorkPatch struct {
	PengurusID   int       `json:"pengurus_id" `
	Title        string    `json:"title" `
	Tagline      string    `json:"tagline"`
	Description  string    `json:"description" `
	Slug         string    `json:"slug" `
	ProjectType  string    `json:"project_type" `
	Technologies []string  `json:"technologies" `
	ProjectDate  time.Time `json:"project_date" `
}

func (Work) TableName() string {
	return "work"
}

func (m *WorkModel) WithTx(tx *gorm.DB) *WorkModel {
	return &WorkModel{DB: tx}
}

func (m *WorkModel) InsertWork(ctx context.Context, work *Work) (WorkResponse, error) {
	work.CreatedAt = time.Now()
	work.UpdatedAt = time.Now()
	result := m.DB.WithContext(ctx).Create(work)

	if result.Error != nil {
		return WorkResponse{}, fmt.Errorf("failed while insert data %w", result.Error)
	}

	responseData := WorkResponse{
		ID:           work.ID,
		Title:        work.Title,
		Tagline:      work.Tagline,
		Description:  work.Description,
		Slug:         work.Slug,
		ProjectType:  work.ProjectType,
		Technologies: work.Technologies,
		ProjectDate:  work.ProjectDate,
		ImageURL:     work.ImageURL,
	}

	return responseData, nil
}

func (m *WorkModel) GetWorkById(id int) (*Work, error) {
	var work Work
	if err := m.DB.First(&work, id).Error; err != nil {
		return nil, err
	}
	return &work, nil
}

func (m *WorkModel) GetAllWorks() ([]Work, error) {
	works := []Work{}
	if err := m.DB.Order("created_at DESC").Find(&works).Error; err != nil {
		return nil, err
	}
	return works, nil
}

func (m *WorkModel) UpdateWork(id int, patch map[string]any) (*Work, error) {
	var existingWork Work
	if err := m.DB.First(&existingWork, id).Error; err != nil {
		return nil, err
	}

	// set allowed fields to update
	allowedFields := map[string]bool{
		"title":        true,
		"description":  true,
		"gallery_id":   true,
		"pengurus_id":  true,
		"project_date": true,
	}

	// compare the data and filter the empty value
	filteredUpdates := make(map[string]any)
	for field, value := range patch {
		if allowedFields[field] {
			filteredUpdates[field] = value
		}
	}

	// check if there is any update
	if len(filteredUpdates) == 0 {
		return &existingWork, nil
	}

	// update
	if err := m.DB.Model(&existingWork).Updates(filteredUpdates).Error; err != nil {
		return nil, err
	}

	// reload data -> make sure the data is updated
	if err := m.DB.First(&existingWork, id).Error; err != nil {
		return nil, err
	}

	return &existingWork, nil

}

func (m *WorkModel) DeleteWork(id int) error {
	var work Work
	if err := m.DB.First(&work, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	if err := m.DB.Delete(&work).Error; err != nil {
		return err
	}

	return nil
}
