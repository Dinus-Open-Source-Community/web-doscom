package service

import "web_doscom/internal/database/model"

type PengurusService struct {
	PengurusModel *model.PengurusModel
	GalleryModel  *GalleryService
}

func NewPengurusService(m *model.PengurusModel, g *GalleryService) *PengurusService {
	return &PengurusService{
		PengurusModel: m,
		GalleryModel:  g,
	}
}
