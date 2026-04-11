package service

import (
	"errors"
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

type WorkService struct {
	Model         *model.WorkModel
	GalleryModel  *model.GalleryModel
	PengurusModel *model.PengurusModel
}

func NewWorkService(m *model.WorkModel, g *model.GalleryModel, p *model.PengurusModel) *WorkService {
	return &WorkService{
		Model:         m,
		GalleryModel:  g,
		PengurusModel: p,
	}
}

func (s *WorkService) CreateWork(work *model.Work) error {
	if work.Title == "" {
		return errors.New("title is required")
	}

	// Validasi Gallery
	if _, err := s.GalleryModel.GetGalleryByID(work.GalleryID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("referenced gallery_id not found")
		}
		return err
	}

	// Validasi Pengurus
	if _, err := s.PengurusModel.GetPengurusById(work.TeamProject); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("referenced pengurus_id (team_project) not found")
		}
		return err
	}

	return s.Model.InsertWork(work)
}

func (s *WorkService) GetAllWorks() ([]model.Work, error) {
	return s.Model.GetAllWorks()
}

func (s *WorkService) GetWorkByID(id int) (*model.Work, error) {
	work, err := s.Model.GetWorkById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("work not found")
		}
		return nil, err
	}
	return work, nil
}

func (s *WorkService) UpdateWork(id int, patch map[string]any) (*model.Work, error) {
	// Cek apakah data ada sebelum update
	_, err := s.Model.GetWorkById(id)
	if err != nil {
		return nil, errors.New("work not found")
	}

	return s.Model.UpdateWork(id, patch)
}

func (s *WorkService) DeleteWork(id int) error {
	err := s.Model.DeleteWork(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("work not found already")
		}
		return err
	}
	return nil
}
