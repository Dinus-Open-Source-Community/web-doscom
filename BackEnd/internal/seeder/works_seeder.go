package seeder

import (
	"context"
	"fmt"
	"log"
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

func uploadFilePhoto(ctx context.Context, galleryService *service.GalleryService) ([]int, string, error) {
	var fileUploads []*dto.UploadFileRequest
	var cleanups []func()

	for imageNumber := 1; imageNumber <= 5; imageNumber++ {
		photoPath := filepath.Join("storage", "uploads", "work", fmt.Sprintf("%d.jpg", imageNumber))

		header, file, cleanup, err := utils.OpenSeedImage(photoPath)
		if err != nil {
			return nil, "", fmt.Errorf("failed to open photo %d: %w", imageNumber, err)
		}
		cleanups = append(cleanups, cleanup)
		fileUploads = append(fileUploads, &dto.UploadFileRequest{
			FileHeader: header,
			File:       file,
			Folder:     "work",
			UserID:     uint(1),
		})
	}

	defer func() {
		for _, cleanup := range cleanups {
			cleanup()
		}
	}()

	galleryData := &dto.GalleryInsert{
		IDUsers:     1,
		GalleryName: "foto untuk works",
		GalleryType: "work",
		Description: "foto ini untuk kepentingan works bwang",
		EventDate:   time.Date(2025, time.January, 1, 0, 0, 0, 0, time.Local),
	}
	galleries, err := galleryService.UploadAndInsertGalleryMultiple(ctx, galleryData, fileUploads)
	if err != nil {
		return nil, "", fmt.Errorf("failed to upload new image and gallery %w", err)
	}

	imageID := make([]int, len(galleries))
	for i, id := range galleries {
		imageID[i] = id.ID
	}
	return imageID, galleries[0].FileURL, nil
}

func SeedWorksGallery(db *gorm.DB, workGallery []*entity.WorkGallery) error {
	workModel := entity.WorkGalleryModel{DB: db}
	ctx := context.Background()

	if err := workModel.InsertWorkGalleryMultiple(ctx, workGallery); err != nil {
		return fmt.Errorf("failed to insert data %w", err)
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
			Technologies: pq.StringArray{"next.js", "tailwindcss", "typescript"},
			ProjectDate:  time.Date(2025, time.February, 10, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   29,
			Title:        "Git-Lazy TUI Tool",
			Tagline:      "Automate your Git workflow smoothly right from the terminal",
			Description:  "Terminal User Interface (TUI) berbasis Rust untuk mempermudah dan mempercepat pengelolaan staging, commit, dan push repository Git.",
			Slug:         "git-lazy-tui-tool",
			ProjectType:  "cli-desktop",
			Technologies: pq.StringArray{"rust", "git-api"},
			ProjectDate:  time.Date(2026, time.January, 5, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   45,
			Title:        "Savorbite Backend API",
			Tagline:      "Scalable backend architecture with complex ERD integration",
			Description:  "Pengembangan modul server inti untuk aplikasi e-commerce makanan, lengkap dengan enkripsi data dan payment gateway logis.",
			Slug:         "savorbite-backend-api",
			ProjectType:  "backend-service",
			Technologies: pq.StringArray{"golang", "postgresql", "midtrans-api"},
			ProjectDate:  time.Date(2026, time.January, 22, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   3,
			Title:        "Microservices Latency Monitor Dashboard",
			Tagline:      "Low-latency real-time tracking for complex systems",
			Description:  "Sistem monitoring berbasis Terminal User Interface untuk memantau health-check dan log antrean dari arsitektur microservices secara real-time.",
			Slug:         "microservices-latency-monitor-dashboard",
			ProjectType:  "terminal-app",
			Technologies: pq.StringArray{"rust", "grpc", "prometheus"},
			ProjectDate:  time.Date(2026, time.March, 14, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "draft",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   54,
			Title:        "Aplikasi Otomatisasi Absensi DOSCOM",
			Tagline:      "Scan, record, and close the sheet instantly",
			Description:  "Sistem pencatatan kehadiran otomatis berbasis QR Code untuk mempermudah rekap absensi rapat mingguan pengurus.",
			Slug:         "aplikasi-otomatisasi-absensi-doscom",
			ProjectType:  "web",
			Technologies: pq.StringArray{"node.js", "express.js", "mongodb"},
			ProjectDate:  time.Date(2025, time.November, 12, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   21,
			Title:        "Linux Custom ISO Builder UI",
			Tagline:      "Ricing your distribution right from a modular panel",
			Description:  "Sebuah utilitas berbasis web untuk melakukan kustomisasi paket dasar distro Linux sebelum membuat berkas berkstensi ISO.",
			Slug:         "linux-custom-iso-builder-ui",
			ProjectType:  "website",
			Technologies: pq.StringArray{"react.js", "python", "bash"},
			ProjectDate:  time.Date(2025, time.August, 19, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   8,
			Title:        "NvChad Shared Configurations Vault",
			Tagline:      "Blazing fast Neovim config setup tailored for Backend Engineers",
			Description:  "Kumpulan kurasi berkas konfigurasi NvChad yang dioptimalkan penuh untuk kenyamanan penulisan kode berskala besar di Golang dan Rust.",
			Slug:         "nvchad-shared-configurations-vault",
			ProjectType:  "open-source-repo",
			Technologies: pq.StringArray{"lua", "neovim", "bash"},
			ProjectDate:  time.Date(2026, time.April, 2, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   37,
			Title:        "Dinus Open Source Competition Platform",
			Tagline:      "Uji ketangkasan ngoding tanpa gangguan server down",
			Description:  "Platform terintegrasi untuk mengadakan kompetisi IT tingkat regional, lengkap dengan auto-grader script.",
			Slug:         "dinus-open-source-competition-platform",
			ProjectType:  "web",
			Technologies: pq.StringArray{"golang", "mysql", "docker"},
			ProjectDate:  time.Date(2025, time.October, 5, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   14,
			Title:        "Technopreneurship Project Marketplace",
			Tagline:      "Ubah ide tugas kuliah menjadi startup bernilai jual",
			Description:  "Sistem pameran proyek digital mahasiswa yang dirancang untuk memenuhi penilaian inkubasi bisnis technopreneurship.",
			Slug:         "technopreneurship-project-marketplace",
			ProjectType:  "website",
			Technologies: pq.StringArray{"laravel", "javascript", "bootstrap"},
			ProjectDate:  time.Date(2025, time.December, 18, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   60,
			Title:        "Intelligent Camera Data Acquisition System",
			Tagline:      "Real-time surveillance environment environment setup",
			Description:  "Pipeline otomatisasi akuisisi data umpan kamera cctv untuk riset pengenalan objek dan pemantauan kluster jaringan.",
			Slug:         "intelligent-camera-data-acquisition-system",
			ProjectType:  "system-integration",
			Technologies: pq.StringArray{"python", "opencv", "mqtt"},
			ProjectDate:  time.Date(2025, time.December, 29, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   42,
			Title:        "Shortlink DOSCOM Generator",
			Tagline:      "Bikin link pendek dengan domain kustom komunitas",
			Description:  "Layanan mandiri internal pengurus untuk mempersingkat tautan promosi kegiatan DOSCOM dengan analitik klik statis.",
			Slug:         "shortlink-doscom-generator",
			ProjectType:  "web-service",
			Technologies: pq.StringArray{"node.js", "redis", "postgresql"},
			ProjectDate:  time.Date(2025, time.June, 4, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   18,
			Title:        "LKS Web Dev Simulator",
			Tagline:      "Bahan ajar simulasi kompetisi untuk siswa SMK binaan",
			Description:  "Aplikasi modul latihan interaktif yang dikembangkan khusus untuk mempersiapkan kompetisi LKS siswa SMK dalam ekosistem full-stack.",
			Slug:         "lks-web-dev-simulator",
			ProjectType:  "web",
			Technologies: pq.StringArray{"laravel", "mysql", "jquery"},
			ProjectDate:  time.Date(2026, time.February, 11, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   33,
			Title:        "Discord Bot Management Utility",
			Tagline:      "Mengatur server DOSCOM biar tidak sepi chat",
			Description:  "Bot serbaguna untuk menyapa anggota baru, moderasi kata terlarang otomatis, dan pemutar musik di kanal suara.",
			Slug:         "discord-bot-management-utility",
			ProjectType:  "bot",
			Technologies: pq.StringArray{"typescript", "discord.js"},
			ProjectDate:  time.Date(2025, time.July, 15, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   5,
			Title:        "Inventaris Alat Laboratorium DOSCOM",
			Tagline:      "Catat barang keluar masuk tanpa ada yang hilang misterius",
			Description:  "Dashboard pelacakan aset perangkat keras keras komunitas seperti router, switch, dan perangkat IoT.",
			Slug:         "inventaris-alat-laboratorium-doscom",
			ProjectType:  "web",
			Technologies: pq.StringArray{"react.js", "express.js", "mysql"},
			ProjectDate:  time.Date(2025, time.September, 9, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   49,
			Title:        "Automated Network Topology Mapper",
			Tagline:      "Scan active hosts and generate instant graphs",
			Description:  "Skrip otomatisasi pemindaian host dalam subnet jaringan lokal untuk menghasilkan diagram peta topologi aktif.",
			Slug:         "automated-network-topology-mapper",
			ProjectType:  "networking-tool",
			Technologies: pq.StringArray{"python", "nmap-api", "graphviz"},
			ProjectDate:  time.Date(2025, time.November, 28, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   2,
			Title:        "Algoritma Load Balancer Sederhana",
			Tagline:      "Distribusi trafik round-robin buatan sendiri",
			Description:  "Eksperimen pembuatan proxy server mini penyeimbang beban lalu lintas jaringan menggunakan bahasa pemrograman Go.",
			Slug:         "algoritma-load-balancer-sederhana",
			ProjectType:  "backend-service",
			Technologies: pq.StringArray{"golang", "net-http"},
			ProjectDate:  time.Date(2026, time.January, 9, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   59,
			Title:        "Media Creative Content Scheduler",
			Tagline:      "Jadwalkan postingan instagram feeds dalam satu klik",
			Description:  "Aplikasi manajemen kalender editorial konten sosial media DOSCOM agar publikasi info tetap konsisten.",
			Slug:         "media-creative-content-scheduler",
			ProjectType:  "web",
			Technologies: pq.StringArray{"next.js", "supabase"},
			ProjectDate:  time.Date(2025, time.August, 30, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "draft",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   16,
			Title:        "Custom Markdown Parser Engine",
			Tagline:      "Render berkas catatan menjadi HTML super cepat",
			Description:  "Pustaka open-source ringan untuk mengubah berkas sintaks teks markdown menjadi struktur elemen web HTML.",
			Slug:         "custom-markdown-parser-engine",
			ProjectType:  "library",
			Technologies: pq.StringArray{"javascript"},
			ProjectDate:  time.Date(2025, time.May, 14, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   31,
			Title:        "Sistem Pengumpul Tugas Praktikum",
			Tagline:      "Kumpulkan tugas tepat waktu sebelum server ditutup otomatis",
			Description:  "Portal pengumpulan file tugas praktikum pemrograman untuk mempermudah asisten dosen melakukan validasi timestamp.",
			Slug:         "sistem-pengumpul-tugas-praktikum",
			ProjectType:  "web",
			Technologies: pq.StringArray{"golang", "gin-gonic", "bootstrap"},
			ProjectDate:  time.Date(2025, time.November, 4, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			PengurusID:   40,
			Title:        "API Gateway Logger Interceptor",
			Tagline:      "Saring semua request masuk tanpa memperlambat performa",
			Description:  "Middleware interceptor berkinerja tinggi untuk mencatat jejak request IP, status code, dan konsumsi memori pada aplikasi backend.",
			Slug:         "api-gateway-logger-interceptor",
			ProjectType:  "backend-service",
			Technologies: pq.StringArray{"golang", "redis"},
			ProjectDate:  time.Date(2026, time.February, 20, 0, 0, 0, 0, time.Local),
			ImageURL:     "",
			Status:       "published",
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	imageID, imageURL, err := uploadFilePhoto(ctx, galleryService)
	if err != nil {
		return fmt.Errorf("failed to upload file photo %w", err)
	}
	for i := range workList {
		work := &workList[i]
		work.ImageURL = imageURL
		_, err := workModel.InsertWork(ctx, work)
		if err != nil {
			log.Printf("Failed to insert data work %v", err)
			return fmt.Errorf("failed to insert data work %w", err)
		}
		workGalleryData := make([]*entity.WorkGallery, len(imageID))
		for i, id := range imageID {
			workGalleryData[i] = &entity.WorkGallery{
				IDWork:    work.ID,
				IDGallery: id,
			}
		}
		err = SeedWorksGallery(db, workGalleryData)
		if err != nil {
			log.Printf("Failed to insert data work gallery %v", err)
			return fmt.Errorf("failed to insert data work gallery %w", err)
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
