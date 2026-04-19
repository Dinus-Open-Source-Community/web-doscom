package seeder

import (
	"log"
	"time"

	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedWorks(db *gorm.DB) {
	now := time.Now()

	workList := []model.Work{
		{
			PengurusID:   2,
			Title:        "Website Profil Doscom",
			Tagline:      "Company Profile Website",
			Description:  "Pengembangan website internal profil dinus open source community.",
			Slug:         "website-profil-doscom",
			ProjectType:  "web",
			Technologies: []string{"golang", "gin", "postgresql"},
			ProjectDate:  now.AddDate(-1, 0, 0),
			ImageURL:     "https://dummyimage.com/800x600/000/fff&text=Doscom",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	for _, w := range workList {
		db.FirstOrCreate(&w, model.Work{Slug: w.Slug})
	}

	log.Println("Seed table 'work' completed.")
}
