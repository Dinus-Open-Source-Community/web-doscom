package seeder

import (
	"log"
	"time"
	"web_doscom/internal/database/model/entity"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedBlogs(db *gorm.DB) {
	now := time.Now()

	blogList := []entity.Blog{
		{
			AuthorID:     1,
			Title:        "Pengenalan Open Source",
			Slug:         "pengenalan-open-source",
			Content:      "Open Source adalah masa depan perangkat lunak.",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "https://dummyimage.com/800x600/000/fff&text=OpenSource",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
		{
			AuthorID:     2,
			Title:        "Cara Membuat Backend Go",
			Slug:         "cara-membuat-backend-go",
			Content:      "Tutorial membuat REST API dengan Gin dan GORM.",
			Kategori:     pq.StringArray{"technology"},
			ThumbnailURL: "https://dummyimage.com/800x600/000/fff&text=Golang+Gin",
			Status:       "published",
			PublishedAt:  &now,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	for _, b := range blogList {
		var existing entity.Blog

		err := db.Where("slug = ?", b.Slug).First(&existing).Error
		if err == gorm.ErrRecordNotFound {
			db.Create(&b)
		}
	}

	log.Println("Seed table 'blog' completed.")
}
