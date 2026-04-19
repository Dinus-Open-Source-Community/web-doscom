package seeder

import (
	"log"
	"time"

	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedPengurus(db *gorm.DB) {
	now := time.Now()

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
			CreatedAt: now,
			UpdatedAt: now,
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
			CreatedAt: now,
			UpdatedAt: now,
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
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	for _, p := range pengurusList {
		var existing model.Pengurus

		err := db.Where("email = ?", p.Email).
			Or("id_user = ?", p.UserID).
			First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			db.Create(&p)
		}
	}

	log.Println("Seed table 'pengurus' completed.")
}
