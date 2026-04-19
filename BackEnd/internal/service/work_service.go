package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

type WorkService struct {
	WorkModel       *model.WorkModel
	GalleryServices *GalleryService
	PengurusModel   *model.PengurusModel
	WorkGallery     *model.WorkGalleryModel
}

func NewWorkService(
	m *model.WorkModel,
	g *GalleryService,
	p *model.PengurusModel,
	w *model.WorkGalleryModel,
) *WorkService {
	return &WorkService{
		WorkModel:       m,
		GalleryServices: g,
		PengurusModel:   p,
		WorkGallery:     w,
	}
}

func (s *WorkService) CreateWork(
	ctx context.Context,
	work *model.RegisterWork,
	newImages []*multipart.FileHeader,
	userRole string,
) (*model.WorkResponse, error) {

	validRole, err := constants.GetRoleInfo(userRole)
	if err != nil {
		return nil, fmt.Errorf("invalid role %w", err)
	}
	if validRole.Role != constants.RoleAdmin && validRole.Role != constants.RoleKoordinator {
		return nil, fmt.Errorf("you are not allowed to access this resource")
	}

	// process gallery
	galleryDataResponse, galleryIDS, err := s.ProcessWorkGalleries(
		ctx,
		work.ExistingID,
		newImages,
		work,
	)
	if err != nil {
		return nil, fmt.Errorf("failed while process gallery %w", err)
	}

	err = s.WorkModel.DB.Transaction(func(tx *gorm.DB) error {

		var txFailed bool
		defer func() {
			if txFailed {
				tx.Rollback()
			}
		}

		modelWork := s.WorkModel.WithTx(tx)
		modelGallery := s.WorkGallery.WithTx(tx)
		var err error

		if _, err := s.PengurusModel.GetPengurusById(ctx, work.PengurusID); err != nil {
			return nil, fmt.Errorf("pengurus tidak valid atau tidak ditemukan %w", err)
		}

		insertData := model.Work{
			PengurusID:   work.PengurusID,
			Title:        work.Title,
			Tagline:      work.Tagline,
			Description:  work.Description,
			Slug:         work.Slug,
			ProjectType:  work.ProjectType,
			Technologies: work.Technologies,
			ProjectDate:  work.ProjectDate,
		}
		workData, err := s.WorkModel.InsertWork(ctx, &insertData)
		if err != nil {
			return nil, fmt.Errorf("failed while insert data to database %w", err)
		}

		workGalleryData := make([]*model.WorkGallery, len(galleryIDS))
		for i, id := range galleryIDS {
			workGalleryData[i] = &model.WorkGallery{
				IDWork:    workData.ID,
				IDGallery: id,
			}
		}
	})

	return s.Model.InsertWork(work)
}

func (s *WorkService) ProcessWorkGalleries(
	ctx context.Context,
	existingID []*int,
	newImages []*multipart.FileHeader,
	work *model.RegisterWork,
) ([]*model.GalleryResponse, []int, error) {
	var galleryIDS []int
	if len(existingID) > 0 {
		isValid, err := s.GalleryServices.CheckExistingGallery(existingID)

		if err != nil {
			return nil, nil, fmt.Errorf("failed to check existing gallery %w", err)
		}

		if !isValid {
			return nil, nil, fmt.Errorf("invalid gallery ids")
		}

		for _, id := range existingID {
			if id != nil {
				galleryIDS = append(galleryIDS, *id)
			}
		}
	}

	const (
		MaxGallery  = 5
		MaxKategori = 3
	)
	totalGallery := len(galleryIDS) + len(newImages)
	if totalGallery > MaxGallery {
		return nil, nil, fmt.Errorf("you can only set 5 gallery to this work")
	}
	if len(work.Technologies) > MaxKategori {
		return nil, nil, fmt.Errorf("you can only set 3 Technologies to this work")
	}

	now := time.Now()
	galleryName := "foto untuk work dengan judul" + work.Title
	galleryData := &model.GalleryInsert{
		IDUsers:     work.PengurusID,
		GalleryName: galleryName,
		GalleryType: "work",
		Description: "foto untuk kepentingan work",
		EventDate: time.Date(
			now.Year(),
			now.Month(),
			now.Day(), 0, 0, 0, 0, time.UTC,
		),
	}

	fileUpload := make([]*model.UploadFileRequest, len(newImages))
	for i, file := range newImages {
		fileContent, err := file.Open()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file")
		}
		fileUpload[i] = &model.UploadFileRequest{
			FileHeader: file,
			File:       fileContent,
			Folder:     "work",
			UserID:     uint(work.PengurusID),
		}
	}

	defer func() {
		for _, file := range fileUpload {
			file.File.Close()
		}
	}()

	// upload foto baru dan insert ke database
	var galleryDataResponse []*model.GalleryResponse
	if len(newImages) > 0 {
		var err error
		galleryDataResponse, err = s.GalleryServices.UploadAndInsertGalleryMultiple(
			ctx,
			galleryData,
			fileUpload,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to upload new image and gallery %w", err)
		}
		for _, gallery := range galleryDataResponse {
			galleryIDS = append(galleryIDS, gallery.ID)
		}
	}

	return galleryDataResponse, galleryIDS, nil
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
