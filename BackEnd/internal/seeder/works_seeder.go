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

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func uploadWorkPhotos(
	ctx context.Context,
	galleryService *service.GalleryService,
	folderNumber int,
	uploaderUserID int,
	workTitle string,
) ([]int, string, error) {
	photoDirectory := filepath.Join(
		"storage",
		"uploads",
		"work",
		fmt.Sprintf("%d", folderNumber),
	)

	entries, err := os.ReadDir(photoDirectory)
	if err != nil {
		return nil, "", fmt.Errorf(
			"failed to read work photo directory %d: %w",
			folderNumber,
			err,
		)
	}

	var fileUploads []*dto.UploadFileRequest
	var cleanups []func()

	defer func() {
		for _, cleanup := range cleanups {
			cleanup()
		}
	}()

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		photoPath := filepath.Join(photoDirectory, entry.Name())

		header, file, cleanup, err := utils.OpenSeedImage(photoPath)
		if err != nil {
			return nil, "", fmt.Errorf(
				"failed to open work photo %q: %w",
				photoPath,
				err,
			)
		}

		cleanups = append(cleanups, cleanup)
		fileUploads = append(fileUploads, &dto.UploadFileRequest{
			FileHeader: header,
			File:       file,
			Folder:     "work",
			UserID:     uint(uploaderUserID),
		})
	}

	if len(fileUploads) == 0 {
		return nil, "", fmt.Errorf(
			"no images found in directory %s",
			photoDirectory,
		)
	}

	galleryData := &dto.GalleryInsert{
		IDUsers:     uploaderUserID,
		GalleryName: "Foto work " + workTitle,
		GalleryType: "work",
		Description: "Dokumentasi untuk work " + workTitle,
		EventDate:   time.Now(),
	}

	galleries, err := galleryService.UploadAndInsertGalleryMultiple(
		ctx,
		galleryData,
		fileUploads,
	)
	if err != nil {
		return nil, "", fmt.Errorf(
			"failed to upload work galleries: %w",
			err,
		)
	}

	if len(galleries) == 0 {
		return nil, "", fmt.Errorf(
			"no gallery created for work %q",
			workTitle,
		)
	}

	galleryIDs := make([]int, len(galleries))
	for i, gallery := range galleries {
		galleryIDs[i] = gallery.ID
	}

	// Gambar pertama dalam folder menjadi thumbnail.
	return galleryIDs, galleries[0].FileURL, nil
}

func SeedWorksGallery(
	db *gorm.DB,
	workGalleries []*entity.WorkGallery,
) error {
	model := entity.WorkGalleryModel{DB: db}

	if err := model.InsertWorkGalleryMultiple(
		context.Background(),
		workGalleries,
	); err != nil {
		return fmt.Errorf("failed to insert work galleries: %w", err)
	}

	return nil
}

func SeedWorks(db *gorm.DB, galleryService *service.GalleryService) error {
	now := time.Now()

	workModel := entity.WorkModel{DB: db}
	ctx := context.Background()

	workList := []entity.Work{
		{
			PengurusID:   12,
			Title:        "DOSCOM University 2025 Landing Page",
			Tagline:      "Gerbang pendaftaran bagi calon talenta open source",
			Description:  "Landing page interaktif untuk manajemen pendaftaran peserta workshop tahunan DOSCOM University.",
			Slug:         "doscom-university-2025-landing-page",
			ProjectType:  "website",
			Technologies: pq.StringArray{"react", "golang", "tripay"},
			ProjectDate:  time.Date(2025, time.February, 10, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   29,
			Title:        "WEB DOSCOM Company Profile",
			Tagline:      "Manage your company profile easily and quickly",
			Description:  "Website for managing company profile DOSCOM University 2025 Landing Page",
			Slug:         "web-doscom-company-profile",
			ProjectType:  "website",
			Technologies: pq.StringArray{"golang", "astro"},
			ProjectDate:  time.Date(2026, time.January, 5, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   45,
			Title:        "DFORM - Website Form Builder",
			Tagline:      "Build and deploy forms easily",
			Description:  "website form builder for creating and deploying forms for your event",
			Slug:         "dform-website-form-builder",
			ProjectType:  "website",
			Technologies: pq.StringArray{"laravel", "mariadb", "midtrans-api"},
			ProjectDate:  time.Date(2026, time.January, 22, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   3,
			Title:        "TEALINUX OS",
			Tagline:      "System operating for desktop",
			Description:  "Operating system for desktop computers that is based on the Linux kernel and lightweight",
			Slug:         "tealinux-os",
			ProjectType:  "operating-system",
			Technologies: pq.StringArray{"rust", "grpc", "prometheus"},
			ProjectDate:  time.Date(2026, time.March, 14, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   3,
			Title:        "DAEMON COMMAND & CONTROL",
			Tagline:      "Cross-platform Command & Control framework built with Rust",
			Description:  "DCC (Daemon Command and Control) is a full-stack C2 framework designed for penetration testing and red team operations. It features: A Rust-based server that manages connected agents (implants) with AES-256 encrypted communications Cross-platform clients (Linux and Windows) that beacon back to the server for command execution A React + TypeScript dashboard (Vite + TailwindCSS) providing real-time telemetry, agent management, payload generation, reverse shell access, secure notes, and system logging An automated payload generator that cross-compiles implants for multiple OS targets",
			Slug:         "daemon-command-control",
			ProjectType:  "Framework",
			Technologies: pq.StringArray{"rust", "grpc", "prometheus"},
			ProjectDate:  time.Date(2026, time.March, 14, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	const uploaderUserID = 1

	for i := range workList {
		work := &workList[i]
		folderNumber := i + 1

		imageIDs, imageURL, err := uploadWorkPhotos(
			ctx,
			galleryService,
			folderNumber,
			uploaderUserID,
			work.Title,
		)
		if err != nil {
			return fmt.Errorf(
				"failed to process photos for work %q: %w",
				work.Title,
				err,
			)
		}

		work.ImageURL = imageURL

		if _, err := workModel.InsertWork(ctx, work); err != nil {
			return fmt.Errorf(
				"failed to insert work %q: %w",
				work.Title,
				err,
			)
		}

		workGalleryData := make(
			[]*entity.WorkGallery,
			len(imageIDs),
		)

		for galleryIndex, galleryID := range imageIDs {
			workGalleryData[galleryIndex] = &entity.WorkGallery{
				IDWork:    work.ID,
				IDGallery: galleryID,
			}
		}

		if err := SeedWorksGallery(db, workGalleryData); err != nil {
			return fmt.Errorf(
				"failed to insert gallery for work %q: %w",
				work.Title,
				err,
			)
		}
	}
	log.Println("Seed table 'work' completed.")

	log.Println("Seed table 'work_gallery' completed.")

	return nil
}

func RunSeedWorks(db *gorm.DB, minioClient *config.MinioClient) {
	models := database.NewModel(db)
	storageService := service.NewStorageService(
		minioClient,
		&models.FileUploads,
	)

	galleryService := service.NewGalleryService(
		&models.Gallery,
		storageService,
	)
	if err := SeedWorks(db, galleryService); err != nil {
		log.Fatalf("Failed to seed works: %v", err)
		return
	}
}
