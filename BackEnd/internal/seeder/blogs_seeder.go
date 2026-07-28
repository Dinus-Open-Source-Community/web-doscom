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
				Content:     "Panduan langkah demi langkah memahami version control system menggunakan Git bagi pemula.",
				Kategori:    pq.StringArray{"technology", "education"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1618401471353-b98afee0b2eb?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1556075798-482aefae607a?auto=format&fit=crop&w=1200&h=800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Panduan Dasar Canva untuk Desain Grafis",
				Slug:        "panduan-dasar-canva-untuk-desain-grafis",
				Content:     "Tips dan trik memanfaatkan Canva untuk membuat microblog dan aset visual media sosial dengan cepat.",
				Kategori:    pq.StringArray{"activity", "education"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1626785774573-4b799315345d?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1618005182384-a83a8bd57fbe?auto=format&fit=crop&w=1200&h=800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Membangun REST API dengan Golang",
				Slug:        "membangun-rest-api-dengan-golang",
				Content:     "Tutorial arsitektur backend clean code untuk membangun REST API performa tinggi menggunakan Go.",
				Kategori:    pq.StringArray{"technology"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1605379399642-870262d3d051?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1526374965328-7f61d4dc18c5?auto=format&fit=crop&w=1200&h=800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Budaya Open Source dan Keunggulannya",
				Slug:        "budaya-open-source-dan-keunggulannya",
				Content:     "Mengenal esensi ekosistem perangkat lunak terbuka, lisensi kode, dan kontribusi komunitas global.",
				Kategori:    pq.StringArray{"technology", "education"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1629654297299-c8506221ca97?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1542831371-29b0f74f9713?auto=format&fit=crop&w=1200&h=800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Mengenal Docker untuk Containerization",
				Slug:        "mengenal-docker-untuk-containerization",
				Content:     "Pondasi dasar menggunakan Docker container demi standarisasi alur kerja deployment tim developer.",
				Kategori:    pq.StringArray{"technology", "work"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1607799279861-4dd421887fb3?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1587620962725-abab7fe55159?auto=format&fit=crop&w=1200&h=800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Tips Desain UI/UX Mobile App di Figma",
				Slug:        "tips-desain-uiux-mobile-app-di-figma",
				Content:     "Strategi penyusunan komponen, auto-layout, dan wireframing interaktif aplikasi mobile di Figma.",
				Kategori:    pq.StringArray{"activity", "education"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1581291518633-83b4ebd1d83e?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1586717791821-3f44a563fa4c?auto=format&fit=crop&w=1200&h=800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Pentingnya Menerapkan Prinsip Clean Code",
				Slug:        "pentingnya-menerapkan-prinsip-clean-code",
				Content:     "Bagaimana menulis kode program yang rapi, efisien, low-latency, dan mudah dirawat secara jangka panjang.",
				Kategori:    pq.StringArray{"technology", "work"},
				Status:      "published",
				PublishedAt: &now,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1555066931-4365d14bab8c?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1517694712202-14dd9538aa97?auto=format&fit=crop&w=1200&h=800",
			},
		},
		{
			Blog: entity.Blog{
				AuthorID:    14,
				Title:       "Manajemen Waktu Antara Kuliah dan Organisasi",
				Slug:        "manajemen-waktu-antara-kuliah-dan-organisasi",
				Content:     "Tips menyeimbangkan prioritas akademik mahasiswa informatika dengan keaktifan berkontribusi di komunitas kampus.",
				Kategori:    pq.StringArray{"education"},
				Status:      "draft",
				PublishedAt: nil,
			},
			ImageURLs: []string{
				"https://images.unsplash.com/photo-1522202176988-66273c2fd55f?auto=format&fit=crop&w=1200&h=800",
				"https://images.unsplash.com/photo-1506784983877-45594efa4cbe?auto=format&fit=crop&w=1200&h=800",
			},
		},
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
