package service

import (
	"context"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
)

type HistoryService struct {
	HistoryTimelineModel *entity.HistoryTimelineModel
	HistoryPhotosModel   *entity.HistoryPhotosModel
}

func NewHistoryService(m *entity.HistoryTimelineModel, n *entity.HistoryPhotosModel) *HistoryService {
	return &HistoryService{
		HistoryTimelineModel: m,
		HistoryPhotosModel:   n,
	}
}

// func (m *HistoryService) ProcessHistoryPhotos(
// 	ctx context.Context,
// 	newImages []*multipart.FileHeader,
// 	historyDetail *dto.HistoryPayload,
// ) ([]*dto.HistoryTimelineResponse, []int, error) {
// 	var all
// }

func (m *HistoryService) GetAllHistoryTimeline(ctx context.Context, offset, limit int) ([]dto.HistoryTimelineResponse, int, error) {
	history, totalData, err := m.HistoryTimelineModel.GetAllHistoryTimeline(ctx, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return history, totalData, nil
}

func (m *HistoryService) GetHistoryTimelineByID(ctx context.Context, id int) (*dto.HistoryTimelineResponse, error) {
	history, err := m.HistoryTimelineModel.GetHistoryTimelineByid(ctx, id)
	if err != nil {
		return nil, err
	}
	return history, nil
}
