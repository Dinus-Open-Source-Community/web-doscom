package model

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ValidKategori = map[string]bool{
	"event":         true,
	"seminar":       true,
	"collaboration": true,
	"education":     true,
	"technology":    true,
	"work":          true,
	"activity":      true,
}

// BlogModel wraps DB for blog operations
type BlogModel struct {
	DB *gorm.DB
}

// Blog represents a blog post
type Blog struct {
<<<<<<< HEAD
	ID           int            `gorm:"primaryKey" json:"id"`
	AuthorID     int            `json:"author_id"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug" gorm:"unique"`
	Content      string         `json:"content"`
	Kategori     pq.StringArray `gorm:"type:text[]" json:"kategori"`
	ThumbnailURL string         `json:"thumbnail_url"`
	Status       string         `json:"status"`
	PublishedAt  time.Time      `json:"published_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type RequestBlog struct {
	AuthorID     int        `json:"author_id"`
	Title        string     `json:"title"`
	Slug         string     `json:"slug" gorm:"unique"`
	Content      string     `json:"content"`
	Kategori     []string   `json:"kategori"`
	ThumbnailURL string     `json:"thumbnail_url"`
	PublishedAt  *time.Time `json:"published_at"`
	Status       string     `json:"status" default:"draft"`
=======
	ID          int       `gorm:"primaryKey" json:"id"`
	GalleryID   int       `json:"id_asset"`
	WorkID      int       `json:"id_work"`
	ActivityID  int       `json:"id_activity"`
	PengurusID  int       `json:"id_pengurus"`
	Kategori    string    `json:"kategori"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug" gorm:"unique"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	IsPublished bool      `json:"is_published" default:"FALSE"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
>>>>>>> master
}

// RegisterBlog is used for creating a new blog
type RegisterBlog struct {
<<<<<<< HEAD
	ExistingID  []int      `form:"existingID_image"`
	Title       string     `form:"title" binding:"required"`
	Slug        string     `form:"slug" binding:"required"`
	Content     string     `form:"content" binding:"required"`
	Kategori    []string   `form:"kategori" binding:"required,kategori"`
	PublishedAt *time.Time `form:"published_at"`
	Status      string     `form:"status" default:"draft"`
=======
	Title       string    `json:"title" binding:"required"`
	GalleryID   int       `json:"id_asset" binding:"required"`
	Slug        string    `json:"slug" binding:"required"`
	Content     string    `json:"content" binding:"required"`
	Kategori    string    `json:"kategori" binding:"required,kategori"`
	PublishedAt time.Time `json:"published_at" binding:"required"`
	IsPublished bool      `json:"is_published" binding:"required"`
	WorkID      int       `json:"id_work" binding:"required"`
	ActivityID  int       `json:"id_activity" binding:"required"`
	PengurusID  int       `json:"id_pengurus" binding:"required"`
>>>>>>> master
}

// BlogPatch is used for updating a blog
type BlogPatch struct {
	ExistingID  []int      `form:"existingID_image" binding:"omitempty"`
	Title       string     `json:"title" binding:"omitempty"`
	Slug        string     `json:"slug" binding:"omitempty"`
	Content     string     `json:"content" binding:"omitempty"`
	Kategori    []string   `json:"kategori" binding:"omitempty,kategori"`
	Status      string     `json:"status" binding:"omitempty"`
	PublishedAt *time.Time `json:"published_at" binding:"omitempty"`
}

type BlogResponse struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	AuthorID     int            `json:"author_id"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug" gorm:"unique"`
	Content      string         `json:"content"`
	Kategori     []string       `json:"kategori"`
	ThumbnailURL string         `json:"thumbnail_url"`
	PublishedAt  time.Time      `json:"published_at"`
	BlogImage    []*BlogGallery `json:"blog_image"`
}

type BlogThumbnail struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Kategori     string `json:"kategori"`
	ThumbnailURL string `json:"thumbnail_url"`
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("kategori", func(fl validator.FieldLevel) bool {
			val := fl.Field().String()
			return ValidKategori[val]
		})
	}
}

// TableName sets the table name for Blog
func (Blog) TableName() string {
	return "blog"
}

// InsertBlog creates a new blog record
func (m *BlogModel) InsertBlog(blog *Blog) error {
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()
	if err := m.DB.Create(blog).Error; err != nil {
		return err
	}
	return nil
}

// GetBlogById fetches a blog by its ID
func (m *BlogModel) GetBlogById(id int) (*BlogResponse, error) {
	var blog BlogResponse

	query := `
		SELECT
			b.id,
			b.author_id,
			b.title,
			b.slug,
			b.content,
b.kategori,
			b.thumbnail_url,
			b.published_at,
			COALESCE(gallery.images, '[]'::json) AS gallery
		FROM blog b
		LEFT JOIN LATERAL (
			SELECT json_agg(
				json_build_object(
					'id', g.id,
					'file_url', g.file_url
				)
			) AS images FROM blog_gallery bg JOIN gallery g ON g.id = bg.id_gallery
			WHERE bg.id_blog = b.id
		)gallery ON true
		WHERE b.id = $1
	`

	if err := m.DB.Raw(query, id).Scan(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

// GetAllBlogs returns all blogs ordered by creation date
func (m *BlogModel) GetAllBlogs(ctx context.Context, limit, offset int) ([]BlogThumbnail, int, error) {
	var blogs []BlogThumbnail

	// pagination
	var totalData int64
	query := m.DB.WithContext(ctx).Where("status = 'published' AND published_at <= Now()")
	if err := query.Model(&Blog{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Model(&Blog{}).Select("id, title, slug, thumbnail_url, kategori").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(&blogs).
		Error; err != nil {
		return nil, 0, err
	}

	return blogs, int(totalData), nil
}

// get all blogs and get blog with kategory
func (m *BlogModel) GetBlogs(ctx context.Context, kategori []string, limit, offset int) ([]BlogThumbnail, int64, error) {

	query := m.DB.WithContext(ctx).Model(&Blog{})

	if len(kategori) >= 0 {
		query = query.Where("kategori && ?", pq.Array(kategori))
	}

	var totalData int64
	if err := query.Count(&totalData).Error; err != nil {
		return nil, 0, fmt.Errorf("error while count the data")
	}

	var blogs []BlogThumbnail
	if err := query.
		Select("id, title, slug, thumbnail_url, kategori").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Scan(&blogs).
		Error; err != nil {
		return nil, 0, fmt.Errorf("error while get data")
	}

	return blogs, totalData, nil
}

// update blog -> flexible update -> update partial and update full
func (m *BlogModel) UpdateBlogPartial(id int, data map[string]any) (*Blog, error) {
	var updateBlog Blog

	if err := m.DB.Model(&Blog{}).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(data).
		Scan(&updateBlog).
		Error; err != nil {
		return nil, err
	}
<<<<<<< HEAD

	return &updateBlog, nil
=======
	blog.Title = pathValue(input.Title, blog.Title)
	blog.Content = pathValue(input.Content, blog.Content)
	blog.Slug = pathValue(input.Slug, blog.Slug)
	blog.GalleryID = pathValue(input.GalleryID, blog.GalleryID)
	blog.WorkID = pathValue(input.WorkID, blog.WorkID)
	blog.ActivityID = pathValue(input.ActivityID, blog.ActivityID)
	blog.PengurusID = pathValue(input.PengurusID, blog.PengurusID)
	blog.Kategori = pathValue(input.Kategori, blog.Kategori)
	blog.PublishedAt = pathValue(input.PublishedAt, blog.PublishedAt)
	blog.IsPublished = pathValue(input.IsPublished, blog.IsPublished)
	blog.UpdatedAt = time.Now()
	if err := m.DB.Save(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
>>>>>>> master
}

// DeleteBlog deletes a blog record by ID
func (m *BlogModel) DeleteBlog(ctx context.Context, tx *gorm.DB, id int) error {
	if err := tx.WithContext(ctx).Delete(&Blog{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetBlogsByKategori returns all blogs by kategori
<<<<<<< HEAD
func (m *BlogModel) GetBlogsByKategori(ctx context.Context, kategori []string, limit, offset int) ([]BlogThumbnail, int, error) {
	var blogs []BlogThumbnail

	var totalData int64
	query := m.DB.WithContext(ctx).Where("kategori && ?", pq.Array(kategori))
	if err := query.Model(&Blog{}).Count(&totalData).Error; err != nil {
		return nil, 0, err
	}

	if err := query.
		Model(&Blog{}).
		Select("id, title, slug, thumbnail_url, kategori").
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(&blogs).
		Error; err != nil {
		return nil, 0, err
	}

	return blogs, int(totalData), nil
=======
func (m *BlogModel) GetBlogsByKategori(kategori string) ([]Blog, error) {
	var blogs []Blog
	if err := m.DB.Where("kategori = ?", kategori).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
>>>>>>> master
}
