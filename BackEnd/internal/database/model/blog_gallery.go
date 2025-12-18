package model

import (
	"gorm.io/gorm"
)

func (BlogGallery) TableName() string {
	return "blog_gallery"
}

type BlogGalleryModel struct {
	DB *gorm.DB
}

type BlogGallery struct {
	ID        int `json:"id"`
	BlogID    int `json:"blog_id"`
	GalleryID int `json:"gallery_id"`
}

type BlogGalleryInsert struct {
	BlogID    int `json:"blog_id"`
	GalleryID int `json:"gallery_id"`
}

// insert blog gallery
func (b *BlogGalleryModel) InsertBlogGallery(blogGallery *BlogGallery) (*BlogGallery, error) {
	if err := b.DB.Create(blogGallery).Error; err != nil {
		return nil, err
	}

	return blogGallery, nil
}

func (b *BlogGalleryModel) GetBlogGalleryByID(id int) (*BlogGallery, error) {
	var gallery BlogGallery
	if err := b.DB.First(&gallery, id).Error; err != nil {
		return nil, err
	}

	return &gallery, nil
}
