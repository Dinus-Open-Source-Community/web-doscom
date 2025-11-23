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

type GalleryUpdate struct {
	GalleryName *string `json:"gallery_name" binding:"omitempty"`
	GalleryType *string `json:"gallery_type" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
	EventDate   *string `json:"event_date" binding:"omitempty"`
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

<<<<<<< HEAD
// get gallery by id
func (g *GalleryModel) GetGalleryByID(id int) (*Gallery, error) {
	var gallery Gallery
	if err := g.DB.First(&gallery, id).Error; err != nil {
		return nil, err
	}

	return &gallery, nil
}

// get all gallery
func (g *GalleryModel) GetAllGallery(page, limit, offset int) ([]*Gallery, int64, error) {
	var GalleryData []*Gallery
	// take data per page
	if err := g.DB.
		Offset(offset).
		Limit(limit).
		Find(&GalleryData).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	// hitung total data
	g.DB.Model(&Gallery{}).
		Count(&total)

	return GalleryData, total, nil
}

// update gallery
func (g *GalleryModel) UpdateGallery(id int, patch GalleryUpdate) (*Gallery, error) {
	var gallery Gallery
	// find gallery by id
	if err := g.DB.First(&gallery, id).Error; err != nil {
		return nil, err
	}
	if err := g.DB.Model(&gallery).Updates(patch).Error; err != nil {
		return nil, err
	}
	return &gallery, nil
}

=======
>>>>>>> fd36652 (add some function, fix some function, malaasss)
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
