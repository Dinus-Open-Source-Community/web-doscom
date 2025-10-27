package model

import (
	"time"

	"gorm.io/gorm"
)

type BlogModel struct {
	DB *gorm.DB
}

type Blog struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	IsPublished bool      `json:"is_published"`
	// PengurusID   int       `json:"pengurus_id"`
	// GalleryID    int       `json:"gallery_id"`
	// ActivitiesID int       `json:"activities_id"`
	// WorkID       int       `json:"work_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BlogContent struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required" regex:"^[a-zA-Z0-9 ]+$"`
	Slug        string    `json:"slug" binding:"required" regex:"^[a-z0-9-]+$"`
	Content     string    `json:"content" binding:"required"`
	PublishedAt time.Time `json:"published_at" binding:"required"`
	IsPublished bool      `json:"is_published" binding:"required : boolean"`
	// PengurusID   int       `json:"pengurus_id" binding:"required"`
	// GalleryID    int       `json:"gallery_id" binding:"required"`
	// WorkID       int       `json:"work_id" binding:"required" default:"false"`
	// ActivitiesID int       `json:"activities_id" binding:"required" default:"false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Blog) TableName() string {
	return "blog"
}

func (m *BlogModel) InsertBlog(blog *Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	return m.DB.Create(blog).Error
}
