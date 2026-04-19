package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type WorkGalleryModel struct {
	DB *gorm.DB
}

type WorkGallery struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	IDWork    int       `gorm:"column:id_work" json:"id_work"`
	IDGallery int       `gorm:"column:id_gallery" json:"id_gallery"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type WorkGalleryInsert struct {
	IDWork    int `gorm:"column:id_work" json:"id_work" binding:"required"`
	IDGallery int `gorm:"column:id_gallery" json:"id_gallery" binding:"required"`
}

type WorkGalleryResponse struct {
	ID        int `json:"id"`
	IDWork    int `json:"id_work"`
	IDGallery int `json:"id_gallery"`
}

func (WorkGallery) TableName() string {
	return "work_gallery"
}

func (m *WorkGalleryModel) WithTx(tx *gorm.DB) *WorkGalleryModel {
	return &WorkGalleryModel{DB: tx}
}

func (m *WorkGalleryModel) InsertWorkGallery(
	ctx context.Context,
	workGallery *WorkGallery,
) (WorkGalleryResponse, error) {
	workGallery.CreatedAt = time.Now()
	workGallery.UpdatedAt = time.Now()

	if err := m.DB.WithContext(ctx).Create(workGallery).Error; err != nil {
		return WorkGalleryResponse{}, fmt.Errorf("failed to insert data %w", err)
	}

	return WorkGalleryResponse{
		ID:        workGallery.ID,
		IDWork:    workGallery.IDWork,
		IDGallery: workGallery.IDGallery,
	}, nil

}

func (m *WorkGalleryModel) InsertWorkGalleryMultiple(ctx context.Context, workGallery []*WorkGallery) error {

	if len(workGallery) == 0 {
		return fmt.Errorf("work gallery is requied not empty")
	}
	now := time.Now()
	for i := range workGallery {
		workGallery[i].CreatedAt = now
		workGallery[i].UpdatedAt = now
	}

	if err := m.DB.WithContext(ctx).Create(workGallery).Error; err != nil {
		return fmt.Errorf("failed while insert data %w", err)
	}

	return nil
}
