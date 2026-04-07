package seeder

import (
	"log"
	"time"

	"web_doscom/internal/database/model"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedBlogs(db *gorm.DB) {
	blogList := []model.Blog{
		{
			AuthorID:     1,
			Title:        "Pengenalan Open Source",
			Slug:         "pengenalan-open-source",
			Content:      "<p>Open Source adalah masa depan perangkat lunak. Artikel ini membahas mengapa kita harus mulai berkontribusi.</p>",
			Kategori:     pq.StringArray{"technology", "education"},
			ThumbnailURL: "https://dummyimage.com/800x600/000/fff&text=OpenSource",
			Status:       "published",
			PublishedAt:  time.Now(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			AuthorID:     2,
			Title:        "Cara Membuat Backend Go",
			Slug:         "cara-membuat-backend-go",
			Content:      "<p>Berikut adalah langkah-langkah membuat REST API dengan Golang menggunakan Gin dan GORM.</p>",
			Kategori:     pq.StringArray{"technology", "event"},
			ThumbnailURL: "https://dummyimage.com/800x600/000/fff&text=Golang+Gin",
			Status:       "published",
			PublishedAt:  time.Now(),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	for _, b := range blogList {
		db.FirstOrCreate(&b, model.Blog{Slug: b.Slug})
	}

	log.Println("Seed table 'blog' completed.")
}
