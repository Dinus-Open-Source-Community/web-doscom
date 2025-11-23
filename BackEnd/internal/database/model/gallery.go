package model

import (
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

type GalleryModel struct {
	DB *gorm.DB
}

type Gallery struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	GalleryName string    `form:"gallery_name"`
	GalleryType string    `form:"gallery_type"`
	Description string    `form:"description"`
	EventDate   string    `form:"event_date"`
	FileSize    int64     `form:"file_size"`
	MimeType    string    `form:"mime_type"`
	AssetUrl    string    `form:"asset_url"`
	Kategori    string    `form:"kategori"`
	CreatedAt   time.Time `form:"created_at"`
	UpdatedAt   time.Time `form:"updated_at"`
}

type GalleryInsert struct {
	FileHeader  *multipart.FileHeader
	AssetUrl    string
	GalleryName string
	Kategori    string
	FileSize    int64
	MimeType    string
}

type CreateGallery struct {
	GalleryType string `form:"gallery_type" binding:"required"`
	Description string `form:"description" binding:"required"`
	EventDate   string `form:"event_date" binding:"required"`
}

type GalleryResponse struct {
	ID          int    `json:"id"`
	GalleryName string `json:"gallery_name"`
	GalleryType string `json:"gallery_type"`
	Description string `json:"description"`
	EventDate   string `json:"event_date"`
	FileSize    int64  `json:"file_size"`
	MimeType    string `json:"mime_type"`
	AssetUrl    string `json:"asset_url"`
}

func (Gallery) TableName() string {
	return "gallery"
}

// insert gallery
func (g *GalleryModel) InsertGallery(gallery *Gallery) (*Gallery, error) {
	if err := g.DB.Create(gallery).Error; err != nil {
		return nil, err
	}

	return gallery, nil
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
