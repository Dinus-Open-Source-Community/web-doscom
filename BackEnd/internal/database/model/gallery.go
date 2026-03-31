package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type GalleryModel struct {
	DB *gorm.DB
}

type Gallery struct {
	ID           int       `form:"primaryKey" json:"id"`
	IDUsers      int       `form:"column:id_pengurus" json:"id_users"`
	FileUploadID int       `form:"column:file_upload_id" json:"file_upload_id"`
	GalleryName  string    `form:"gallery_name"`
	GalleryType  string    `form:"gallery_type"`
	Description  string    `form:"description"`
	EventDate    time.Time `form:"event_date"`
	FileURL      string    `form:"file_url"`
	CreatedAt    time.Time `form:"created_at"`
	UpdatedAt    time.Time `form:"updated_at"`
}

type GalleryInsert struct {
	IDUsers      int       `form:"id_users"`
	FileUploadID int       `form:"file_upload_id"`
	GalleryName  string    `form:"gallery_name"`
	GalleryType  string    `form:"gallery_type"`
	Description  string    `form:"description"`
	EventDate    time.Time `form:"event_date" time_format:"2006-01-02"`
}

type CreateGallery struct {
	GalleryName string    `form:"gallery_name" binding:"required"`
	GalleryType string    `form:"gallery_type" binding:"required"`
	Description string    `form:"description" binding:"required"`
	EventDate   time.Time `form:"event_date" binding:"required"`
}

type GalleryResponse struct {
<<<<<<< HEAD
	ID           int       `json:"id"`
	IDUsers      int       `json:"id_users"`
	FileUploadID int       `json:"file_upload_id"`
	GalleryName  string    `json:"gallery_name"`
	GalleryType  string    `json:"gallery_type"`
	Description  string    `json:"description"`
	EventDate    time.Time `json:"event_date"`
	FileURL      string    `json:"file_url"`
=======
	ID          int    `json:"id"`
	GalleryName string `json:"gallery_name"`
	GalleryType string `json:"gallery_type"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
	FileSize    int64  `json:"file_size"`
	MimeType    string `json:"mime_type"`
	AssetUrl    string `json:"asset_url"`
>>>>>>> master
}

func (Gallery) TableName() string {
	return "gallery"
}

// insert gallery
func (g *GalleryModel) InsertGallery(gallery *Gallery) (*GalleryResponse, error) {
	if err := g.DB.Create(gallery).Error; err != nil {
		return nil, err
	}

	return &GalleryResponse{
		ID:           gallery.ID,
		IDUsers:      gallery.IDUsers,
		FileUploadID: gallery.FileUploadID,
		GalleryName:  gallery.GalleryName,
		GalleryType:  gallery.GalleryType,
		Description:  gallery.Description,
		EventDate:    gallery.EventDate,
	}, nil
}

func (g *GalleryModel) InsertGalleryMultiple(gallery []*Gallery) ([]*GalleryResponse, error) {
	if err := g.DB.Create(&gallery).Error; err != nil {
		return nil, err
	}

	response := make([]*GalleryResponse, len(gallery))
	for i, data := range gallery {
		response[i] = &GalleryResponse{
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

// get gallery by id multiple
// func (g *GalleryModel) GetGalleryByID(id []int) ([]*GalleryResponse, error) {
// 	var gallery []*Gallery
// 	if err := g.DB.Where("id = ?", &id).Find(&gallery).Error; err != nil {
// 		return nil, err
// 	}
//
// 	response := make([]*GalleryResponse, len(gallery))
// 	for i, data := range gallery {
// 		response[i] = &GalleryResponse{
// 			ID:           data.ID,
// 			IDUsers:      data.IDUsers,
// 			FileUploadID: data.FileUploadID,
// 			GalleryName:  data.GalleryName,
// 			GalleryType:  data.GalleryType,
// 			Description:  data.Description,
// 			EventDate:    data.EventDate,
// 			FileURL:      data.FileURL,
// 		}
// 	}
// 	return response, nil
// }

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

func (g *GalleryModel) CheckExistingGallery(id []int) (bool, error) {
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
) ([]GalleryResponse, int64, error) {

	var galleryResponse []GalleryResponse
	var total int64

	query := g.DB.WithContext(ctx).Model(&Gallery{})
	if startYear != nil && endYear != nil {
		query = query.Where("created_at >= ? AND created_at < ?", startYear, endYear)
	}

	if err := query.Model(&Gallery{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Model(&Gallery{}).
		Select("id, description, file_url, event_date").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Scan(&galleryResponse).
		Error; err != nil {
		return nil, 0, err
	}

	return galleryResponse, total, nil
}
