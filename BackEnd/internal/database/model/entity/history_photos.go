package entity

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type HistoryPhotosModel struct {
	DB *gorm.DB
}

type HistoryPhotos struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	IDHistory int       `gorm:"column:id_history" json:"id_history"`
	ImagerURL string    `gorm:"column:image_url" json:"image_url"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (HistoryPhotos) TableName() string {
	return "history_photos"
}

func (m *HistoryPhotosModel) InsertHistoryPhotos(ctx context.Context, historyPhotos *HistoryPhotos) error {
	historyPhotos.CreatedAt = time.Now()
	historyPhotos.UpdatedAt = time.Now()

	if err := m.DB.WithContext(ctx).Create(historyPhotos).Error; err != nil {
		return err
	}
	return nil
}
