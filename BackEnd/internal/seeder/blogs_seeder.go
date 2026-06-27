package seeder

import (
	"context"
	"fmt"
	"log"
	"time"
	"web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/service"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func seedBlogsGallery(db *gorm.DB, blogGallery []*entity.BlogGallery) error {
	blogModel := entity.BlogGalleryModel{DB: db}
	ctx := context.Background()

	_, err := blogModel.InsertBlogGalleryMultiple(ctx, blogGallery)
	if err != nil {
		return fmt.Errorf("failed to insert data %w", err)
	}

	return nil
}

func SeedBlogs(db *gorm.DB, galleryService *service.GalleryService) error {
	now := time.Now()

	blogModel := entity.BlogModel{DB: db}
	ctx := context.Background()

	blogList := []entity.Blog{
		{
			AuthorID:     14,
			Title:        "Pengenalan Git untuk Pemula",
			Slug:         "pengenalan-git-untuk-pemula",
			Content:      "Panduan langkah demi langkah memahami version control system menggunakan Git bagi pemula.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Panduan Dasar Canva untuk Desain Grafis",
			Slug:         "panduan-dasar-canva-untuk-desain-grafis",
			Content:      "Tips dan trik memanfaatkan Canva untuk membuat microblog dan aset visual media sosial dengan cepat.",
			Kategori:     pq.StringArray{"activity", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Desain UI/UX Menggunakan Figma",
			Slug:         "desain-uiux-menggunakan-figma",
			Content:      "Belajar memahami wireframing, prototyping, dan komponen dasar desain UI/UX di Figma.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Tips Konsistensi Belajar Programming",
			Slug:         "tips-konsistensi-belajar-programming",
			Content:      "Membangun kebiasaan ngoding setiap hari tanpa terkena burnout atau kehilangan motivasi.",
			Kategori:     pq.StringArray{"education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Apa itu Open Source Software?",
			Slug:         "apa-itu-open-source-software",
			Content:      "Mengenal budaya open-source, lisensi software, dan bagaimana cara berkontribusi ke repositori publik.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Strategi Branding Komunitas Kampus",
			Slug:         "strategi-branding-komunitas-kampus",
			Content:      "Bagaimana mengelola visual identitas dan gaya bahasa organisasi agar menarik minat mahasiswa baru.",
			Kategori:     pq.StringArray{"activity", "work"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Dasar-dasar HTML dan CSS",
			Slug:         "dasar-dasar-html-dan-css",
			Content:      "Mengenal struktur markup web dan dasar-dasar styling menggunakan properti CSS modern.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Cara Membuat Komposisi Warna yang Menarik",
			Slug:         "cara-membuat-komposisi-warna-yang-menarik",
			Content:      "Teori warna dasar untuk desainer grafis dan web agar konten visual terlihat harmonis.",
			Kategori:     pq.StringArray{"activity", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Mengenal Framework Tailwind CSS",
			Slug:         "mengenal-framework-tailwind-css",
			Content:      "Mengapa utilitas kelas dari Tailwind CSS mempermudah proses slicing desain web.",
			Kategori:     pq.StringArray{"technology"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Manajemen Waktu Kuliah Sambil Organisasi",
			Slug:         "manajemen-waktu-kuliah-sambil-organisasi",
			Content:      "Berbagi pengalaman membagi waktu antara tugas akademik dan tanggung jawab sebagai pengurus komunitas.",
			Kategori:     pq.StringArray{"education", "work"},
			ThumbnailURL: "",
			Status:       "draft",
			PublishedAt:  nil,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Pengenalan Linux untuk Sehari-hari",
			Slug:         "pengenalan-linux-untuk-sehari-hari",
			Content:      "Migrasi ke ekosistem open-source dengan menggunakan distro Linux ramah pemula.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Pentingnya Tipografi dalam Desain Konten",
			Slug:         "pentingnya-tipografi-dalam-desain-konten",
			Content:      "Memilih font, mengatur kerning, dan leading agar informasi di media sosial lebih mudah dibaca.",
			Kategori:     pq.StringArray{"activity", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Tutorial Sederhana JavaScript DOM",
			Slug:         "tutorial-sederhana-javascript-dom",
			Content:      "Manipulasi elemen HTML secara dinamis menggunakan vanilla JavaScript Document Object Model.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Mengenal Struktur Data Array dan Object",
			Slug:         "mengenal-struktur-data-array-dan-object",
			Content:      "Pondasi logika dasar pemrograman mengenai cara menyimpan koleksi data secara terstruktur.",
			Kategori:     pq.StringArray{"education", "technology"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Alasan Harus Ikut Komunitas DOSCOM",
			Slug:         "alasan-harus-ikut-komunitas-doscom",
			Content:      "Benefit belajar bareng, memperluas relasi, hingga membangun portofolio proyek open-source nyata.",
			Kategori:     pq.StringArray{"activity", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Cara Mengatasi Writer's Block saat Bikin Konten",
			Slug:         "cara-mengatasi-writers-block-saat-bikin-konten",
			Content:      "Trik menemukan ide segar saat kehabisan bahan tulisan artikel blog atau caption instagram.",
			Kategori:     pq.StringArray{"activity", "work"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Apa Itu REST API dan Cara Kerjanya?",
			Slug:         "apa-itu-rest-api-dan-cara-kerjanya",
			Content:      "Memahami jembatan komunikasi data antara sisi frontend dan backend secara arsitektural.",
			Kategori:     pq.StringArray{"technology"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Memahami Konsep Clean Code",
			Slug:         "memahami-konsep-clean-code",
			Content:      "Menulis kode yang rapi, mudah dibaca, dan gampang di-maintain berkolaborasi dalam tim.",
			Kategori:     pq.StringArray{"technology", "work"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Tips Lolos Kurasi Konten Instagram",
			Slug:         "tips-lolos-kurasi-konten-instagram",
			Content:      "Alur standard operating procedure (SOP) internal divisi media kreatif sebelum mempublikasikan desain.",
			Kategori:     pq.StringArray{"activity"},
			ThumbnailURL: "",
			Status:       "draft",
			PublishedAt:  nil,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     14,
			Title:        "Mengenal Version Control System",
			Slug:         "mengenal-version-control-system",
			Content:      "Mengapa tim developer besar wajib memakai sistem pelacakan revisi kode program.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	imageID, imageURL, err := uploadFilePhoto(ctx, galleryService)
	if err != nil {
		log.Fatalf("Failed to upload file photo %v", err)
	}
	for i := range blogList {
		blog := &blogList[i]
		blog.ThumbnailURL = imageURL
		if err := blogModel.InsertBlog(ctx, &blogList[i]); err != nil {
			log.Printf("Failed to insert data blog %v", err)
			return fmt.Errorf("failed to insert data %w", err)
		}
		blogGalleryData := make([]*entity.BlogGallery, len(imageID))
		for i, id := range imageID {
			blogGalleryData[i] = &entity.BlogGallery{
				BlogID:    blog.ID,
				GalleryID: id,
			}
		}
		err = seedBlogsGallery(db, blogGalleryData)
		if err != nil {
			log.Printf("Failed to insert data blog gallery %v", err)
			return fmt.Errorf("failed to insert data %w", err)
		}
	}

	log.Println("Seed table 'blog' completed.")
	return nil
}

func RunSeedBlogs(db *gorm.DB, minioClient *config.MinioClient) {

	models := database.NewModel(db)
	storageService := service.NewStorageService(
		minioClient,
		&models.FileUploads,
	)

	galleryService := service.NewGalleryService(
		&models.Gallery,
		storageService,
	)

	if err := SeedBlogs(db, galleryService); err != nil {
		log.Fatalf("Failed to seed blogs: %v", err)
		return
	}
}
