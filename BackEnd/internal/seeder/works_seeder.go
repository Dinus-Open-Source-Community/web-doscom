package seeder

import (
	"context"
	"log"
	"time"
	"web_doscom/internal/database/model/entity"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedWorks(db *gorm.DB) {
	now := time.Now()

	workModel := entity.WorkModel{DB: db}
	ctx := context.Background()

	workList := []entity.Work{
		{
			PengurusID:   2,
			Title:        "Website Profil Doscom",
			Tagline:      "Bikin website tanpa rancangan awal, go with the flow",
			Description:  "Pengembangan website internal profil dinus open source community.",
			Slug:         "website-profil-doscom",
			ProjectType:  "website",
			Technologies: pq.StringArray{"golang", "postgresql", "astro"},
			ProjectDate:  time.Date(2025, time.December, 7, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   2,
			Title:        "Website Deteksi Femboy",
			Tagline:      "Deteksi femboy tanpa ribet",
			Description:  "bagi anda yang ingin mengetahui teman atau anda sendiri yang femboy, maka website ini cocok karena bisa deteksi femboy dengan mudah",
			Slug:         "website-femboy-deteksi",
			ProjectType:  "website",
			Technologies: pq.StringArray{"html", "css", "javascript", "mysql"},
			ProjectDate:  time.Date(2026, time.January, 17, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   3,
			Title:        "Ricesource",
			Tagline:      "Opensource ur ricing",
			Description:  "platform to share ur linux configuration and using other people configuration on ur machine",
			Slug:         "wong sangar rak perlu rapat, seng penting garap jadi",
			ProjectType:  "web",
			Technologies: pq.StringArray{"express.js", "next.js", "postgresql"},
			ProjectDate:  time.Date(2024, time.December, 20, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "draft",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	for i := range workList {
		_, err := workModel.InsertWork(ctx, &workList[i])
		if err != nil {
			log.Printf("Failed to insert data work %v", err)
			return
		}
	}
	log.Println("Seed table 'work' completed.")
}

func RunSeedWorks(db *gorm.DB) {
	// for now, only run works seeder -> because im too lazy to create a seeder for work_gallery
	SeedWorks(db)
	// call seeder for work_gallery -> if it exist
}
