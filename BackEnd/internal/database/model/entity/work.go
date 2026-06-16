package entity

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type WorkModel struct {
	DB *gorm.DB
}

type Work struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	PengurusID   int            `gorm:"column:pengurus_id" json:"pengurus_id"`
	Title        string         `gorm:"column:title" json:"title"`
	Tagline      string         `gorm:"column:tagline" json:"tagline"`
	Description  string         `gorm:"column:description" json:"description"`
	Slug         string         `gorm:"column:slug" json:"slug"`
	ProjectType  string         `gorm:"column:project_type" json:"project_type"`
	Technologies pq.StringArray `gorm:"type:text[]" json:"technologies"`
	ProjectDate  time.Time      `gorm:"column:project_date" json:"project_date"`
	ImageURL     string         `gorm:"column:image_url" json:"image_url"`
	Status       string         `gorm:"column:status" json:"status"`
	Division     string         `gorm:"column:division" json:"division"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

func (Work) TableName() string {
	return "work"
}

func (m *WorkModel) WithTx(tx *gorm.DB) *WorkModel {
	return &WorkModel{DB: tx}
}

func (m *WorkModel) InsertWork(ctx context.Context, work *Work) (*dto.WorkResponseClient, error) {
	work.CreatedAt = time.Now()
	work.UpdatedAt = time.Now()
	result := m.DB.WithContext(ctx).Create(work)

	if result.Error != nil {
		return nil, fmt.Errorf("failed while insert data %w", result.Error)
	}

	responseData := &dto.WorkResponseClient{
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

func (m *WorkModel) GetWorkById(ctx context.Context, id int) (*dto.WorkResponseInternal, error) {
	return m.getWorkById(ctx, id, true)
}

func (m *WorkModel) GetWorkByIDForAdmin(ctx context.Context, id int) (*dto.WorkResponseInternal, error) {
	return m.getWorkById(ctx, id, false)
}

func (m *WorkModel) getWorkById(ctx context.Context, id int, onlyPublished bool) (*dto.WorkResponseInternal, error) {
	var work dto.WorkResponseInternal

	query := `
		SELECT
			w.id,
			w.pengurus_id,
			w.title,
			w.tagline,
			w.description,
			w.slug,
			w.status,
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
		WHERE w.id = @id
	`
	args := map[string]any{
		"id": id,
	}
	if onlyPublished {
		query = query + ` AND w.status = @status`
		args["status"] = constants.StatusPublished
	}

	result := m.DB.WithContext(ctx).Raw(query, args).Scan(&work)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get work by id %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("work not found")
	}

	return &work, nil
}

func (m *WorkModel) GetWorkWithOutGallery(ctx context.Context, workID int) (Work, error) {

	var work Work
	result := m.DB.WithContext(ctx).
		Model(&Work{}).
		Where("id = ?", workID).
		First(&work)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Work{}, fmt.Errorf("work not found")
		}
		return Work{}, fmt.Errorf("failed to get work by id %w", result.Error)
	}

	return work, nil
}

func (m *WorkModel) GetAllWorks(ctx context.Context, offset, limit int) ([]dto.WorkResponseClient, int64, error) {

	baseQuery := m.DB.WithContext(ctx).Model(&Work{}).Where("status = ?", constants.StatusPublished)

	var totalData int64
	if err := baseQuery.Session(&gorm.Session{}).Count(&totalData).Error; err != nil {
		return nil, 0, fmt.Errorf("error while count the data %w", err)
	}

	var worksDataResponse []dto.WorkResponseClient
	if err := baseQuery.
		Select("id, title, tagline, description, slug, project_type, technologies, project_date, image_url").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(&worksDataResponse).
		Error; err != nil {
		return nil, 0, fmt.Errorf("failed while get the data %w", err)
	}

	return worksDataResponse, totalData, nil
}

func (m *WorkModel) GetAllWorksAdmin(
	ctx context.Context,
	division string,
	status []string,
	filterByDivision bool,
	offset, limit int,
) ([]dto.WorkResponseInternal, int64, error) {

	baseQuery := m.DB.WithContext(ctx).Model(&Work{}).Where("status IN ?", status)
	log.Printf("[model] status: %v", status)
	if filterByDivision {
		baseQuery = baseQuery.Where("division = ?", division)
	}

	var totalData int64
	if err := baseQuery.Session(&gorm.Session{}).Count(&totalData).Error; err != nil {
		return nil, 0, fmt.Errorf("error while count the data %w", err)
	}

	var worksDataResponse []dto.WorkResponseInternal
	if err := baseQuery.
		Select("id, title, tagline, description, slug, status, project_type, technologies, Project_date, image_url").
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
) ([]dto.WorkResponseClient, int64, error) {
	if strings.TrimSpace(projectType) == "" {
		return nil, 0, fmt.Errorf("project type is required not empty, if empty what should i get")
	}

	baseQuery := m.DB.WithContext(ctx).Model(&Work{}).Where("project_type = ? AND status = ?", projectType, constants.StatusPublished)

	var totalData int64
	if err := baseQuery.Session(&gorm.Session{}).Count(&totalData).Error; err != nil {
		return nil, 0, fmt.Errorf("error while count the data %w", err)
	}

	var worksDataResponse []dto.WorkResponseClient

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
		"status":       true,
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

func (m *WorkModel) DeleteWork(ctx context.Context, id int) error {

	result := m.DB.WithContext(ctx).Delete(&Work{}, id)

	if result.Error != nil {
		return fmt.Errorf("failed to delete work %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("work not found")
	}
	return nil
}

func (m *WorkModel) GetWorkTypes(ctx context.Context) ([]string, error) {
	var projectTypes []string
	result := m.DB.WithContext(ctx).Model(&Work{}).Distinct("project_type").Where("project_type <> ''").Pluck("project_type", &projectTypes)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to get work types %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no work types found")
	}

	return projectTypes, nil
}
