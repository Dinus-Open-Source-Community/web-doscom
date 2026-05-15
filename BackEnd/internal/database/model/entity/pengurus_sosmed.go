package entity

import (
	"context"
	"fmt"
	"time"
	"web_doscom/internal/database/model/dto"

	"gorm.io/gorm"
)

type PengurusSosmedModel struct {
	DB *gorm.DB
}

func (PengurusSosmed) TableName() string {
	return "pengurus_sosmed"
}

type PengurusSosmed struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	PengurusID int       `gorm:"column:pengurus_id" json:"pengurus_id"`
	Platform   string    `json:"platform"`
	Username   string    `json:"username"`
	Url        string    `json:"url"`
	IsPrimary  bool      `json:"is_primary"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *PengurusSosmedModel) InsertPengurusSosmed(
	ctx context.Context,
	data []dto.CreatePengurusSosmedPayload,
) ([]dto.PengurusSosmedResponse, error) {

	if len(data) == 0 {
		return nil, fmt.Errorf("sosmed data is required to insert")
	}

	entity := make([]PengurusSosmed, len(data))
	for i, d := range data {
		entity[i] = PengurusSosmed{
			PengurusID: d.PengurusID,
			Platform:   d.Platform,
			Username:   d.Username,
			Url:        d.Url,
			IsPrimary:  d.IsPrimary,
		}
	}

	if err := m.DB.WithContext(ctx).Create(entity).Error; err != nil {
		return nil, fmt.Errorf("failed to insert data pengurus sosmed %w", err)
	}

	result := make([]dto.PengurusSosmedResponse, len(entity))
	for i, data := range entity {
		result[i] = dto.PengurusSosmedResponse{
			Platform:  data.Platform,
			Username:  data.Username,
			Url:       data.Url,
			IsPrimary: data.IsPrimary,
		}
	}

	return result, nil
}

func (m *PengurusSosmedModel) UpdatePengurusSosmed(ctx context.Context, pengurusID int, sosmedPayload []dto.CreatePengurusSosmedPayload) ([]dto.PengurusSosmedResponse, error) {
	if len(sosmedPayload) == 0 {
		return nil, fmt.Errorf("sosmed data is required to update")
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

	if err := tx.Where("pengurus_id = ?", pengurusID).Delete(&PengurusSosmed{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	data := make([]*PengurusSosmed, len(sosmedPayload))
	for i, d := range sosmedPayload {
		data[i] = &PengurusSosmed{
			PengurusID: d.PengurusID,
			Platform:   d.Platform,
			Username:   d.Username,
			Url:        d.Url,
			IsPrimary:  i == 0, // true hanya untuk index 0
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

	responseData := make([]dto.PengurusSosmedResponse, len(data))
	for i, sosmed := range data {
		responseData[i] = dto.PengurusSosmedResponse{
			Platform:  sosmed.Platform,
			Username:  sosmed.Username,
			Url:       sosmed.Url,
			IsPrimary: sosmed.IsPrimary,
		}
	}

	return responseData, nil
}

func (m *PengurusSosmedModel) GetSosmedByPengurusID(ctx context.Context, pengurusID int) ([]dto.PengurusSosmedResponse, error) {
	var sosmed []dto.PengurusSosmedResponse
	if err := m.DB.WithContext(ctx).
		Model(&PengurusSosmed{}).
		Where("pengurus_id = ?", pengurusID).
		Select("platform, username, url, is_primary").
		Scan(&sosmed).
		Error; err != nil {
		return nil, fmt.Errorf("failed to get data pengurus sosmed: %w", err)
	}

	return sosmed, nil
}

func (m *PengurusSosmedModel) DeleteByPengurusID(ctx context.Context, pengurusID int) error {
	if err := m.DB.WithContext(ctx).
		Where("pengurus_id = ?", pengurusID).
		Delete(&PengurusSosmed{}).
		Error; err != nil {
		return fmt.Errorf("failed to delete data pengurus sosmed: %w", err)
	}

	return nil
}
