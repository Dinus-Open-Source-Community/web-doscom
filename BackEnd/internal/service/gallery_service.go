package service

import (
	"context"
	"fmt"
	"math"
	"mime/multipart"
	"time"

	"web_doscom/internal/database/model"
)

type GalleryService struct {
	Model   *model.GalleryModel
	Storage *StorageService
}

func NewGalleryService(m *model.GalleryModel, s *StorageService) *GalleryService {
	return &GalleryService{Model: m, Storage: s}
}

const (
	maxUploadSize = 20 << 20 // 20mb
	maxFileSize   = 5 << 20  // 5mb
)

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
	gallery *model.GalleryInsert,
	fileUpload *model.UploadFileRequest,
) (*model.GalleryResponse, string, error) {
	// insert file first
	fileURL, fileUploadID, err := m.Storage.UploadFileAndCreateMetadata(ctx, fileUpload)
	if err != nil {
		return nil, "", err
	}

	// insert gallery
	galleryUpload := &model.Gallery{
		IDUsers:      gallery.IDUsers,
		FileUploadID: fileUploadID,
		GalleryName:  gallery.GalleryName,
		GalleryType:  gallery.GalleryType,
		Description:  gallery.Description,
		EventDate:    gallery.EventDate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	galleryResponse, err := m.Model.InsertGallery(galleryUpload)
	if err != nil {
		return nil, "", err
	}

	return &model.GalleryResponse{
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
	gallery *model.GalleryInsert,
	fileUpload []*model.UploadFileRequest,
) ([]*model.GalleryResponse, error) {
	fileUploadHeader := make([]*multipart.FileHeader, len(fileUpload))
	for i, file := range fileUpload {
		fileUploadHeader[i] = file.FileHeader
	}
	fileUrl, fileUploadID, err := m.Storage.UploadFileAndCreateMetadataMultiple(
		ctx,
		fileUploadHeader,
		fileUpload[0].Folder,
		int(fileUpload[0].UserID),
	)
	if err != nil {
		return nil, err
	}

	galleryUpload := make([]*model.Gallery, len(fileUpload))
	for i, _ := range fileUpload {
		galleryUpload[i] = &model.Gallery{
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

	// upload to database
	responseGallery, err := m.Model.InsertGalleryMultiple(galleryUpload)
	if err != nil {
		return nil, err
	}

	return responseGallery, nil
}

// wrapper for get gallery by type
func (m *GalleryService) GetAllGalleryByDate(
	ctx context.Context,
	startDate, endDate string,
	limit, offset int,
) ([]*model.GalleryResponse, int64, int64, error) {

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

	var response []*model.GalleryResponse
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
		response = append(response, &model.GalleryResponse{
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

// wrapper for delete gallery
func (m *GalleryService) DeleteGallery(id int) error {
	return m.Model.DeleteGallery(id)
}
