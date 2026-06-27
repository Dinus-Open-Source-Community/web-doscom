package seeder

import (
	"context"
	"fmt"
	"log"
	"time"

	"web_doscom/internal/database/model/entity"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type blogSeedData struct {
	Blog      entity.Blog
	ImageURLs []string
}

func SeedBlogs(db *gorm.DB) error {
	ctx := context.Background()
	now := time.Now()

	blogModel := entity.BlogModel{DB: db}
	galleryModel := entity.GalleryModel{DB: db}
	blogGalleryModel := entity.BlogGalleryModel{DB: db}

	const uploaderUserID = 14

	blogList := []blogSeedData{
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Pengenalan Git untuk Pemula",
				Slug:        "pengenalan-git-untuk-pemula",
				Content:     "Panduan memahami version control menggunakan Git.",
				Kategori:    pq.StringArray{"technology", "education"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://picsum.photos/seed/git-1/1200/800",
				"https://picsum.photos/seed/git-2/1200/800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Panduan Dasar Canva",
				Slug:        "panduan-dasar-canva",
				Content:     "Tips menggunakan Canva untuk membuat desain grafis.",
				Kategori:    pq.StringArray{"activity", "education"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://picsum.photos/seed/canva-1/1200/800",
				"https://picsum.photos/seed/canva-2/1200/800",
			},
		},

		// Tambahkan data blog lainnya menggunakan pola yang sama.
	}

	for i := range blogList {
		seed := &blogList[i]
		blog := &seed.Blog

		if len(seed.ImageURLs) == 0 {
			return fmt.Errorf("blog %q does not have images", blog.Title)
		}

		// URL pertama digunakan sebagai thumbnail.
		blog.ThumbnailURL = seed.ImageURLs[0]

		if err := blogModel.InsertBlog(ctx, blog); err != nil {
			return fmt.Errorf(
				"failed to insert blog %q: %w",
				blog.Title,
				err,
			)
		}

		for _, imageURL := range seed.ImageURLs {
			gallery := &entity.Gallery{
				IDUsers:     uploaderUserID,
				GalleryName: "Gambar blog " + blog.Title,
				GalleryType: "blog",
				Description: "Gambar eksternal blog " + blog.Title,
				EventDate:   now,
				FileURL:     imageURL,
			}

			// file_upload_id harus NULL karena gambar tidak di-upload
			// ke MinIO dan tidak memiliki metadata file_uploads.
			result := galleryModel.DB.
				WithContext(ctx).
				Omit("FileUploadID").
				Create(gallery)

			if result.Error != nil {
				return fmt.Errorf(
					"failed to insert gallery blog %q: %w",
					blog.Title,
					result.Error,
				)
			}

			relation := &entity.BlogGallery{
				BlogID:    blog.ID,
				GalleryID: gallery.ID,
			}

			if _, err := blogGalleryModel.InsertBlogGallery(
				relation,
			); err != nil {
				return fmt.Errorf(
					"failed to relate gallery with blog %q: %w",
					blog.Title,
					err,
				)
			}
		}
	}

	log.Println("Seed table 'blog' completed.")
	log.Println("Seed table 'blog_gallery' completed.")

	return nil
}

func RunSeedBlogs(db *gorm.DB) {
	if err := SeedBlogs(db); err != nil {
		log.Fatalf("Failed to seed blogs: %v", err)
	}
}
