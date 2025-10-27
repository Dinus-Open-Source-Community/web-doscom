package model

import (
	"time"

	"gorm.io/gorm"
)

type WorkModel struct {
	DB *gorm.DB
}

type Work struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// PengurusID  int       `json:"pengurus_id"`
	ProjectDate string `json:"project_date"`
	// GalleryID   int       `json:"gallery_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkContent struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required" regex:"^[a-zA-Z0-9 ]+$"`
	Description string `json:"description" binding:"required" regex:"^[a-zA-Z0-9 ]+$"`
	// PengurusID  int       `json:"pengurus_id" binding:"required"`
	ProjectDate string `json:"project_date" binding:"required" regex:"^[0-9-]+$"`
	// GalleryID   int       `json:"gallery_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *WorkModel) InsertWork(work *Work) error {
	work.CreatedAt = time.Now()
	work.UpdatedAt = time.Now()
	return m.DB.Create(work).Error
}
