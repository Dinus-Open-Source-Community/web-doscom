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
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (m *ActivityModel) InsertActivity(activity *Activities) error {
	activity.CreatedAt = time.Now()
	activity.UpdatedAt = time.Now()
	return m.DB.Create(activity).Error
}
