package seeder

import (
	"log"
	"time"

	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedGallery(db *gorm.DB) {
	galleryList := []model.Gallery{
		{
			IDUsers:      1,
			FileUploadID: 1,
			GalleryName:  "Gathering Doscom 2024",
			GalleryType:  "event",
			Description:  "Foto kegiatan gathering tahun 2024 bersama anggota baru.",
			EventDate:    time.Now().AddDate(0, -1, 0), // 1 bulan lalu
			FileURL:      "https://dummyimage.com/800x600/000/fff&text=Gathering2024",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			IDUsers:      2,
			FileUploadID: 1,
			GalleryName:  "Workshop Go",
			GalleryType:  "activity",
			Description:  "Foto kegiatan workshop pemrograman Go.",
			EventDate:    time.Now().AddDate(0, -2, 0), // 2 bulan lalu
			FileURL:      "https://dummyimage.com/800x600/000/fff&text=WorkshopGo",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	for _, g := range galleryList {
		db.FirstOrCreate(&g, model.Gallery{GalleryName: g.GalleryName})
	}

	log.Println("Seed table 'gallery' completed.")
}
