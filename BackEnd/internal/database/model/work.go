package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"web_doscom/internal/constants"

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
	Status       string    `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type CreateRequestWork struct {
	ExistingID   []*int    `form:"existingID_image"`
	PengurusID   int       `form:"pengurus_id" binding:"required"`
	Title        string    `form:"title" binding:"required"`
	Tagline      string    `form:"tagline"`
	Description  string    `form:"description" binding:"required"`
	Slug         string    `form:"slug" binding:"required"`
	ProjectType  string    `form:"project_type" binding:"required"`
	Technologies []string  `form:"technologies[]" binding:"required"`
	ProjectDate  time.Time `form:"project_date" binding:"required"`
	Status       string    `form:"status" binding:"required"`
}

type WorkResponse struct {
	ID           int             `json:"id"`
	Title        string          `json:"title"`
	Tagline      string          `json:"tagline"`
	Description  string          `json:"description"`
	Slug         string          `json:"slug"`
	ProjectType  string          `json:"project_type"`
	Technologies []string        `json:"technologies"`
	ProjectDate  time.Time       `json:"project_date"`
	ImageURL     string          `json:"image_url"`
	Gallery      json.RawMessage `json:"gallery"`
}

type WorkPatch struct {
	PengurusID   int       `json:"pengurus_id"`
	Title        string    `json:"title"`
	Tagline      string    `json:"tagline"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ProjectType  string    `json:"project_type"`
	ProjectDate  time.Time `json:"project_date"`
	Status       string    `json:"status"`
	Technologies []string  `json:"technologies"`
	ExistingID   []*int    `json:"existingID_image"`
}

type WorkPayloadUpdate struct {
	PengurusID   int       `json:"pengurus_id"`
	Title        string    `json:"title"`
	Tagline      string    `json:"tagline"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ProjectType  string    `json:"project_type"`
	ProjectDate  time.Time `json:"project_date"`
	Status       string    `json:"status"`
	Technologies []string  `json:"technologies"`
	UpdatedAt    time.Time `json:"updated_at"`
	ImageURL     string    `json:"image_url"`
}

type WorkUpdateResponse struct {
	ID           int                    `json:"id"`
	PengurusID   int                    `json:"pengurus_id"`
	Title        string                 `json:"title"`
	Tagline      string                 `json:"tagline"`
	Description  string                 `json:"description"`
	Slug         string                 `json:"slug"`
	ProjectType  string                 `json:"project_type"`
	Technologies []string               `json:"technologies"`
	ProjectDate  time.Time              `json:"project_date"`
	ImageURL     string                 `json:"image_url"`
	Status       string                 `json:"status"`
	UpdatedAt    time.Time              `json:"updated_at"`
	CreatedAt    time.Time              `json:"created_at"`
	WorkGallery  []*WorkGalleryResponse `json:"work_gallery"`
}

func (Work) TableName() string {
	return "work"
}

func (m *WorkModel) WithTx(tx *gorm.DB) *WorkModel {
	return &WorkModel{DB: tx}
}

func (m *WorkModel) InsertWork(ctx context.Context, work *Work) (*WorkResponse, error) {
	work.CreatedAt = time.Now()
	work.UpdatedAt = time.Now()
	result := m.DB.WithContext(ctx).Create(work)

	if result.Error != nil {
		return nil, fmt.Errorf("failed while insert data %w", result.Error)
	}

	responseData := &WorkResponse{
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

func (m *WorkModel) GetWorkById(ctx context.Context, id int) (*WorkResponse, error) {
	var work WorkResponse

	query := `
		SELECT
			w.id,
			w.title,
			w.tagline,
			w.description,
			w.slug,
			w.project_type,
			w.technologies,
			w.project_date,
			w.image_url,
			COALESCE(gallery.images, '[]'::json) AS gallery
		FROM work w
		LEFT JOIN LATERAL (
			SELECT json_agg(
				json_build_object(
					'id', g.id,
					'file_url', g.file_url
				)
			) AS images FROM work_gallery wg JOIN gallery g ON g.id = wg.id_gallery
			WHERE wg.id_work = w.id
		)gallery ON true
		WHERE w.id = $1
		  AND w.status = 'published'
	`

	result := m.DB.WithContext(ctx).Raw(query, id).Scan(&work)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get work by id %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("work not found")
	}

	return &work, nil
}

func (m *WorkModel) GetAllWorks(ctx context.Context, offset, limit int) ([]WorkResponse, int64, error) {

	baseQuery := m.DB.WithContext(ctx).Model(&Work{}).Where("status = ?", constants.StatusPublished)

	var totalData int64
	if err := baseQuery.Session(&gorm.Session{}).Count(&totalData).Error; err != nil {
		return nil, 0, fmt.Errorf("error while count the data %w", err)
	}

	var worksDataResponse []WorkResponse
	if err := baseQuery.
		Select("id, title, tagline, description, slug, project_type, technologies, Project_date, image_url").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(&worksDataResponse).
		Error; err != nil {
		return nil, 0, fmt.Errorf("failed while get the data %w", err)
	}

	return worksDataResponse, totalData, nil
}

func (m *WorkModel) GetAllWorksByProjectType(
	ctx context.Context,
	offset, limit int,
	projectType string,
) ([]WorkResponse, int64, error) {
	if strings.TrimSpace(projectType) == "" {
		return nil, 0, fmt.Errorf("project type is required not empty, if empty what should i get")
	}

	baseQuery := m.DB.WithContext(ctx).Model(&Work{}).Where("project_type = ? AND status = ?", projectType, constants.StatusPublished)

	var totalData int64
	if err := baseQuery.Session(&gorm.Session{}).Count(&totalData).Error; err != nil {
		return nil, 0, fmt.Errorf("error while count the data %w", err)
	}

	var worksDataResponse []WorkResponse

	if err := baseQuery.
		Select("id, title, tagline, description, slug, project_type, technologies, Project_date, image_url").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(&worksDataResponse).
		Error; err != nil {
		return nil, 0, fmt.Errorf("failed while get the data %w", err)
	}

	return worksDataResponse, totalData, nil
}

func (m *WorkModel) UpdateWork(id int, patch map[string]any) (*Work, error) {
	var existingWork Work
	if err := m.DB.First(&existingWork, id).Error; err != nil {
		return nil, err
	}

	// set allowed fields to update
	allowedFields := map[string]bool{
		"pengurus_id":  true,
		"title":        true,
		"tagline":      true,
		"description":  true,
		"slug":         true,
		"project_type": true,
		"technologies": true,
		"project_date": true,
		"Status":       true,
		"updated_at":   true,
		"image_url":    true,
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
