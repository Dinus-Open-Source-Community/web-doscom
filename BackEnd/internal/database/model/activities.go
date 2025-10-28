package model

import (
	"time"

	"gorm.io/gorm"
)

type ActivityModel struct {
	DB *gorm.DB
}

type Activities struct {
	ID              int       `gorm:"primaryKey" json:"id"`
	GalleryID       int       `json:"id_asset"`
	ActivitiesTitle string    `json:"activities_title"`
	ActivitiesDesc  string    `json:"activities_desc"`
	ActivitiesDate  time.Time `json:"activities_date"`
	PengurusID      int       `json:"pengurus_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type RegisterActivities struct {
	GalleryID       int    `json:"id_asset" binding:"required"`
	ActivitiesTitle string `json:"activities_title" binding:"required"`
	ActivitiesDesc  string `json:"activities_desc" binding:"required"`
	ActivitiesDate  string `json:"activities_date" binding:"required"`
	PengurusID      int    `json:"pengurus_id" binding:"required"`
}

type ActivitiesPatch struct {
	GalleryID       *int    `json:"id_asset" binding:"omitempty"`
	ActivitiesTitle *string `json:"activities_title" binding:"omitempty"`
	ActivitiesDesc  *string `json:"activities_desc" binding:"omitempty"`
	ActivitiesDate  *string `json:"activities_date" binding:"omitempty"`
	PengurusID      *int    `json:"pengurus_id" binding:"omitempty"`
}

func (Activities) TableName() string {
	return "activities"
}

// CRUD implementation

func (m *ActivityModel) InsertActivity(activity *Activities) error {
	activity.CreatedAt = time.Now()
	activity.UpdatedAt = time.Now()
	return m.DB.Create(activity).Error
}

func (m *ActivityModel) GetActivitiesById(id int) (*Activities, error) {
	var activity Activities
	if err := m.DB.First(&activity, id).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

func (m *ActivityModel) GetAllActivities() ([]Activities, error) {
	var activities []Activities
	if err := m.DB.Order("created_at desc").Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (m *ActivityModel) UpdateActivities(id int, input *Activities) (*Activities, error) {
	var activity Activities
	if err := m.DB.First(&activity, id).Error; err != nil {
		return nil, err
	}
	// Update fields
	activity.GalleryID = input.GalleryID
	activity.ActivitiesTitle = input.ActivitiesTitle
	activity.ActivitiesDesc = input.ActivitiesDesc
	activity.ActivitiesDate = input.ActivitiesDate
	activity.PengurusID = input.PengurusID
	activity.UpdatedAt = time.Now()
	if err := m.DB.Save(&activity).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

func (m *ActivityModel) DeleteActivities(id int) error {
	return m.DB.Delete(&Activities{}, id).Error
}
