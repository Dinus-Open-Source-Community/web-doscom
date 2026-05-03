package entity

import (
	"context"
	"fmt"
	"time"
	"web_doscom/internal/database/model/dto"

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

func (WorkGallery) TableName() string {
	return "work_gallery"
}

func (m *WorkGalleryModel) WithTx(tx *gorm.DB) *WorkGalleryModel {
	return &WorkGalleryModel{DB: tx}
}

func (m *WorkGalleryModel) InsertWorkGallery(
	ctx context.Context,
	workGallery *WorkGallery,
) (dto.WorkGalleryResponse, error) {
	workGallery.CreatedAt = time.Now()
	workGallery.UpdatedAt = time.Now()

	if err := m.DB.WithContext(ctx).Create(workGallery).Error; err != nil {
		return dto.WorkGalleryResponse{}, fmt.Errorf("failed to insert data %w", err)
	}

	return dto.WorkGalleryResponse{
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

func (m *WorkGalleryModel) UpdateWorkGallery(galleryID []int, idWork int) ([]*dto.WorkGalleryResponse, error) {
	if len(galleryID) == 0 {
		return nil, fmt.Errorf("galleryID is empty cannot update database")
	}

	tx := m.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Where("id_work = ?", idWork).Delete(&WorkGallery{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	data := make([]*WorkGallery, len(galleryID))
	for i, id := range galleryID {
		data[i] = &WorkGallery{
			IDWork:    idWork,
			IDGallery: id,
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

	responseData := make([]*dto.WorkGalleryResponse, len(data))
	for i, workGallery := range data {
		responseData[i] = &dto.WorkGalleryResponse{
			ID:        workGallery.ID,
			IDWork:    workGallery.IDWork,
			IDGallery: workGallery.IDGallery,
		}
	}
	return responseData, nil
}

func (m *WorkGalleryModel) DeleteWorkGalleryByID(ctx context.Context, workID int) error {
	result := m.DB.WithContext(ctx).Where("id_work = ?", workID).Delete(&WorkGallery{})

	if result.Error != nil {
		return fmt.Errorf("failed to delete work gallery %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("work gallery not found")
	}

	return nil
}
