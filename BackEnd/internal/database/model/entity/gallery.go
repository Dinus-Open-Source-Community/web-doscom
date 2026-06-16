package entity

import (
	"context"
	"fmt"
	"time"
	"web_doscom/internal/database/model/dto"

	"gorm.io/gorm"
)

type GalleryModel struct {
	DB *gorm.DB
}

type Gallery struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	IDUsers      int       `gorm:"column:id_users" json:"id_users"`
	FileUploadID int       `gorm:"column:file_upload_id" json:"file_upload_id"`
	GalleryName  string    `gorm:"column:gallery_name" json:"gallery_name"`
	GalleryType  string    `gorm:"column:gallery_type" json:"gallery_type"`
	Description  string    `gorm:"column:description" json:"description"`
	EventDate    time.Time `gorm:"column:event_date" json:"event_date"`
	FileURL      string    `gorm:"column:file_url" json:"file_url"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Gallery) TableName() string {
	return "gallery"
}

func (g *GalleryModel) WithTx(tx *gorm.DB) *GalleryModel {
	return &GalleryModel{DB: tx}
}

// insert gallery
func (g *GalleryModel) InsertGallery(gallery *Gallery) (*dto.GalleryResponse, error) {
	if err := g.DB.Create(gallery).Error; err != nil {
		return nil, err
	}

	return &dto.GalleryResponse{
		ID:           gallery.ID,
		IDUsers:      gallery.IDUsers,
		FileUploadID: gallery.FileUploadID,
		GalleryName:  gallery.GalleryName,
		GalleryType:  gallery.GalleryType,
		Description:  gallery.Description,
		EventDate:    gallery.EventDate,
		FileURL:      gallery.FileURL,
	}, nil
}

func (g *GalleryModel) InsertGalleryMultiple(gallery []*Gallery) ([]*dto.GalleryResponse, error) {
	if err := g.DB.Create(&gallery).Error; err != nil {
		return nil, err
	}

	response := make([]*dto.GalleryResponse, len(gallery))
	for i, data := range gallery {
		response[i] = &dto.GalleryResponse{
			ID:           data.ID,
			IDUsers:      data.IDUsers,
			FileUploadID: data.FileUploadID,
			GalleryName:  data.GalleryName,
			GalleryType:  data.GalleryType,
			Description:  data.Description,
			EventDate:    data.EventDate,
			FileURL:      data.FileURL,
		}
	}
	return response, nil
}

func (g *GalleryModel) GetGalleryByID(ctx context.Context, id int) (*dto.GalleryResponse, error) {
	var gallery Gallery
	if err := g.DB.WithContext(ctx).First(&gallery, id).Error; err != nil {
		return nil, err
	}

	return &dto.GalleryResponse{
		ID:           gallery.ID,
		IDUsers:      gallery.IDUsers,
		FileUploadID: gallery.FileUploadID,
		GalleryName:  gallery.GalleryName,
		GalleryType:  gallery.GalleryType,
		Description:  gallery.Description,
		EventDate:    gallery.EventDate,
		FileURL:      gallery.FileURL,
	}, nil
}

// get gallery by id multiple
func (g *GalleryModel) GetGalleryByIDMultiple(ctx context.Context, id []int) ([]*dto.GalleryResponse, error) {
	var gallery []*Gallery
	if err := g.DB.WithContext(ctx).Where("id IN ?", id).Find(&gallery).Error; err != nil {
		return nil, err
	}

	response := make([]*dto.GalleryResponse, len(gallery))
	for i, data := range gallery {
		response[i] = &dto.GalleryResponse{
			ID:           data.ID,
			IDUsers:      data.IDUsers,
			FileUploadID: data.FileUploadID,
			GalleryName:  data.GalleryName,
			GalleryType:  data.GalleryType,
			Description:  data.Description,
			EventDate:    data.EventDate,
			FileURL:      data.FileURL,
		}
	}
	return response, nil
}

// get gallery by type
func (g *GalleryModel) GetGalleryByType(galleryType string, page, limit, offset int) ([]*Gallery, int64, error) {

	var GalleryData []*Gallery
	// take data per page
	if err := g.DB.Where("gallery_type = ?", galleryType).
		Offset(offset).
		Limit(limit).
		Find(&GalleryData).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	// hitung total data
	g.DB.Model(&Gallery{}).
		Where("gallery_type = ?", galleryType).
		Count(&total)

	return GalleryData, total, nil
}

// delete gallery by id
func (g *GalleryModel) DeleteGallery(id int) error {
	var gallery Gallery
	if err := g.DB.First(&gallery, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	if err := g.DB.Delete(&gallery).Error; err != nil {
		return err
	}

	return nil
}

func (g *GalleryModel) DeleteGalleryByIdMultiple(ctx context.Context, id []int) error {
	if len(id) == 0 {
		return fmt.Errorf("ids is required, not empty, if empty what should i delete")
	}

	result := g.DB.WithContext(ctx).Where("id IN ?", id).Delete(&Gallery{})

	if result.Error != nil {
		return fmt.Errorf("failed to delete data: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("failed, no data match to delete")
	}

	return nil
}

func (g *GalleryModel) CheckExistingGallery(id []*int) (bool, error) {
	if len(id) == 0 {
		return false, nil
	}
	var count int64
	result := g.DB.Model(&Gallery{}).Where("id in ?", id).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return int(count) == len(id), nil
}

// get all gallery -> if theres no filter year applied
// get gallery by event_date -> if theres filter year applied
func (g *GalleryModel) GetAllGalleryAndByYear(
	ctx context.Context,
	startYear, endYear *time.Time,
	limit, offset int,
) ([]dto.GalleryResponse, int64, error) {

	var total int64

	query := g.DB.WithContext(ctx).Model(&Gallery{})
	if startYear != nil && endYear != nil {
		query = query.Where("created_at >= ? AND created_at < ?", startYear, endYear)
	}

	if err := query.Model(&Gallery{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	galleryResponse := []dto.GalleryResponse{}
	if err := query.Model(&Gallery{}).
		Select("id, id_users, file_upload_id, description, file_url, event_date").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Scan(&galleryResponse).
		Error; err != nil {
		return nil, 0, err
	}

	return galleryResponse, total, nil
}
