package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strings"
	"time"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"
	"web_doscom/internal/utils"

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

func (s *WorkService) ProcessWorkGalleries(
	ctx context.Context,
	newImages []*multipart.FileHeader,
	workData *model.CreateRequestWork,
) ([]*model.GalleryResponse, []int, error) {
	var allGalleryIDS []int
	if len(workData.ExistingID) > 0 {
		isValid, err := s.GalleryServices.CheckExistingGallery(workData.ExistingID)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to check existing gallery %w", err)
		}

		if !isValid {
			return nil, nil, fmt.Errorf("invalid gallery ids")
		}

		for _, id := range workData.ExistingID {
			if id != nil {
				allGalleryIDS = append(allGalleryIDS, *id)
			}
		}
	}

	const (
		MaxGallery  = 5
		MaxKategori = 3
	)
	totalGallery := len(allGalleryIDS) + len(newImages)
	if totalGallery > MaxGallery {
		return nil, nil, fmt.Errorf("you can only set 5 gallery to this work")
	}
	if len(workData.Technologies) > MaxKategori {
		return nil, nil, fmt.Errorf("you can only tag %d Technologies in one work", MaxKategori)
	}

	now := time.Now()
	galleryName := "foto untuk work dengan judul" + workData.Title
	galleryData := &model.GalleryInsert{
		IDUsers:     workData.PengurusID,
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
			UserID:     uint(workData.PengurusID),
		}
	}

	defer func() {
		for _, file := range fileUpload {
			file.File.Close()
		}
	}()

	var newgalleryDataResponse []*model.GalleryResponse
	if len(newImages) > 0 {
		var err error
		newgalleryDataResponse, err = s.GalleryServices.UploadAndInsertGalleryMultiple(
			ctx,
			galleryData,
			fileUpload,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to upload new image and gallery %w", err)
		}
		for _, gallery := range newgalleryDataResponse {
			allGalleryIDS = append(allGalleryIDS, gallery.ID)
		}
	}

	return newgalleryDataResponse, allGalleryIDS, nil
}

func (s *WorkService) CreateWork(
	ctx context.Context,
	work *model.CreateRequestWork,
	newImages []*multipart.FileHeader,
	userRole string,
) (*model.WorkResponse, error) {

	if err := utils.CheckRolePermission(userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		return nil, fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	if work.Status == "" {
		work.Status = constants.StatusDraft
	}
	if userRole == constants.RoleKoordinator &&
		(work.Status != constants.StatusDraft || work.Status != constants.StatusPending) {
		return nil, fmt.Errorf("you are can only set status to draft or pending")
	}

	newgGalleryDataResponse, allgalleryIDS, err := s.ProcessWorkGalleries(
		ctx,
		newImages,
		work,
	)
	if err != nil {
		return nil, fmt.Errorf("failed while process gallery %w", err)
	}

	newGalleryIDS := make([]int, 0, len(newgGalleryDataResponse))
	for _, idGallery := range newgGalleryDataResponse {
		newGalleryIDS = append(newGalleryIDS, idGallery.ID)
	}

	var workDataResponse *model.WorkResponse
	err = s.WorkModel.DB.Transaction(func(tx *gorm.DB) error {

		var txFailed bool
		defer func() {
			if txFailed {
				s.GalleryServices.DeleteGalleryMultiple(ctx, newGalleryIDS)
			}
		}()

		modelWork := s.WorkModel.WithTx(tx)
		modelWorkGallery := s.WorkGallery.WithTx(tx)
		var err error

		if _, err := s.PengurusModel.GetPengurusById(ctx, work.PengurusID); err != nil {
			return fmt.Errorf("pengurus tidak valid atau tidak ditemukan %w", err)
		}

		technologies := make([]string, len(work.Technologies))
		for i, tech := range work.Technologies {
			technologies[i] = strings.ToLower(strings.TrimSpace(tech))
		}
		insertData := model.Work{
			PengurusID:   work.PengurusID,
			Title:        work.Title,
			Tagline:      work.Tagline,
			Description:  work.Description,
			Slug:         work.Slug,
			ProjectType:  strings.ToLower(work.ProjectType),
			Technologies: technologies,
			ProjectDate:  work.ProjectDate,
			ImageURL:     newgGalleryDataResponse[0].FileURL,
			Status:       work.Status,
		}
		workDataResponse, err = modelWork.InsertWork(ctx, &insertData)
		if err != nil {
			txFailed = true
			log.Printf("Failed to insert work %v: %v", insertData, err)
			return fmt.Errorf("failed while insert data to database %w", err)
		}

		workGalleryData := make([]*model.WorkGallery, len(allgalleryIDS))
		for i, id := range allgalleryIDS {
			workGalleryData[i] = &model.WorkGallery{
				IDWork:    workDataResponse.ID,
				IDGallery: id,
			}
		}

		if err := modelWorkGallery.InsertWorkGalleryMultiple(ctx, workGalleryData); err != nil {
			txFailed = true
			log.Printf("Failed to insert work gallery %v: %v", workGalleryData, err)
			return fmt.Errorf("failed to insert work gallery %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert database and Transaction failed %w", err)
	}

	return workDataResponse, nil

}

func (s *WorkService) GetAllWorksAndByTechnologi(
	ctx context.Context,
	offset, limit int,
	filterProjectType string,
) ([]model.WorkResponse, int64, error) {

	var (
		worksDataResponse []model.WorkResponse
		totalData         int64
		err               error
	)

	filterProjectType = strings.ToLower(strings.TrimSpace(filterProjectType))

	if filterProjectType == "" {
		worksDataResponse, totalData, err = s.WorkModel.GetAllWorks(ctx, offset, limit)
	} else {
		worksDataResponse, totalData, err = s.WorkModel.GetAllWorksByProjectType(ctx, offset, limit, filterProjectType)
	}

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get data %w", err)
	}

	return worksDataResponse, totalData, err
}

func (s *WorkService) GetWorkByID(ctx context.Context, id int) (*model.WorkResponse, error) {
	workResponseData, err := s.WorkModel.GetWorkById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("work not found or something wrong while fetch data %w", err)
	}
	return workResponseData, nil
}

func (s *WorkService) UpdateWorkByID(
	ctx context.Context,
	idWork int,
	updatedWork *model.WorkPatch,
	newImages []*multipart.FileHeader,
) (*model.WorkResponse, error) {
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
