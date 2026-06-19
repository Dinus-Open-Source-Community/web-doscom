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

	"gorm.io/gorm"
)

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

	historyList := []entity.HistoryTimeline{
		{
			IDAuthor:     authorID,
			Title:        "Awal Berdirinya DOSCOM",
			Year:         "2008",
			Description:  "Perjalanan awal berdirinya Dinus Open Source Community oleh para founder dengan misi memmasyarakatkan open source.",
			DisplayOrder: 1,
		},
		{
			IDAuthor:     authorID,
			Title:        "Gerakan Migrasi Linux Kampus",
			Year:         "2009",
			Description:  "Menggalakkan kampanye migrasi sistem operasi open-source di lingkungan internal mahasiswa Fakultas Ilmu Komputer.",
			DisplayOrder: 2,
		},
		{
			IDAuthor:     authorID,
			Title:        "DOSCOM Goes to School",
			Year:         "2010",
			Description:  "Inisiasi program sosialisasi, instalasi, dan pelatihan sistem operasi Linux ke berbagai sekolah menengah di Semarang.",
			DisplayOrder: 3,
		},
		{
			IDAuthor:     authorID,
			Title:        "Rilis Distro Lokal Kampus",
			Year:         "2011",
			Description:  "Berkolaborasi dalam merilis varian distro Linux lokal yang dikustomisasi khusus untuk kebutuhan edukasi mahasiswa.",
			DisplayOrder: 4,
		},
		{
			IDAuthor:     authorID,
			Title:        "Workshop Pemrograman Terbuka",
			Year:         "2013",
			Description:  "Mulai mengadakan kelas pelatihan rutin bahasa pemrograman gratis seperti Python dan PHP yang terbuka untuk umum.",
			DisplayOrder: 5,
		},
		{
			IDAuthor:     authorID,
			Title:        "Peresmian Laboratorium Resmi DOSCOM",
			Year:         "2015",
			Description:  "Mendapatkan fasilitas ruang laboratorium resmi dari fakultas sebagai pusat riset dan tempat berkumpulnya anggota.",
			DisplayOrder: 6,
		},
		{
			IDAuthor:     authorID,
			Title:        "Kontribusi Kustomisasi Distro Nasional",
			Year:         "2016",
			Description:  "Aktif berkontribusi dalam pengembangan paket komponen dan pembaruan repositori lokal untuk distro IGOS Nusantara.",
			DisplayOrder: 7,
		},
		{
			IDAuthor:     authorID,
			Title:        "DOSCOM Hackathon Pertama",
			Year:         "2017",
			Description:  "Menyelenggarakan kompetisi hackathon internal untuk memecahkan berbagai permasalahan utilitas kampus.",
			DisplayOrder: 8,
		},
		{
			IDAuthor:     authorID,
			Title:        "Perayaan Satu Dekade DOSCOM",
			Year:         "2018",
			Description:  "Memperingati 10 tahun kontribusi aktif menyebarkan semangat open-source dan melahirkan talenta IT berbakat.",
			DisplayOrder: 9,
		},
		{
			IDAuthor:     authorID,
			Title:        "Restrukturisasi Divisi Komunitas",
			Year:         "2019",
			Description:  "Penyegaran struktur organisasi dengan memetakan fokus keahlian menjadi Pemrograman, Jaringan, Media, dan Data.",
			DisplayOrder: 10,
		},
		{
			IDAuthor:     authorID,
			Title:        "Adaptasi DOSCOM Online",
			Year:         "2020",
			Description:  "Menghadapi tantangan pandemi dengan memindahkan seluruh kegiatan workshop rutin menjadi webinar interaktif online.",
			DisplayOrder: 11,
		},
		{
			IDAuthor:     authorID,
			Title:        "Peluncuran Open Source Repository Mirror",
			Year:         "2021",
			Description:  "Menyediakan server mirror lokal mandiri untuk mempercepat proses unduhan berbagai distro Linux populer di area kampus.",
			DisplayOrder: 12,
		},
		{
			IDAuthor:     authorID,
			Title:        "Inisiasi DOSCOM University",
			Year:         "2022",
			Description:  "Meluncurkan program bootcamp intensif bertajuk DOSCOM University untuk menyaring dan membimbing calon talenta baru.",
			DisplayOrder: 13,
		},
		{
			IDAuthor:     authorID,
			Title:        "Fokus Pengembangan Proyek Riil",
			Year:         "2023",
			Description:  "Mulai mewajibkan pengerjaan produk open-source skala tim kecil sebagai syarat kelulusan magang anggota baru.",
			DisplayOrder: 14,
		},
		{
			IDAuthor:     authorID,
			Title:        "Ekspansi Angkatan Baru",
			Year:         "2023",
			Description:  "Rekrutmen besar-besaran yang sukses menyerap antusiasme puluhan pengurus aktif dari berbagai lintas program studi.",
			DisplayOrder: 15,
		},
		{
			IDAuthor:     authorID,
			Title:        "DOSCOM University 2024",
			Year:         "2024",
			Description:  "Sukses menyelenggarakan rangkaian workshop berskala besar yang fokus pada tech stack modern seperti Go, Rust, dan DevOps.",
			DisplayOrder: 16,
		},
		{
			IDAuthor:     authorID,
			Title:        "Modernisasi Server Infrastruktur Lab",
			Year:         "2024",
			Description:  "Migrasi arsitektur server laboratorium DOSCOM menggunakan teknologi containerization berbasis Docker.",
			DisplayOrder: 17,
		},
		{
			IDAuthor:     authorID,
			Title:        "DOSCOM University 2025",
			Year:         "2025",
			Description:  "Melatih ratusan peserta luar lewat kelas intensif web development, database management, dan advanced networking.",
			DisplayOrder: 18,
		},
		{
			IDAuthor:     authorID,
			Title:        "Rancang Bangun Platform Portofolio Terpadu",
			Year:         "2025",
			Description:  "Mulai mendevelop platform internal terintegrasi untuk mendata seluruh aset digital dan hasil karya pengurus.",
			DisplayOrder: 19,
		},
		{
			IDAuthor:     authorID,
			Title:        "Rencana Inovasi Global Terbuka",
			Year:         "2026",
			Description:  "Terus berkomitmen menciptakan ekosistem backend yang efisien, latency rendah, serta berkontribusi di kancah open-source global.",
			DisplayOrder: 20,
		},
	}

	fileUploads := make([]*dto.UploadFileRequest, 0, totalPhoto)
	cleanups := make([]func(), 0, totalPhoto)

	// Buka enam gambar satu kali.
	for imageNumber := 1; imageNumber <= totalPhoto; imageNumber++ {
		imagePath := filepath.Join(
			"storage",
			"uploads",
			"history",
			fmt.Sprintf("%d.jpg", imageNumber),
		)

		header, file, cleanup, err := utils.OpenSeedImage(imagePath)
		if err != nil {
			for _, cleanup := range cleanups {
				cleanup()
			}

			return fmt.Errorf(
				"failed to open history photo %d: %w",
				imageNumber,
				err,
			)
		}

		cleanups = append(cleanups, cleanup)
		fileUploads = append(fileUploads, &dto.UploadFileRequest{
			FileHeader: header,
			File:       file,
			Folder:     "history",
			UserID:     uint(authorID),
		})
	}

	galleryPayload := &dto.GalleryInsert{
		IDUsers:     authorID,
		GalleryName: "Dokumentasi sejarah DOSCOM",
		GalleryType: "history",
		Description: "Foto dokumentasi perjalanan DOSCOM",
		EventDate:   time.Now(),
	}

	// Upload enam gambar hanya satu kali.
	galleries, err := galleryService.UploadAndInsertGalleryMultiple(
		ctx,
		galleryPayload,
		fileUploads,
	)

	// File lokal sudah tidak digunakan setelah proses upload selesai.
	for _, cleanup := range cleanups {
		cleanup()
	}

	if err != nil {
		return fmt.Errorf(
			"failed to upload history galleries: %w",
			err,
		)
	}

	if len(galleries) != totalPhoto {
		return fmt.Errorf(
			"expected %d history photos, got %d",
			totalPhoto,
			len(galleries),
		)
	}

	// Setiap timeline menggunakan enam URL gallery yang sama.
	for i := range historyList {
		history := &historyList[i]

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
					"failed to insert photo for history %q: %w",
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
