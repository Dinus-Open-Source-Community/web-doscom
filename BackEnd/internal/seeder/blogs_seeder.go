package seeder

import (
	"context"
	"log"
	"time"
	"web_doscom/internal/database/model/entity"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedBlogs(db *gorm.DB) {
	now := time.Now()

	blogModel := entity.BlogModel{DB: db}
	ctx := context.Background()

	blogList := []entity.Blog{
		{
			AuthorID:     3,
			Title:        "Pengenalan Open Source",
			Slug:         "pengenalan-open-source",
			Content:      "Open Source adalah masa depan perangkat lunak.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     3,
			Title:        "Cara Membuat Backend Go",
			Slug:         "cara-membuat-backend-go",
			Content:      "Tutorial membuat REST API dengan Gin dan GORM.",
			Kategori:     pq.StringArray{"technology"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     3,
			Title:        "Mengenal JWT Authentication di Golang",
			Slug:         "mengenal-jwt-authentication-di-golang",
			Content:      "Panduan implementasi JWT authentication menggunakan Gin framework.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},

		{
			AuthorID:     3,
			Title:        "Belajar PostgreSQL untuk Backend Developer",
			Slug:         "belajar-postgresql-untuk-backend-developer",
			Content:      "Dasar penggunaan PostgreSQL untuk aplikasi backend modern.",
			Kategori:     pq.StringArray{"education", "technology"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},

		{
			AuthorID:     3,
			Title:        "Deploy Aplikasi Go ke VPS Linux",
			Slug:         "deploy-aplikasi-go-ke-vps-linux",
			Content:      "Tutorial deploy aplikasi Golang menggunakan svstemd dan Nginx.",
			Kategori:     pq.StringArray{"technology", "work"},
			ThumbnailURL: "",
			Status:       "draft",
			PublishedAt:  nil,
			CreatedAt:    now,
			UpdatedAt:    now,
		},

		{
			AuthorID:     3,
			Title:        "Membuat CRUD API dengan Gin dan GORM",
			Slug:         "membuat-crud-api-dengan-gin-dan-gorm",
			Content:      "Step by step membuat CRUD API menggunakan Golang, Gin, dan GORM.",
			Kategori:     pq.StringArray{"activity", "technology"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},

		{
			AuthorID:     3,
			Title:        "Pengenalan Docker untuk Developer",
			Slug:         "pengenalan-docker-untuk-developer",
			Content:      "Belajar dasar Docker dan containerization untuk development workflow.",
			Kategori:     pq.StringArray{"education", "technology", "work"},
			ThumbnailURL: "",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	for i := range blogList {
		if err := blogModel.InsertBlog(ctx, &blogList[i]); err != nil {
			log.Printf("Failed to insert data blog %v", err)
			return
		}
	}

	log.Println("Seed table 'blog' completed.")
}

func RunSeedBlogs(db *gorm.DB) {
	// for now, only run blogs seeder -> bacause im too lazy to create a seeder for blog_gallery
	SeedBlogs(db)

	// call seeder for blog_gallery -> if it exist
}
