package model

import (
	"context"
	"gorm.io/gorm"
	"time"
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

type BlogGalleryInsert struct {
	BlogID    int `gorm:"column:id_blog" json:"blog_id"`
	GalleryID int `gorm:"column:id_gallery" json:"gallery_id"`
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

func (b *BlogGalleryModel) InsertBlogGalleryMultiple(BlogGallery []*BlogGallery) ([]*BlogGallery, error) {
	if err := b.DB.Create(&BlogGallery).Error; err != nil {
		return nil, err
	}

	return BlogGallery, nil
}

func (b *BlogGalleryModel) UpdateBlogGallery(galleryID []*int, idBlog int) ([]*BlogGallery, error) {
	tx := b.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&BlogGallery{}).
		Where("id_blog = ?", idBlog).
		Delete(&BlogGallery{}).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	data := make([]*BlogGallery, len(galleryID))
	for i, galleryIDs := range galleryID {
		data[i] = &BlogGallery{
			BlogID:    idBlog,
			GalleryID: *galleryIDs,
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
