package entity

import (
	"context"
	"fmt"
	"time"
	"web_doscom/internal/database/model/dto"

	"gorm.io/gorm"
)

type HistoryTimelineModel struct {
	DB *gorm.DB
}

type HistoryTimeline struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	IDAuthor     int       `gorm:"column:id_author" json:"id_author"`
	Title        string    `gorm:"column:title" json:"title"`
	Year         string    `gorm:"column:year" json:"year"`
	Description  string    `gorm:"column:description" json:"description"`
	DisplayOrder int       `gorm:"column:display_order" json:"display_order"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (HistoryTimeline) TableName() string {
	return "history_timeline"
}

func (m *HistoryTimelineModel) InsertHistoryTimeline(ctx context.Context, historyData *HistoryTimeline) error {
	historyData.CreatedAt = time.Now()
	historyData.UpdatedAt = time.Now()

	if err := m.DB.WithContext(ctx).Create(historyData).Error; err != nil {
		return err
	}
	return nil
}

func (m *HistoryTimelineModel) GetHistoryTimelineByid(ctx context.Context, id int) (*dto.HistoryTimelineResponse, error) {
	var history dto.HistoryTimelineResponse

	query := `
		SELECT
			ht.id,
			ht.id_author,
			ht.title,
			ht.description,
			ht.year,
			ht.display_order,
			COALESCE(
				(
					SELECT json_agg(
						json_build_object(
							'id', hp.id,
							'id_history', hp.id_history,
							'image_url', hp.image_url
						)
						ORDER BY hp.id DESC
					)
					FROM history_photos hp
					WHERE hp.id_history = ht.id
				),
				'[]'::json
			)AS photos
		FROM history_timeline ht
		WHERE ht.id = $1
	`

	result := m.DB.WithContext(ctx).Raw(query, id).Scan(&history)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get history timeline %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("history timeline not found")
	}

	return &history, nil
}

func (m *HistoryTimelineModel) GetAllHistoryTimeline(ctx context.Context, offset, limit int) ([]dto.HistoryTimelineResponse, int, error) {
	var history []dto.HistoryTimelineResponse

	var totalData int64
	if err := m.DB.WithContext(ctx).Model(&HistoryTimeline{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	if err := m.DB.WithContext(ctx).
		Model(&HistoryTimeline{}).
		Select("id, id_author, title, year, description, display_order").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(&history).
		Error; err != nil {
		return nil, 0, err
	}

	return history, int(totalData), nil
}
