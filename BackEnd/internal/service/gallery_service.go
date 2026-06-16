package service

import (
	"context"
	"fmt"
	"log"
	"math"
	"mime/multipart"
	"time"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
)

type GalleryService struct {
	Model   *entity.GalleryModel
	Storage *StorageService
}

func NewGalleryService(m *entity.GalleryModel, s *StorageService) *GalleryService {
	return &GalleryService{Model: m, Storage: s}
}

type validateFile struct {
	Fileheader *multipart.FileHeader
	MimeType   string
	Folder     string
	fileSize   int64
	kategori   string
}

func ParseYearRange(startYear, endYear string) (*time.Time, *time.Time, error) {
	// parse from string to time
	start, err := time.Parse("2006", startYear)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse time %w", err)
	}

	end, err := time.Parse("2006", endYear)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse time %w", err)
	}

	startTime := time.Date(start.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(end.Year(), 12, 31, 23, 59, 59, 0, time.UTC)

	return &startTime, &endTime, nil
}

func (m *GalleryService) InsertGalleryAndFileUpload(
	ctx context.Context,
	gallery *dto.GalleryInsert,
	fileUpload *dto.UploadFileRequest,
) (*dto.GalleryResponse, string, error) {
	// insert file first
	fileURL, fileUploadID, err := m.Storage.UploadFileAndCreateMetadata(ctx, fileUpload)
	if err != nil {
		return nil, "", err
	}

	// insert gallery
	log.Printf("[gallery service] id users: %d", gallery.IDUsers)
	galleryUpload := &entity.Gallery{
		IDUsers:      gallery.IDUsers,
		FileUploadID: fileUploadID,
		GalleryName:  gallery.GalleryName,
		GalleryType:  gallery.GalleryType,
		Description:  gallery.Description,
		EventDate:    gallery.EventDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	log.Printf("[gallery service] gallery: %d", galleryUpload.IDUsers)
	galleryResponse, err := m.Model.InsertGallery(galleryUpload)
	if err != nil {
		return nil, "", err
	}

	return &dto.GalleryResponse{
		ID:           galleryResponse.ID,
		IDUsers:      galleryResponse.IDUsers,
		FileUploadID: galleryResponse.FileUploadID,
		GalleryName:  galleryResponse.GalleryName,
		GalleryType:  galleryResponse.GalleryType,
		Description:  galleryResponse.Description,
		EventDate:    galleryResponse.EventDate,
	}, fileURL, nil
}

func (m *GalleryService) UploadAndInsertGalleryMultiple(
	ctx context.Context,
	gallery *dto.GalleryInsert,
	fileUpload []*dto.UploadFileRequest,
) ([]*dto.GalleryResponse, error) {
	// Gunakan background context agar upload tidak terputus jika client disconnect prematur
	uploadCtx := context.Background()

	fileUploadHeader := make([]*multipart.FileHeader, len(fileUpload))
	for i, file := range fileUpload {
		fileUploadHeader[i] = file.FileHeader
	}
	fileUrl, fileUploadID, err := m.Storage.UploadFileAndCreateMetadataMultiple(
		uploadCtx,
		fileUploadHeader,
		fileUpload[0].Folder,
		int(fileUpload[0].UserID),
	)
	if err != nil {
		return nil, err
	}

	galleryUpload := make([]*entity.Gallery, len(fileUpload))
	for i, _ := range fileUpload {
		galleryUpload[i] = &entity.Gallery{
			IDUsers:      gallery.IDUsers,
			FileUploadID: fileUploadID[i],
			GalleryName:  gallery.GalleryName,
			GalleryType:  gallery.GalleryType,
			Description:  gallery.Description,
			EventDate:    gallery.EventDate,
			FileURL:      fileUrl[i],
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
	}

	// upload to database menggunakan uploadCtx agar tidak terputus
	responseGallery, err := m.Model.InsertGalleryMultiple(galleryUpload)
	if err != nil {
		return nil, err
	}

	return responseGallery, nil
}

func (m *GalleryService) GetAllGalleryAndByDate(
	ctx context.Context,
	startDate, endDate string,
	limit, offset int,
) ([]*dto.GalleryResponse, int64, int64, error) {

	var (
		dateStart, dateEnd *time.Time
		err                error
	)
	if startDate != "" || endDate != "" {
		dateStart, dateEnd, err = ParseYearRange(startDate, endDate)
		if err != nil {
			return nil, 0, 0, err
		}
	} else {
		dateStart = nil
		dateEnd = nil
	}

	var response []*dto.GalleryResponse
	galleries, count, err := m.Model.GetAllGalleryAndByYear(
		ctx,
		dateStart,
		dateEnd,
		limit,
		offset,
	)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error while get data %w", err)
	}

	for _, data := range galleries {
		response = append(response, &dto.GalleryResponse{
			ID:          data.ID,
			GalleryName: data.GalleryName,
			GalleryType: data.GalleryType,
			Description: data.Description,
			EventDate:   data.EventDate,
			FileURL:     data.FileURL,
		})
	}

	totalPage := int(math.Ceil(float64(count) / float64(limit)))
	currentPage := (offset / limit) + 1

	return response, int64(totalPage), int64(currentPage), nil
}

func (m *GalleryService) GetGalleryByID(ctx context.Context, id int) (dto.GalleryResponse, error) {
	galleryData, err := m.Model.GetGalleryByID(ctx, id)
	if err != nil {
		return dto.GalleryResponse{}, fmt.Errorf("gallery not found %w", err)
	}
	return dto.GalleryResponse{
		ID:          galleryData.ID,
		GalleryName: galleryData.GalleryName,
		GalleryType: galleryData.GalleryType,
		Description: galleryData.Description,
		EventDate:   galleryData.EventDate,
		FileURL:     galleryData.FileURL,
	}, nil
}

func (m *GalleryService) DeleteGallery(ctx context.Context, id int) error {
	galleryData, err := m.Model.GetGalleryByID(ctx, id)
	if err != nil {
		return fmt.Errorf("gallery not found %w", err)
	}

	fileData, err := m.Storage.GetFileUploadByID(galleryData.FileUploadID)
	if err != nil {
		return fmt.Errorf("file upload not found %w", err)
	}

	if err := m.Storage.DeleteFile(ctx, fileData.StoredFilename); err != nil {
		return fmt.Errorf("failed to delete file from the bucket: %w", err)
	}

	tx := m.Model.DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	// kurang inject tx

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := m.Model.DeleteGallery(galleryData.ID); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to gallery: %w", err)
	}

	if err := m.Storage.DeleteFileById(int(fileData.ID)); err != nil {
		tx.Rollback()
		log.Println("orphand DB record s3 file already deleted but DB failed to delete",
			"file_upload_id", fileData.ID,
			"stored_filename", fileData.StoredFilename,
			"error: %w", err,
		)
		return fmt.Errorf("failed to delete file data: %w", err)
	}

	if err := m.Model.DeleteGallery(galleryData.ID); err != nil {
		tx.Rollback()
		log.Println("orphand DB record s3 file already deleted but DB failed to delete",
			"gallery_id", galleryData.ID,
			"file_upload_id", fileData.ID,
			"error: %w", err,
		)

		return fmt.Errorf("failed to gallery: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil

}

func (m *GalleryService) DeleteGalleryMultiple(ctx context.Context, galleryIDS []int) error {

	_, err := m.Model.GetGalleryByIDMultiple(ctx, galleryIDS)
	if err != nil {
		return fmt.Errorf("gallery not found %w", err)
	}

	fileData, err := m.Storage.GetFileUploadByIDMultiple(ctx, galleryIDS)
	if err != nil {
		return fmt.Errorf("file upload not found %w", err)
	}

	fileIDToDelete := make([]int, 0, len(fileData))
	fileNamesToDelete := make([]string, 0, len(fileData))
	for _, file := range fileData {
		fileNamesToDelete = append(fileNamesToDelete, file.StoredFilename)
		fileIDToDelete = append(fileIDToDelete, int(file.ID))
	}
	if err := m.Storage.DeleteFileMultiple(ctx, fileNamesToDelete); err != nil {
		return fmt.Errorf("failed to delete file from bucket: %w", err)
	}

	tx := m.Model.DB.Begin()
	if tx.Error != nil {
		log.Printf("Warning: file deleted from storage but failed to begin transaction to delete from database")
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	galleryModel := m.Model.WithTx(tx)
	storageModel := m.Storage.WithTx(tx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Panic: file deleted from storage but failed to delete from database")
		}
	}()

	if err := galleryModel.DeleteGalleryByIdMultiple(ctx, galleryIDS); err != nil {
		tx.Rollback()
		log.Printf("orphand DB record s3 file already deleted but DB failed to delete for ids %v: %v", galleryIDS, err)
		return fmt.Errorf("failed to delete gallery %w", err)
	}

	if err := storageModel.DeleteFileUploadByIdMultiple(ctx, fileIDToDelete); err != nil {
		tx.Rollback()

		log.Printf("orphand DB record s3 file already deleted but DB failed to delete for ids %v: %v", fileIDToDelete, err)
		return fmt.Errorf("failed to delete file data: %w", err)
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Critical: file deleted and DB changes made but commit transaction failed for gallery ids %v: %v", galleryIDS, err)
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (s *GalleryService) CheckExistingGallery(id []*int) (bool, error) {
	if len(id) == 0 {
		return false, fmt.Errorf("id gallery is required not empty")
	}

	return s.Model.CheckExistingGallery(id)
}
