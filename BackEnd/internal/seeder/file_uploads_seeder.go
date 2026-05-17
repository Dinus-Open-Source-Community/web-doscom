package seeder

// import (
// 	"fmt"
// 	"log"
// 	"mime"
// 	"os"
// 	"path/filepath"
// 	"time"
// 	"web_doscom/internal/config"
// 	"web_doscom/internal/database/model/dto"
// 	"web_doscom/internal/database/model/entity"
// 	"web_doscom/internal/service"
//
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// func OpenLocalFile(pathFile string) (*os.File, string, int64, string, error) {
// 	file, err := os.Open(pathFile)
// 	if err != nil {
// 		return nil, "", 0, "", err
// 	}
//
// 	stat, err := file.Stat()
// 	if err != nil {
// 		return nil, "", 0, "", err
// 	}
//
// 	fileName := filepath.Base(pathFile)
// 	contentType := mime.TypeByExtension(filepath.Ext(fileName))
// 	if contentType == "" {
// 		contentType = "application/octet-stream"
// 	}
//
// 	return file, fileName, stat.Size(), contentType, nil
// }
//
// func UploadFileToStorage(minioClient *config.MinioClient, pathFile []string, folder string) (entity.FileUpload, error) {
// 	// upload file to minio
//
// 	fileUpload := make([]entity.FileUpload, len(pathFile))
// 	for i, path := range pathFile {
// 		file, fileName, fileSize, contentType, err := OpenLocalFile(path)
// 		if err != nil {
// 			return entity.FileUpload{}, err
// 		}
//
// 		fileUpload, err :=
//
// 		storedFilename := fmt.Sprintf("%s/%s-%s%s", folder, time.Now().Format("20060102"), uuid.New().String(), fileName)
// 		fileUpload[i] = entity.FileUpload{
// 			Category:         folder,
// 			OriginalFilename: fileName,
// 			StoredFilename:   storedFilename,
// 			FileSize:         fileSize,
// 			ContentType:      contentType,
// 			FileURL:          "https://dummyimage.com/800x600/000/fff&text=Gallery1",
// 			UploadedAt:       time.Now(),
// 			UpdatedAt:        time.Now(),
// 		}
// 	}
//
// }
//
// func SeedFileUploads(db *gorm.DB) (dto.FileUploadResponse, error) {
//
// 	// ambil file dari folder test asset
//
// 	// upload ke storage minio
// 	// save metadata to database
// 	files := []entity.FileUpload{
// 		{
// 			UserID:           1,
// 			Category:         "gallery",
// 			OriginalFilename: "dummy_gallery.jpg",
// 			StoredFilename:   "12345_dummy_gallery.jpg",
// 			FileSize:         102400,
// 			ContentType:      "image/jpeg",
// 			FileURL:          "https://dummyimage.com/800x600/000/fff&text=Gallery1",
// 			UploadedAt:       time.Now(),
// 			UpdatedAt:        time.Now(),
// 		},
// 		{
// 			UserID:           2,
// 			Category:         "blog",
// 			OriginalFilename: "dummy_blog_thumb.png",
// 			StoredFilename:   "67890_dummy_blog_thumb.png",
// 			FileSize:         204800,
// 			ContentType:      "image/png",
// 			FileURL:          "https://dummyimage.com/800x600/000/fff&text=BlogThumb1",
// 			UploadedAt:       time.Now(),
// 			UpdatedAt:        time.Now(),
// 		},
// 	}
//
// 	for _, f := range files {
// 		db.FirstOrCreate(&f, entity.FileUpload{StoredFilename: f.StoredFilename})
// 	}
//
// 	log.Println("Seed table 'file_uploads' completed.")
// }
