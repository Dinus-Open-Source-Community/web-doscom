package seeder

import (
	"log"
	"time"

	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedPengurus(db *gorm.DB) {
	// Assumes users with ID 1, 2, 3 have been seeded.
	pengurusList := []model.Pengurus{
		{
			UserID:    1,
			PhotoURL:  "https://dummyimage.com/200x200/000/fff&text=SuperAdmin",
			Email:     "superadmin@doscom.org",
			Divisi:    "bph",
			Name:      "Super Admin Doscom",
			Position:  "ketum",
			Sosmed:    "instagram",
			Period:    "2023/2024",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    2,
			PhotoURL:  "https://dummyimage.com/200x200/000/fff&text=KoorPemro",
			Email:     "pemro@doscom.org",
			Divisi:    "pemro",
			Name:      "Koor Pemrograman",
			Position:  "KoorPemro",
			Sosmed:    "github",
			Period:    "2023/2024",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			UserID:    3,
			PhotoURL:  "https://dummyimage.com/200x200/000/fff&text=AnggotaPemro",
			Email:     "anggotapemro@doscom.org",
			Divisi:    "pemro",
			Name:      "Anggota Pemrograman",
			Position:  "PemroAng",
			Sosmed:    "linkedin",
			Period:    "2023/2024",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, p := range pengurusList {
		db.FirstOrCreate(&p, model.Pengurus{Email: p.Email})
	}

	log.Println("Seed table 'pengurus' completed.")
}
