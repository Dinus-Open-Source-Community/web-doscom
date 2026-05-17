package seeder

import (
	"log"
	"time"

	"web_doscom/internal/auth"
	"web_doscom/internal/database/model/entity"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	now := time.Now()

	passwordHash := auth.HashPassword("password123")

	users := []entity.User{
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
			Username:  "koor jaringan",
			Email:     "jaringan@doscom.org",
			Full_name: "muhammad daniyal haq",
			Password:  passwordHash,
			Role:      "KoorJaringan",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Username:  "anggota jaringan",
			Email:     "anggotajaringan@doscom.org",
			Full_name: "fujikawa nippon",
			Password:  passwordHash,
			Role:      "jaringanAnggota",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Username:  "dhion sangar boss",
			Email:     "medcrev@doscom.org",
			Full_name: "dhion oppa",
			Password:  passwordHash,
			Role:      "KoorMedcrev",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			Username:  "dhion tapi anggota",
			Email:     "anggotamedcrev@doscom.org",
			Full_name: "dhion keren",
			Password:  passwordHash,
			Role:      "medcrevAnggota",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	for _, user := range users {
		db.Where(entity.User{Email: user.Email}).
			Or(entity.User{Username: user.Username}).
			FirstOrCreate(&user)
	}

	log.Println("Seed table 'users' completed.")
}
