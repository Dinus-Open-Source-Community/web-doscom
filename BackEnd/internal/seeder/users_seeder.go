package seeder

import (
	"log"
	"time"

	"web_doscom/internal/auth"
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	passwordHash := auth.HashPassword("password123")

	users := []model.User{
		{
			Username:  "superadmin",
			Email:     "superadmin@doscom.org",
			Full_name: "Super Admin Doscom",
			Password:  passwordHash,
			Role:      "SuperAdmin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Username:  "koorpemro",
			Email:     "pemro@doscom.org",
			Full_name: "Koor Pemrograman",
			Password:  passwordHash,
			Role:      "KoorPemro",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Username:  "anggotapemro",
			Email:     "anggotapemro@doscom.org",
			Full_name: "Anggota Pemrograman",
			Password:  passwordHash,
			Role:      "pemroAnggota",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for _, user := range users {
		db.FirstOrCreate(&user, model.User{Email: user.Email})
	}

	log.Println("Seed table 'users' completed.")
}
