package seeder

import (
	"context"
	"log"
	"time"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/utils"

	"gorm.io/gorm"
)

func SeedPengurus(db *gorm.DB) {

	pengurusModel := entity.PengurusModel{DB: db}

	now := time.Now()

	pengurusList := []entity.Pengurus{
		{
			UserID:    2,
			PhotoURL:  "",
			Email:     "jaringan@doscom.org",
			Divisi:    "jaringan",
			Name:      "muhammad daniyal haq",
			Position:  "KoorJaringan",
			Period:    "2023/2024",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			UserID:    3,
			PhotoURL:  "",
			Email:     "anggotajaringan@doscom.org",
			Divisi:    "jaringan",
			Name:      "fujikawa shinici",
			Position:  "JaringanAng",
			Period:    "2023/2024",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			UserID:    4,
			PhotoURL:  "",
			Email:     "medcrev@doscom.org",
			Divisi:    "medcrev",
			Name:      "dhion oppa",
			Position:  "KoorMedcrev",
			Period:    "2023/2024",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			UserID:    5,
			PhotoURL:  "",
			Email:     "anggotamedcrev@doscom.org",
			Divisi:    "medcrev",
			Name:      "dhion tapi keren",
			Position:  "MedcrevAng",
			Period:    "2023/2024",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	for i := range pengurusList {
		if err := pengurusModel.InsertPengurus(&pengurusList[i]); err != nil {
			log.Printf("Failed to insert data pengurus %v", err)
			return
		}
	}

	log.Println("Seed table 'pengurus' completed.")
}

func SeedPengurusSosmed(db *gorm.DB, pengurusID int, sosmedUrl []string) {
	ctx := context.Background()
	pengurusSosmedModel := entity.PengurusSosmedModel{DB: db}
	// parse url info
	socialMediaInfo, err := utils.ExtractSocialMediaBatch(sosmedUrl)
	if err != nil {
		log.Printf("Failed to extract social media info %v", err)
		return
	}

	sosmedPayload := make([]dto.CreatePengurusSosmedPayload, len(socialMediaInfo))
	for i, info := range socialMediaInfo {
		sosmedPayload[i] = dto.CreatePengurusSosmedPayload{
			PengurusID: pengurusID,
			Platform:   info.Platform,
			Username:   info.Username,
			Url:        info.URL,
			IsPrimary:  i == 0, // true hanya untuk index 0
		}
	}

	_, err = pengurusSosmedModel.InsertPengurusSosmed(ctx, sosmedPayload)
	if err != nil {
		log.Printf("Failed to insert data pengurus sosmed %v", err)
		return
	}

	log.Println("Seed table 'pengurus' completed.")
}

func RunSeedPengurus(db *gorm.DB) {
	SeedPengurus(db)

	pengurusID := []int{1, 2, 3, 4}
	sosmedURL := []string{
		"https://www.instagram.com/dhioonnn/",
		"https://www.linkedin.com/in/dhion-nur-damanhuri-2bb863275/",
		"https://github.com/IKOPOO",
	}
	for i := range pengurusID {
		SeedPengurusSosmed(db, pengurusID[i], sosmedURL)
	}
}
