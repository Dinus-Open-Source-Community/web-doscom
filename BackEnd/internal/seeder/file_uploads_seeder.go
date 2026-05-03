package seeder

import (
	"log"
	"time"

	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedFileUploads(db *gorm.DB) {
	files := []model.FileUpload{
		{
			UserID:           1,
			Category:         "gallery",
			OriginalFilename: "dummy_gallery.jpg",
			StoredFilename:   "12345_dummy_gallery.jpg",
			FileSize:         102400,
			ContentType:      "image/jpeg",
			FileURL:          "https://dummyimage.com/800x600/000/fff&text=Gallery1",
			UploadedAt:       time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			UserID:           2,
			Category:         "blog",
			OriginalFilename: "dummy_blog_thumb.png",
			StoredFilename:   "67890_dummy_blog_thumb.png",
			FileSize:         204800,
			ContentType:      "image/png",
			FileURL:          "https://dummyimage.com/800x600/000/fff&text=BlogThumb1",
			UploadedAt:       time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	for _, f := range files {
		db.FirstOrCreate(&f, model.FileUpload{StoredFilename: f.StoredFilename})
	}

	log.Println("Seed table 'file_uploads' completed.")
}
