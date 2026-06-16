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

	"github.com/lib/pq"
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

type WorkGalleryResult struct {
	NewGallery      []*dto.GalleryResponse
	AllgalleryIDS   []int
	PrimaryImageURl string
}

type UserData struct {
	UserRole string
	UserID   int
}

func (s *WorkService) ProcessWorkGalleries(
	ctx context.Context,
	userID int,
	newImages []*multipart.FileHeader,
	workData *dto.CreateRequestWork,
) (*WorkGalleryResult, error) {

	var allGalleryIDS []int
	primaryImageURL := ""
	if len(workData.ExistingID) > 0 {
		isValid, err := s.GalleryServices.CheckExistingGallery(workData.ExistingID)
		if err != nil {
			return nil, fmt.Errorf("failed to check existing gallery %w", err)
		}

		if !isValid {
			return nil, fmt.Errorf("invalid gallery ids")
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
		return nil, fmt.Errorf("you can only set 5 gallery to this work")
	}
	if len(workData.Technologies) > MaxKategori {
		return nil, fmt.Errorf("you can only tag %d Technologies in one work", MaxKategori)
	}

	// if there is no new image, return
	if len(newImages) == 0 {
		if len(allGalleryIDS) > 0 {
			gallery, err := s.GalleryServices.GetGalleryByID(ctx, allGalleryIDS[0])
			if err != nil {
				return nil, fmt.Errorf("failed to get gallery image URL: %w", err)
			}
			return &WorkGalleryResult{
				NewGallery:      nil,
				AllgalleryIDS:   allGalleryIDS,
				PrimaryImageURl: gallery.FileURL,
			}, nil
		}
		return &WorkGalleryResult{
			NewGallery:      nil,
			AllgalleryIDS:   allGalleryIDS,
			PrimaryImageURl: primaryImageURL,
		}, nil
	}

	now := time.Now()
	galleryName := "foto untuk work dengan judul" + workData.Title
	galleryData := &dto.GalleryInsert{
		IDUsers:     userID,
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
			return nil, fmt.Errorf("failed to open file")
		}
		fileUpload[i] = &dto.UploadFileRequest{
			FileHeader: file,
			File:       fileContent,
			Folder:     "work",
			UserID:     uint(userID),
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
			return nil, fmt.Errorf("failed to upload new image and gallery %w", err)
		}
		for _, gallery := range newgalleryDataResponse {
			allGalleryIDS = append(allGalleryIDS, gallery.ID)
		}
		if len(newgalleryDataResponse) > 0 {
			primaryImageURL = newgalleryDataResponse[0].FileURL
		}
	}

	return &WorkGalleryResult{
		NewGallery:      newgalleryDataResponse,
		AllgalleryIDS:   allGalleryIDS,
		PrimaryImageURl: primaryImageURL,
	}, nil
}

func (s *WorkService) CreateWork(
	ctx context.Context,
	work *dto.CreateRequestWork,
	newImages []*multipart.FileHeader,
	userRole string,
	userID int,
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
		return nil, fmt.Errorf("failed to set status: %w", err)
	}
	galleryrResult, err := s.ProcessWorkGalleries(
		ctx,
		userID,
		newImages,
		work,
	)
	if err != nil {
		return nil, fmt.Errorf("failed while process gallery %w", err)
	}

	newGalleryIDS := make([]int, 0, len(galleryrResult.NewGallery))
	for _, idGallery := range galleryrResult.NewGallery {
		newGalleryIDS = append(newGalleryIDS, idGallery.ID)
	}

	var workDataResponse *dto.WorkResponseClient
	err = s.WorkModel.DB.Transaction(func(tx *gorm.DB) error { // ini error
		var txFailed bool
		defer func() {
			if txFailed {
				s.GalleryServices.DeleteGalleryMultiple(ctx, newGalleryIDS)
			}
		}()

		modelWork := s.WorkModel.WithTx(tx)
		modelWorkGallery := s.WorkGallery.WithTx(tx)
		// var err error

		if _, err := s.PengurusModel.GetPengurusByID(ctx, work.PengurusID); err != nil {
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
			Technologies: pq.StringArray(technologies),
			ProjectDate:  work.ProjectDate,
			ImageURL:     galleryrResult.PrimaryImageURl, // ini error
			Status:       statusSet,
			Division:     validDivision.Divisi,
		}

		workDataResponse, err = modelWork.InsertWork(ctx, &insertData)
		if err != nil {
			txFailed = true
			log.Printf("Failed to insert work %v: %v", insertData, err)
			return fmt.Errorf("failed while insert data to database %w", err)
		}

		workGalleryData := make([]*entity.WorkGallery, len(galleryrResult.AllgalleryIDS))
		for i, id := range galleryrResult.AllgalleryIDS {
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

func (s *WorkService) GetAllWorksAndByProjectType(
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
) ([]dto.WorkResponseInternal, int64, error) {

	var (
		worksDataResponse []dto.WorkResponseInternal
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

	log.Printf("[service] validDivision: %v", validDivision)
	log.Printf("[service] validViewStatus: %v", validViewStatus)
	// call function query to get all works by division and status
	filterByDivision := true
	if userRole == constants.RoleKeySuperAdmin || userRole == constants.RoleKeyBPH {
		filterByDivision = false
	}
	log.Printf("[service] filterByDivision: %v", filterByDivision)
	worksDataResponse, totalData, err = s.WorkModel.GetAllWorksAdmin(
		ctx,
		validDivision.Divisi,
		validViewStatus,
		filterByDivision,
		offset, limit,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("failed while get the data %w", err)
	}

	return worksDataResponse, totalData, nil
}

func (s *WorkService) GetWorkByID(ctx context.Context, userRole string, id int, isPublic bool) (*dto.WorkResponseInternal, error) {

	if isPublic {
		workResponseData, err := s.WorkModel.GetWorkById(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("work not found or something wrong while fetch data %w", err)
		}

		return workResponseData, nil

	} else {
		validUser, err := authorization.GetRoleInfo(userRole)
		if err != nil {
			return nil, fmt.Errorf("role not valid: %w", err)
		}
		if validUser.Role != constants.RoleAdmin && validUser.Role != constants.RoleKoordinator {
			return nil, fmt.Errorf("you are not allowed to access this resource")
		}

		workResponseData, err := s.WorkModel.GetWorkByIDForAdmin(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("work not found or something wrong while fetch data %w", err)
		}

		return workResponseData, nil
	}

}

func (s *WorkService) UpdateStatusWork(
	ctx context.Context,
	idWork, userID int,
	targetStatus string,
	userRole string,
) (*dto.WorkUpdateResponse, error) {
	oldWorkData, err := s.WorkModel.GetWorkWithOutGallery(ctx, idWork)
	if err != nil {
		return nil, fmt.Errorf("work not found you can't do update: %w", err)
	}

	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleAdmin,
		constants.RoleKeyBPH,
	); err != nil {
		return nil, fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	// check if current status can be changed to target status
	if err = workAuthorization.CanModerateWork(userRole, oldWorkData.Status, targetStatus); err != nil {
		return nil, fmt.Errorf("error while check status: %w", err)
	}

	updatedWorkData, err := s.WorkModel.UpdateWork(idWork, map[string]any{
		"status":     targetStatus,
		"updated_at": time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update work %w", err)
	}

	return &dto.WorkUpdateResponse{
		ID:           updatedWorkData.ID,
		PengurusID:   updatedWorkData.PengurusID,
		Title:        updatedWorkData.Title,
		Tagline:      updatedWorkData.Tagline,
		Description:  updatedWorkData.Description,
		Slug:         updatedWorkData.Slug,
		ProjectType:  updatedWorkData.ProjectType,
		Technologies: updatedWorkData.Technologies,
		ProjectDate:  updatedWorkData.ProjectDate,
		ImageURL:     updatedWorkData.ImageURL,
		Status:       updatedWorkData.Status,
		UpdatedAt:    updatedWorkData.UpdatedAt,
		CreatedAt:    updatedWorkData.CreatedAt,
		WorkGallery:  nil,
	}, nil
}

func canUpdateWork(userRole string, currentStatus string) error {
	validUser, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("invalid user role")
	}

	currentStatus = strings.ToLower(strings.TrimSpace(currentStatus))

	if validUser.Role == constants.RoleAdmin {
		return nil
	}

	if validUser.Role == constants.RoleKoordinator {
		switch currentStatus {
		case constants.StatusDraft, constants.StatusRejected:
			return nil
		default:
			return fmt.Errorf("work with status %s cannot be edited koordinator", currentStatus)
		}
	}

	return fmt.Errorf("role %s cannot edit work", userRole)
}

func (s *WorkService) UpdateWorkByID(
	ctx context.Context,
	idWork, userID int,
	updatedWork *dto.WorkPatch,
	newImages []*multipart.FileHeader,
	userRole string,
) (*dto.WorkUpdateResponse, error) {

	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleAdmin,
		constants.RoleKoordinator,
	); err != nil {
		return nil, fmt.Errorf("you are not allowed to access this resource %w", err)
	}

	oldWorkData, err := s.WorkModel.GetWorkByIDForAdmin(ctx, idWork)
	if err != nil {
		return nil, fmt.Errorf("work not found you can't do update: %w", err)
	}

	if err := canUpdateWork(userRole, oldWorkData.Status); err != nil {
		return nil, fmt.Errorf("cannot update work: %w", err)
	}

	status, err := workAuthorization.CanSetStatusWork(userRole, updatedWork.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to set status %w", err)
	}

	galleryResult := &WorkGalleryResult{}
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

		galleryResult, err = s.ProcessWorkGalleries(
			ctx,
			userID,
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
		Technologies: pq.StringArray(updatedWork.Technologies),
		ProjectDate:  updatedWork.ProjectDate,
		ImageURL:     galleryResult.PrimaryImageURl,
		Status:       status,
		UpdatedAt:    time.Now(),
	}

	updatedWorkData := utils.StructToMap(workUpdateDataPayload)

	workUpdateData, err := s.WorkModel.UpdateWork(idWork, updatedWorkData)
	if err != nil {
		return nil, fmt.Errorf("failed to update work %w", err)
	}
	log.Printf("[service] workUpdateData: %v", workUpdateData)
	var workGallery []*dto.WorkGalleryResponse
	if galleryTouch {
		workGallery, err = s.WorkGallery.UpdateWorkGallery(galleryResult.AllgalleryIDS, idWork)
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

		//delete work gallery
		if err := modelWorkGallery.DeleteWorkGalleryByWorkID(ctx, workID); err != nil {
			tx.Rollback()
			return fmt.Errorf("Failed to delete work gallery data: %w", err)
		}
		// delete work
		if err := modelWork.DeleteWork(ctx, workID); err != nil {
			tx.Rollback()
			return fmt.Errorf("Failed to delete work data: %w", err)
		}
		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("Failed to commit transaction: %w", err)
		}
	}
	return nil
}

func (s *WorkService) GetWorkTypes(ctx context.Context) ([]string, error) {

	projectTypes, err := s.WorkModel.GetWorkTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get work types %w", err)
	}

	return projectTypes, nil
}
