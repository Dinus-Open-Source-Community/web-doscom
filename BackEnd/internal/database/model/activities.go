package model

import (
	"time"

	"gorm.io/gorm"
)

type ActivityModel struct {
	DB *gorm.DB
}

type Activities struct {
	ID              uint      `json:"id"`
	ActivitiesTitle string    `json:"activities_title"`
	ActivitiesDesc  string    `json:"activities_desc"`
	ActivitiesDate  time.Time `json:"activities_date"`
	// PengurusID      int       `json:"pengurus_id"`
	// GalleryID       int       `json:"gallery_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ActivityContent struct {
	ID              uint      `json:"id"`
	ActivitiesTitle string    `json:"activities_title" binding:"required" regex:"^[a-zA-Z0-9 ]+$"`
	ActivitiesDesc  string    `json:"activities_desc" binding:"required" regex:"^[a-zA-Z0-9 ]+$"`
	ActivitiesDate  time.Time `json:"activities_date" binding:"required"`
	// PengurusID      int       `json:"pengurus_id" binding:"required"`
	// GalleryID       int       `json:"gallery_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ActivityModel) InsertActivity(activity *Activities) error {
	activity.CreatedAt = time.Now()
	activity.UpdatedAt = time.Now()
	return m.DB.Create(activity).Error
}
