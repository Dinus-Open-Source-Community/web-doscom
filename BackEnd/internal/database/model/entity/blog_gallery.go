package entity

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"

	"web_doscom/internal/database/model/dto"
)

func (BlogGallery) TableName() string {
	return "blog_gallery"
}

type BlogGalleryModel struct {
	DB *gorm.DB
}

type BlogGallery struct {
	ID        int       `json:"id"`
	BlogID    int       `gorm:"column:id_blog" json:"blog_id"`
	GalleryID int       `gorm:"column:id_gallery" json:"gallery_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (b *BlogGalleryModel) WithTx(tx *gorm.DB) *BlogGalleryModel {
	return &BlogGalleryModel{DB: tx}
}

// insert blog gallery
func (b *BlogGalleryModel) InsertBlogGallery(blogGallery *BlogGallery) (*BlogGallery, error) {
	blogGallery.CreatedAt = time.Now()
	blogGallery.UpdatedAt = time.Now()
	if err := b.DB.Create(blogGallery).Error; err != nil {
		return nil, err
	}

	return blogGallery, nil
}

func (b *BlogGalleryModel) InsertBlogGalleryMultiple(ctx context.Context, galleries []*BlogGallery) ([]*dto.BlogGalleryResponse, error) {

	if len(galleries) == 0 {
		return nil, fmt.Errorf("gallery cannot be empty")
	}

	result := b.DB.WithContext(ctx).Create(&galleries)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to insert data: %w", result.Error)
	}

	response := make([]*dto.BlogGalleryResponse, len(galleries))
	for i, data := range galleries {
		response[i] = &dto.BlogGalleryResponse{
			ID:        data.ID,
			BlogID:    data.BlogID,
			GalleryID: data.GalleryID,
		}
	}

	return response, nil

}

func (b *BlogGalleryModel) UpdateBlogGallery(galleryID []int, idBlog int) ([]*BlogGallery, error) {
	if len(galleryID) == 0 {
		return nil, fmt.Errorf("galleryID is empty cannot update database")
	}

	tx := b.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Where("id_blog = ?", idBlog).Delete(&BlogGallery{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	data := make([]*BlogGallery, len(galleryID))
	for i, id := range galleryID {
		data[i] = &BlogGallery{
			BlogID:    idBlog,
			GalleryID: id,
		}
	}

	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return data, nil
}

func (b *BlogGalleryModel) GetBlogGalleryByID(id int) (*BlogGallery, error) {
	var gallery BlogGallery
	if err := b.DB.First(&gallery, id).Error; err != nil {
		return nil, err
	}

	return &gallery, nil
}

func (b *BlogGalleryModel) DeleteBlogGalleryByBlogID(ctx context.Context, tx *gorm.DB, id int) error {

	if err := tx.WithContext(ctx).
		Where("id_blog = ?", id).
		Delete(&BlogGallery{}).
		Error; err != nil {
		return err
	}
	return nil
}
