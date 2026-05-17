package service

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"strings"
	"time"
	"web_doscom/internal/authorization"
	workAuthorization "web_doscom/internal/authorization/work"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/utils"

	"gorm.io/gorm"
)

type WorkService struct {
	WorkModel       *entity.WorkModel
	GalleryServices *GalleryService
	PengurusModel   *entity.PengurusModel
	WorkGallery     *entity.WorkGalleryModel
}

func NewWorkService(
	m *entity.WorkModel,
	g *GalleryService,
	p *entity.PengurusModel,
	w *entity.WorkGalleryModel,
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
	workData *dto.CreateRequestWork,
) ([]*dto.GalleryResponse, []int, error) {

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

	// if there is no new image, return
	if len(newImages) == 0 {
		return nil, allGalleryIDS, nil
	}

	now := time.Now()
	galleryName := "foto untuk work dengan judul" + workData.Title
	galleryData := &dto.GalleryInsert{
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

	fileUpload := make([]*dto.UploadFileRequest, len(newImages))
	for i, file := range newImages {
		fileContent, err := file.Open()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file")
		}
		fileUpload[i] = &dto.UploadFileRequest{
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

	var newgalleryDataResponse []*dto.GalleryResponse
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
	work *dto.CreateRequestWork,
	newImages []*multipart.FileHeader,
	userRole string,
) (*dto.WorkResponseClient, error) {

	validDivision, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return nil, fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	if err := authorization.CheckRolePermission(userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		return nil, fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	statusSet, err := workAuthorization.CanSetStatusWork(userRole, work.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to set status %w", err)
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

	var workDataResponse *dto.WorkResponseClient
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
		insertData := entity.Work{
			PengurusID:   work.PengurusID,
			Title:        work.Title,
			Tagline:      work.Tagline,
			Description:  work.Description,
			Slug:         work.Slug,
			ProjectType:  strings.ToLower(work.ProjectType),
			Technologies: technologies,
			ProjectDate:  work.ProjectDate,
			ImageURL:     newgGalleryDataResponse[0].FileURL,
			Status:       statusSet,
			Division:     validDivision.Divisi,
		}

		workDataResponse, err = modelWork.InsertWork(ctx, &insertData)
		if err != nil {
			txFailed = true
			log.Printf("Failed to insert work %v: %v", insertData, err)
			return fmt.Errorf("failed while insert data to database %w", err)
		}

		workGalleryData := make([]*entity.WorkGallery, len(allgalleryIDS))
		for i, id := range allgalleryIDS {
			workGalleryData[i] = &entity.WorkGallery{
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
) ([]dto.WorkResponseClient, int64, error) {

	var (
		worksDataResponse []dto.WorkResponseClient
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

func (s *WorkService) GetWorksByDivision(
	ctx context.Context,
	userRole string,
	limit, offset int,
) ([]dto.WorkResponseClient, int64, error) {

	var (
		worksDataResponse []dto.WorkResponseClient
		totalData         int64
		err               error
	)

	validDivision, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return nil, 0, fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	validViewStatus, err := workAuthorization.GetViewableStatus(userRole)
	if err != nil {
		return nil, 0, fmt.Errorf("there is error: %w", err)
	}

	// call function query to get all works by division and status
	worksDataResponse, totalData, err = s.WorkModel.GetAllWorksAdmin(
		ctx,
		validDivision.Divisi,
		validViewStatus,
		offset,
		limit,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("failed while get the data %w", err)
	}

	return worksDataResponse, totalData, nil
}

func (s *WorkService) GetWorkByID(ctx context.Context, id int) (*dto.WorkResponseClient, error) {
	workResponseData, err := s.WorkModel.GetWorkById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("work not found or something wrong while fetch data %w", err)
	}
	return workResponseData, nil
}

func (s *WorkService) UpdateWorkByID(
	ctx context.Context,
	idWork int,
	updatedWork *dto.WorkPatch,
	newImages []*multipart.FileHeader,
	userRole string,
) (*dto.WorkUpdateResponse, error) {

	_, err := s.WorkModel.GetWorkById(ctx, idWork)
	if err != nil {
		return nil, fmt.Errorf("work not found you can't do update %w", err)
	}

	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		return nil, fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	status, err := workAuthorization.CanSetStatusWork(userRole, updatedWork.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to set status %w", err)
	}

	var (
		newGalleryDataResponse []*dto.GalleryResponse
		allGalleryIDS          []int
	)

	galleryTouch := len(newImages) != 0 || len(updatedWork.ExistingID) != 0
	if galleryTouch {
		workUpdate := &dto.CreateRequestWork{
			ExistingID:   updatedWork.ExistingID,
			PengurusID:   updatedWork.PengurusID,
			Title:        updatedWork.Title,
			Tagline:      updatedWork.Tagline,
			Description:  updatedWork.Description,
			Slug:         updatedWork.Slug,
			ProjectType:  updatedWork.ProjectType,
			Technologies: updatedWork.Technologies,
			ProjectDate:  updatedWork.ProjectDate,
			Status:       status,
		}

		newGalleryDataResponse, allGalleryIDS, err = s.ProcessWorkGalleries(
			ctx,
			newImages,
			workUpdate,
		)
	}

	workUpdateDataPayload := &dto.WorkPayloadUpdate{
		PengurusID:   updatedWork.PengurusID,
		Title:        updatedWork.Title,
		Tagline:      updatedWork.Tagline,
		Description:  updatedWork.Description,
		Slug:         updatedWork.Slug,
		ProjectType:  updatedWork.ProjectType,
		Technologies: updatedWork.Technologies,
		ProjectDate:  updatedWork.ProjectDate,
		Status:       status,
		UpdatedAt:    time.Now(),
	}

	if len(newGalleryDataResponse) > 0 {
		workUpdateDataPayload.ImageURL = newGalleryDataResponse[0].FileURL
	}

	updatedWorkData := utils.StructToMap(workUpdateDataPayload)

	workUpdateData, err := s.WorkModel.UpdateWork(idWork, updatedWorkData)
	if err != nil {
		return nil, fmt.Errorf("failed to update work %w", err)
	}

	var workGallery []*dto.WorkGalleryResponse
	if galleryTouch {
		workGallery, err = s.WorkGallery.UpdateWorkGallery(allGalleryIDS, idWork)
		if err != nil {
			return nil, fmt.Errorf("failed to update work gallery %w", err)
		}
	}

	workUpdateDataResponse := &dto.WorkUpdateResponse{
		ID:           workUpdateData.ID,
		PengurusID:   workUpdateData.PengurusID,
		Title:        workUpdateData.Title,
		Tagline:      workUpdateData.Tagline,
		Description:  workUpdateData.Description,
		Slug:         workUpdateData.Slug,
		ProjectType:  workUpdateData.ProjectType,
		Technologies: workUpdateData.Technologies,
		ProjectDate:  workUpdateData.ProjectDate,
		ImageURL:     workUpdateData.ImageURL,
		Status:       workUpdateData.Status,
		UpdatedAt:    workUpdateData.UpdatedAt,
		CreatedAt:    workUpdateData.CreatedAt,
		WorkGallery:  workGallery,
	}
	return workUpdateDataResponse, nil
}

func (s *WorkService) DeleteWork(ctx context.Context, workID int, userRole string) error {

	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("invalid user role")
	}

	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		return fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	workData, err := s.WorkModel.GetWorkWithOutGallery(ctx, workID)
	if err != nil {
		return fmt.Errorf("work not found, the process cannot continue")
	}

	grantedDelete, err := workAuthorization.CanDeleteWork(userRole, workData.Status)
	if err != nil {
		return fmt.Errorf("something wrong or denied: %w", err)
	}
	if grantedDelete {
		tx := s.WorkModel.DB.Begin()
		if tx.Error != nil {
			return fmt.Errorf("failed to begin transaction: %w", tx.Error)
		}

		modelWorkGallery := s.WorkGallery.WithTx(tx)
		modelWork := s.WorkModel.WithTx(tx)
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()

		// delete work first
		if err := modelWork.DeleteWork(ctx, workID); err != nil {
			tx.Rollback()
			return fmt.Errorf("Failed to delete work data: %w", err)
		}

		//delete work gallery
		if err := modelWorkGallery.DeleteWorkGalleryByID(ctx, workID); err != nil {
			tx.Rollback()
			return fmt.Errorf("Failed to delete work gallery data: %w", err)
		}
		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("Failed to commit transaction: %w", err)
		}
	}
	return nil
}
