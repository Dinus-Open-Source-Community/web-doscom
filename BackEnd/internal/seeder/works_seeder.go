package seeder

import (
	"log"
	"time"

	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedWorks(db *gorm.DB) {
	workList := []model.Work{
		{
			Title:       "Website Profil Doscom",
			GalleryID:   1, // Asumsi ID gallery 1 sudah ada
			Description: "Pengembangan website internal profil dinus open source community.",
			ProjectDate: time.Now().AddDate(-1, 0, 0), // 1 tahun lalu
			TeamProject: 2, // ID koorpemro
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, w := range workList {
		db.FirstOrCreate(&w, model.Work{Title: w.Title})
	}

	log.Println("Seed table 'work' completed.")
}
