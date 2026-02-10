package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"web_doscom/internal/config"
	"web_doscom/internal/database/model"
)

type StorageService struct {
	minioClient *config.MinioClient
	db          *gorm.DB
}

func NewStorageService(minioClient *config.MinioClient, db *gorm.DB) *StorageService {
	return &StorageService{
		minioClient: minioClient,
		db:          db,
	}
}

// AllowedImageExtensions defines allowed image file extensions
var AllowedImageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg"}

// MaxFileSize defines maximum file size (10MB)
const MaxFileSize = 10 * 1024 * 1024

// ValidateImageFile validates file extension, size, and actual content
func (s *StorageService) ValidateImageFile(fileHeader *multipart.FileHeader) error {
	// Check file size
	if fileHeader.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds maximum limit of 10MB")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	isValid := false
	for _, allowedExt := range AllowedImageExtensions {
		if ext == allowedExt {
			isValid = true
			break
		}
	}

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
func (s *StorageService) UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, folder string, userID uint) (string, error) {
	// Validate file
	if err := s.ValidateImageFile(fileHeader); err != nil {
		return "", err
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
		return "", fmt.Errorf("failed to upload file: %w", err)
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

	// Save metadata to database
	fileUpload := &model.FileUpload{
		UserID:           userID,
		Category:         folder,
		OriginalFilename: fileHeader.Filename,
		StoredFilename:   storedFilename,
		FileSize:         fileHeader.Size,
		ContentType:      contentType,
		FileURL:          fileURL,
	}

	if err := s.db.Create(fileUpload).Error; err != nil {
		// If database save fails, try to delete the uploaded file
		s.minioClient.Client.RemoveObject(ctx, s.minioClient.BucketName, storedFilename, minio.RemoveObjectOptions{})
		return "", fmt.Errorf("failed to save file metadata: %w", err)
	}

	return fileURL, nil
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
