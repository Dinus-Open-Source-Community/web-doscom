package model

import (
	"time"

	"gorm.io/gorm"
)

type WorkModel struct {
	DB *gorm.DB
}

type Work struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ProjectDate string    `json:"project_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (m *WorkModel) InsertWork(work *Work) error {
	work.CreatedAt = time.Now()
	work.UpdatedAt = time.Now()
	return m.DB.Create(work).Error
}
