package seeder

import (
	"log"
	"time"

	"web_doscom/internal/auth"
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	now := time.Now()

	passwordHash := auth.HashPassword("password123")

	users := []model.User{
		{
			Username:  "superadmin",
			Email:     "superadmin@doscom.org",
			Full_name: "Super Admin Doscom",
			Password:  passwordHash,
			Role:      "SuperAdmin",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Username:  "koorpemro",
			Email:     "pemro@doscom.org",
			Full_name: "Koor Pemrograman",
			Password:  passwordHash,
			Role:      "KoorPemro",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Username:  "anggotapemro",
			Email:     "anggotapemro@doscom.org",
			Full_name: "Anggota Pemrograman",
			Password:  passwordHash,
			Role:      "pemroAnggota",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	for _, user := range users {
		db.Where(model.User{Email: user.Email}).
			Or(model.User{Username: user.Username}).
			FirstOrCreate(&user)
	}

	log.Println("Seed table 'users' completed.")
}
