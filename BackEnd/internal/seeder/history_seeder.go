package seeder

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/service"
	"web_doscom/internal/utils"

	"gorm.io/gorm"
)

type historySeedData struct {
	History     entity.HistoryTimeline
	PhotoFolder string
}

func SeedHistory(
	db *gorm.DB,
	galleryService *service.GalleryService,
) error {
	ctx := context.Background()

	historyModel := entity.HistoryTimelineModel{DB: db}
	historyPhotosModel := entity.HistoryPhotosModel{DB: db}

	const (
		authorID   = 1
		totalPhoto = 6
	)

	historyList := []historySeedData{
		{
			History: entity.HistoryTimeline{
				IDAuthor:     authorID,
				Title:        "Open Recruitment dan Pembaruan Generasi",
				Year:         "2024",
				Description:  "Menjaga keberlanjutan organisasi melalui regenerasi berkala.",
				DisplayOrder: 1,
			},
			PhotoFolder: "oprec",
		},
		{
			History: entity.HistoryTimeline{
				IDAuthor:     authorID,
				Title:        "Edukasi Publik dan Penyebaran Manfaat",
				Year:         "2025",
				Description:  "Membagikan ilmu teknologi lewat workshop interaktif.",
				DisplayOrder: 2,
			},
			PhotoFolder: "2025",
		},
		{
			History: entity.HistoryTimeline{
				IDAuthor:     authorID,
				Title:        "Harmonisasi Internal dan Ruang Dengar",
				Year:         "2026",
				Description:  "Mempererat rasa kekeluargaan dan solidaritas pengurus.",
				DisplayOrder: 3,
			},
			PhotoFolder: "2024",
		},
	}

	for i := range historyList {
		seed := &historyList[i]
		history := &seed.History

		photoDirectory := filepath.Join(
			"storage",
			"uploads",
			"history",
			seed.PhotoFolder,
		)

		entries, err := os.ReadDir(photoDirectory)
		if err != nil {
			return fmt.Errorf(
				"failed to read history folder %q: %w",
				seed.PhotoFolder,
				err,
			)
		}

		fileUploads := make([]*dto.UploadFileRequest, 0, len(entries))
		cleanups := make([]func(), 0, len(entries))

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			imagePath := filepath.Join(
				photoDirectory,
				entry.Name(),
			)

			header, file, cleanup, err := utils.OpenSeedImage(imagePath)
			if err != nil {
				for _, cleanup := range cleanups {
					cleanup()
				}

				return fmt.Errorf(
					"failed to open history image %q: %w",
					imagePath,
					err,
				)
			}

			cleanups = append(cleanups, cleanup)
			fileUploads = append(fileUploads, &dto.UploadFileRequest{
				FileHeader: header,
				File:       file,
				Folder:     "history",
				UserID:     uint(history.IDAuthor),
			})
		}

		if len(fileUploads) == 0 {
			return fmt.Errorf(
				"no images found in history folder %q",
				seed.PhotoFolder,
			)
		}

		galleryPayload := &dto.GalleryInsert{
			IDUsers:     history.IDAuthor,
			GalleryName: "Dokumentasi " + history.Title,
			GalleryType: "history",
			Description: history.Description,
			EventDate:   time.Now(),
		}

		galleries, err := galleryService.UploadAndInsertGalleryMultiple(
			ctx,
			galleryPayload,
			fileUploads,
		)

		for _, cleanup := range cleanups {
			cleanup()
		}

		if err != nil {
			return fmt.Errorf(
				"failed to upload gallery history %q: %w",
				history.Title,
				err,
			)
		}

		if len(galleries) == 0 {
			return fmt.Errorf(
				"no gallery created for history %q",
				history.Title,
			)
		}

		if err := historyModel.InsertHistoryTimeline(ctx, history); err != nil {
			return fmt.Errorf(
				"failed to insert history %q: %w",
				history.Title,
				err,
			)
		}

		for _, gallery := range galleries {
			photo := &entity.HistoryPhotos{
				IDHistory: history.ID,
				ImagerURL: gallery.FileURL,
			}

			if err := historyPhotosModel.InsertHistoryPhotos(
				ctx,
				photo,
			); err != nil {
				return fmt.Errorf(
					"failed to insert photo history %q: %w",
					history.Title,
					err,
				)
			}
		}
	}

	return nil
}

func RunSeedHistory(
	db *gorm.DB,
	minioClient *config.MinioClient,
) {
	models := database.NewModel(db)

	storageService := service.NewStorageService(
		minioClient,
		&models.FileUploads,
	)

	galleryService := service.NewGalleryService(
		&models.Gallery,
		storageService,
	)

	if err := SeedHistory(db, galleryService); err != nil {
		log.Fatalf("Failed to seed history: %v", err)
	}

	log.Println("Seed table 'history_timeline' completed.")
	log.Println("Seed table 'history_photos' completed.")
}
