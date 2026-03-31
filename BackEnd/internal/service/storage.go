package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"web_doscom/internal/config"
	"web_doscom/internal/database/model"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"golang.org/x/sync/errgroup"
)

type StorageService struct {
	minioClient *config.MinioClient
	fileUpload  *model.FileUploadModel
}

func NewStorageService(minioClient *config.MinioClient, fileUploadModel *model.FileUploadModel) *StorageService {
	return &StorageService{
		minioClient: minioClient,
		fileUpload:  fileUploadModel,
	}
}

// AllowedImageExtensions defines allowed image file extensions
var AllowedImageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg"}

// MaxFileSize defines maximum file size (10MB)
const MaxFileSize = 5 * 1024 * 1024

// ValidateImageFile validates file extension, size, and actual content
func (s *StorageService) ValidateImageFile(fileHeader *multipart.FileHeader) error {
	// Check file size
	if fileHeader.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds maximum limit of 10MB")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	isValid := slices.Contains(AllowedImageExtensions, ext)

	if !isValid {
		return fmt.Errorf("invalid file extension. Allowed: %v", AllowedImageExtensions)
	}

	// Validate actual file content (magic numbers)
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file for validation: %w", err)
	}
	defer file.Close()

	// Read first 512 bytes for content type detection
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to read file content: %w", err)
	}

	// Detect content type
	contentType := http.DetectContentType(buffer[:n])

	// Allow image types and SVG (which is detected as text/xml)
	if !strings.HasPrefix(contentType, "image/") &&
		contentType != "text/xml; charset=utf-8" &&
		contentType != "text/plain; charset=utf-8" { // SVG can be detected as text
		return fmt.Errorf("file is not a valid image (detected type: %s)", contentType)
	}

	// Reset file pointer to beginning for upload
	if seeker, ok := file.(io.Seeker); ok {
		seeker.Seek(0, 0)
	}

	return nil
}

// UploadFile uploads a file to MinIO and saves metadata to database
func (s *StorageService) UploadFile(
	ctx context.Context,
	file multipart.File,
	fileHeader *multipart.FileHeader,
	folder string,
) (*model.FileUpload, error) {
	// Validate file
	if err := s.ValidateImageFile(fileHeader); err != nil {
		return nil, err
	}

	// Generate unique filename
	ext := filepath.Ext(fileHeader.Filename)
	storedFilename := fmt.Sprintf("%s/%s-%s%s", folder, time.Now().Format("20060102"), uuid.New().String(), ext)

	// Determine content type
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload to MinIO
	_, err := s.minioClient.Client.PutObject(
		ctx,
		s.minioClient.BucketName,
		storedFilename,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	// Generate file URL using public URL from environment
	publicURL := os.Getenv("MINIO_PUBLIC_URL")
	if publicURL == "" {
		// Fallback to internal URL if not set
		publicURL = fmt.Sprintf("http://%s", s.minioClient.Client.EndpointURL().Host)
	}

	fileURL := fmt.Sprintf("%s/%s/%s",
		publicURL,
		s.minioClient.BucketName,
		storedFilename,
	)

	// return metadata to be saved in database
	fileUpload := &model.FileUpload{
		Category:         folder,
		OriginalFilename: fileHeader.Filename,
		StoredFilename:   storedFilename,
		FileSize:         fileHeader.Size,
		ContentType:      contentType,
		FileURL:          fileURL,
	}

	return fileUpload, nil
}

func (s *StorageService) UploadFileAndCreateMetadata(ctx context.Context, request *model.UploadFileRequest) (string, int, error) {
	// upload file to minIO
	fileURL, err := s.UploadFile(ctx, request.File, request.FileHeader, request.Folder)
	if err != nil {
		return "", 0, err
	}

	uploadedFile := &model.FileUpload{
		UserID:           request.UserID,
		Category:         fileURL.Category,
		OriginalFilename: fileURL.OriginalFilename,
		StoredFilename:   fileURL.StoredFilename,
		FileSize:         fileURL.FileSize,
		ContentType:      fileURL.ContentType,
		FileURL:          fileURL.FileURL,
	}

	// create record in database
	responseFileUpload, err := s.fileUpload.CreateMetaData(uploadedFile)
	if err != nil {
		s.minioClient.Client.RemoveObject(
			ctx,
			s.minioClient.BucketName,
			fileURL.StoredFilename,
			minio.RemoveObjectOptions{},
		)
	}

	return responseFileUpload.FileURL, int(responseFileUpload.ID), nil

}

func (s *StorageService) UploadFileAndCreateMetadataMultiple(
	ctx context.Context,
	files []*multipart.FileHeader,
	folder string,
	currentUserID int,
) ([]string, []int, error) {

	// multiple file upload max = 10 files
	if len(files) > 10 {
		return nil, nil, fmt.Errorf("maximum file upload is 5")
	}

	// validate file first
	for _, file := range files {
		if err := s.ValidateImageFile(file); err != nil {
			return nil, nil, err
		}
		fileContent, err := file.Open()
		if err != nil {
			return nil, nil, err
		}
		defer fileContent.Close()
	}

	result := make([]*model.FileUpload, len(files))
	eg, ctx := errgroup.WithContext(ctx)
	// upload file to minIO
	for i, fileHeader := range files {
		i, fileHeader := i, fileHeader
		eg.Go(func() error {
			file, err := fileHeader.Open()
			if err != nil {
				return fmt.Errorf("failed to open file")
			}
			defer file.Close()

			// upload file to minio
			fileUpload, err := s.UploadFile(ctx, file, fileHeader, folder)
			if err != nil {
				return fmt.Errorf("failed to upload file")
			}

			result[i] = fileUpload
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}

	// upload to database
	responseFileUpload, err := s.fileUpload.CreateMetaDataMultiple(ctx, result)
	if err != nil {
		return nil, nil, err
	}

	var urlFileUpload []string
	var idFileUpload []int
	for _, file := range responseFileUpload {
		urlFileUpload = append(urlFileUpload, file.FileURL)
		idFileUpload = append(idFileUpload, int(file.ID))
	}

	return urlFileUpload, idFileUpload, nil
}

// DeleteFile deletes a file from MinIO
func (s *StorageService) DeleteFile(ctx context.Context, filename string) error {
	err := s.minioClient.Client.RemoveObject(
		ctx,
		s.minioClient.BucketName,
		filename,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	return nil
}

func (s *StorageService) UpdateFile(
	ctx context.Context,
	oldfilename string,
	request *model.UploadFileRequest,
) (string, error) {
	// remove old file first
	err := s.DeleteFile(ctx, oldfilename)
	if err != nil {
		return "", err
	}

	// upload new file
	newFileMetadata, err := s.UploadFile(ctx, request.File, request.FileHeader, request.Folder)
	if err != nil {
		return "", err
	}

	// update file_url in database
	updatedMetadata, err := s.fileUpload.UpdateFileUpload(int(request.UserID), newFileMetadata)
	if err != nil {
		return "", err
	}

	return updatedMetadata.FileURL, nil
}

// GetPresignedURL generates a presigned URL for private file access
func (s *StorageService) GetPresignedURL(ctx context.Context, filename string, expiry time.Duration) (string, error) {
	url, err := s.minioClient.Client.PresignedGetObject(
		ctx,
		s.minioClient.BucketName,
		filename,
		expiry,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return url.String(), nil
}

// ListFiles lists all files in a specific folder
func (s *StorageService) ListFiles(ctx context.Context, folder string) ([]string, error) {
	var files []string

	objectCh := s.minioClient.Client.ListObjects(
		ctx,
		s.minioClient.BucketName,
		minio.ListObjectsOptions{
			Prefix:    folder,
			Recursive: true,
		},
	)

	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("error listing objects: %w", object.Err)
		}
		files = append(files, object.Key)
	}

	return files, nil
}

// CopyFile copies a file to a new location
func (s *StorageService) CopyFile(ctx context.Context, srcFilename, destFilename string) error {
	src := minio.CopySrcOptions{
		Bucket: s.minioClient.BucketName,
		Object: srcFilename,
	}

	dst := minio.CopyDestOptions{
		Bucket: s.minioClient.BucketName,
		Object: destFilename,
	}

	_, err := s.minioClient.Client.CopyObject(ctx, dst, src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

// DownloadFile downloads a file from MinIO
func (s *StorageService) DownloadFile(ctx context.Context, filename string) (io.ReadCloser, error) {
	object, err := s.minioClient.Client.GetObject(
		ctx,
		s.minioClient.BucketName,
		filename,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}

	return object, nil
}
