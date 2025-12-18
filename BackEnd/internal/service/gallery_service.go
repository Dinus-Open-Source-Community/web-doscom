package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	env "web_doscom/internal/config"
	"web_doscom/internal/database/model"

	"github.com/google/uuid"
)

type GalleryService struct {
	Model *model.GalleryModel
}

func NewGalleryService(m *model.GalleryModel) *GalleryService {
	return &GalleryService{Model: m}
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

// wrapper for insert gallery
func (m *GalleryService) InsertGallery(gallery *model.Gallery) (*model.Gallery, error) {
	return m.Model.InsertGallery(gallery)
}

// wrapper for get gallery by type
func (m *GalleryService) GetAllGalleryByType(tipe string, page, limit, offset int) ([]*model.GalleryResponse, int64, error) {

	var response []*model.GalleryResponse
	galleries, count, err := m.Model.GetGalleryByType(tipe, page, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	for _, data := range galleries {
		response = append(response, &model.GalleryResponse{
			ID:          data.ID,
			GalleryName: data.GalleryName,
			GalleryType: data.GalleryType,
			Description: data.Description,
			EventDate:   data.EventDate,
			FileSize:    data.FileSize,
			MimeType:    data.MimeType,
			AssetUrl:    data.AssetUrl,
		})
	}

	return response, count, nil
}

// wrapper for delete gallery
func (m *GalleryService) DeleteGallery(id int) error {
	return m.Model.DeleteGallery(id)
}

// check if file is valid or not
func isValidFile(file multipart.File) (bool, string) {
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return false, ""
	}

	fileType := http.DetectContentType(buffer)

	if _, err := file.Seek(0, 0); err != nil {
		return false, ""
	}
	allowedExt := map[string]bool{
		"image/jpg":       true,
		"image/jpeg":      true,
		"image/png":       true,
		"image/webp":      true,
		"video/mp4":       true,
		"video/quicktime": true,
	}

	return allowedExt[fileType], fileType

}

// generate random name file
func GenerateRandomName(originalName string) string {
	// get the extension of the file
	extension := strings.ToLower(filepath.Ext(originalName))

	// generate uuid for random name
	fileID := uuid.New()

	return fileID.String() + extension

}

// save uploaded file to storage
func SaveUploadedFile(fileHeader *multipart.FileHeader, savePath string) error {
	// make sure the path is valid or exist
	if err := os.MkdirAll(path.Dir(savePath), os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create folder")
	}

	// make temp folder
	tempFolder := savePath + ".tmp"

	// buka file
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// tulis dulu ke temp folder
	tempFile, err := os.Create(tempFolder)
	if err != nil {
		return fmt.Errorf("Failed to create file")
	}

	// copy dari temp ke folder asli
	if _, err := io.Copy(tempFile, file); err != nil {
		tempFile.Close()
		os.Remove(tempFolder)
		return fmt.Errorf("Failed to copy file")
	}

	defer tempFile.Close()

	return os.Rename(tempFolder, savePath)
}

// service for uploading image or video to storage and insert to database
func (m *GalleryService) UploadImage(files []*multipart.FileHeader) ([]*model.GalleryInsert, error) {
	env.LoadEnv()
	// get the storage path
	storagePath := os.Getenv("STORAGE_PATH")
	if storagePath == "" {
		storagePath = "storage/uploads"
	}

	// validasi -> cek total semua file dan file size, cek file type
	var (
		totalSize    int64
		validateData []validateFile
	)

	for _, fileHeader := range files {
		// sum total size file upload
		totalSize += fileHeader.Size

		// check size per file
		if fileHeader.Size > maxFileSize {
			return nil, fmt.Errorf("file size is too large, tf nigga")
		}

		file, err := fileHeader.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to parse multipart form: %s", fileHeader.Filename)
		}
		defer file.Close()

		//check if the file is valid and allowed or not
		isvalid, mimeType := isValidFile(file)
		if !isvalid {
			return nil, fmt.Errorf("file type not allowed, hayoo script apa ituuu")
		}

		kategori := strings.Split(mimeType, "/")[0]

		// save the file base on mimetype
		var folder string
		switch {
		case strings.HasPrefix(mimeType, "image"):
			folder = "image"

		case strings.HasPrefix(mimeType, "video"):
			folder = "video"
		default:
			return nil, fmt.Errorf("failed to save file")
		}

		validateData = append(validateData, validateFile{
			Fileheader: fileHeader,
			MimeType:   mimeType,
			Folder:     folder,
			fileSize:   fileHeader.Size,
			kategori:   kategori,
		})
	}

	// save all the file
	var uploadedFile []*model.GalleryInsert
	for _, file := range validateData {
		// make sure the storage folder exists
		dirPath := filepath.Join(storagePath, file.Folder)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return nil, fmt.Errorf("Failed to create storage folder")
		}

		// make random name file
		newFileName := GenerateRandomName(file.Fileheader.Filename)
		fmt.Println(newFileName)

		// save the fuckin file
		savePath := filepath.Join(dirPath, newFileName)
		if err := SaveUploadedFile(file.Fileheader, savePath); err != nil {
			return nil, fmt.Errorf("Failed to save file")
		}

		uploadedFile = append(uploadedFile, &model.GalleryInsert{
			AssetUrl:    savePath,
			GalleryName: file.Fileheader.Filename,
			Kategori:    file.kategori, // ini ambil dari mimeType dia video atau image
			FileSize:    file.fileSize,
			MimeType:    file.MimeType,
		})
	}

	return uploadedFile, nil
}

// service for uploading single (like me) image or video to storage
func (m *GalleryService) UploadSingleImage(files *multipart.FileHeader) (*model.GalleryInsert, error) {
	// get storage path
	env.LoadEnv()
	storagePath := os.Getenv("STORAGE_PATH")
	if storagePath == "" {
		storagePath = "storage/uploads"
	}

	if files.Size > maxFileSize {
		return nil, fmt.Errorf("file size is too large, tf nigga")
	}

	photo, err := files.Open()
	if err != nil {
		return nil, fmt.Errorf("Failed to read file")
	}
	defer photo.Close()

	// check if the file is valid and allowed or not
	isvalid, mimeType := isValidFile(photo)
	if !isvalid {
		return nil, fmt.Errorf("file type not allowed, hayoo script apa ituuu")
	}

	kategori := strings.Split(mimeType, "/")[0]

	// make random name file
	newFileName := GenerateRandomName(files.Filename)

	// save the file
	dirPath := filepath.Join(storagePath, kategori)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return nil, fmt.Errorf("Failed to create storage folder")
	}

	savePath := filepath.Join(dirPath, newFileName)
	if err := SaveUploadedFile(files, savePath); err != nil {
		return nil, fmt.Errorf("Failed to save file")
	}

	uploadedFile := &model.GalleryInsert{
		AssetUrl:    savePath,
		GalleryName: files.Filename,
		Kategori:    kategori,
		FileSize:    files.Size,
		MimeType:    mimeType,
	}

	return uploadedFile, nil
}

// func save photo to gallery dan insert to database -> multiple file
func (m *GalleryService) UploadAndInsertGallery(files []*multipart.FileHeader, gallery *model.CreateGallery) ([]*model.GalleryResponse, error) {
	// upload file first
	uploadedFile, err := m.UploadImage(files)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload file to storage")
	}

	// insert to database
	result := []*model.GalleryResponse{}
	for _, files := range uploadedFile {
		file_upload := &model.Gallery{
			GalleryName: files.GalleryName,
			GalleryType: gallery.GalleryType,
			Description: gallery.Description,
			EventDate:   gallery.EventDate,
			FileSize:    files.FileSize,
			MimeType:    files.MimeType,
			AssetUrl:    files.AssetUrl,
			Kategori:    files.Kategori,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		fileUpload, err := m.InsertGallery(file_upload)
		if err != nil {
			return nil, fmt.Errorf("Failed to insert data")
		}

		result = append(result, &model.GalleryResponse{
			ID:          fileUpload.ID,
			GalleryName: fileUpload.GalleryName,
			GalleryType: fileUpload.GalleryType,
			Description: fileUpload.Description,
			EventDate:   fileUpload.EventDate,
			FileSize:    fileUpload.FileSize,
			MimeType:    fileUpload.MimeType,
			AssetUrl:    fileUpload.AssetUrl,
		})
	}
	return result, nil
}

// func save profile photo to gallery and insert to database
func (m *GalleryService) UploadInsertSingleImage(files *multipart.FileHeader) (*model.GalleryResponse, error) {
	// upload file first
	uploadedFile, err := m.UploadSingleImage(files)
	if err != nil {
		return nil, fmt.Errorf("Failed to upload file to storage: %v", err)
	}

	// insert to database
	fileUpload := &model.Gallery{
		GalleryName: uploadedFile.GalleryName,
		GalleryType: "pengurus",
		Description: "foto Profile",
		EventDate:   time.Now().Format("2006-01-02"),
		FileSize:    uploadedFile.FileSize,
		MimeType:    uploadedFile.MimeType,
		AssetUrl:    uploadedFile.AssetUrl,
		Kategori:    uploadedFile.Kategori,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	upload, err := m.InsertGallery(fileUpload)
	if err != nil {
		return nil, fmt.Errorf("Failed to insert data: %v", err)
	}

	return &model.GalleryResponse{
		ID:          upload.ID,
		GalleryName: upload.GalleryName,
		GalleryType: upload.GalleryType,
		Description: upload.Description,
		EventDate:   upload.EventDate,
		FileSize:    upload.FileSize,
		MimeType:    upload.MimeType,
		AssetUrl:    upload.AssetUrl,
	}, nil
}
