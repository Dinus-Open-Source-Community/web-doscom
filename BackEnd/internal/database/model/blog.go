package model

import (
	"context"
	"fmt"
	"strings"
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
	ID           int            `gorm:"primaryKey" json:"id"`
	AuthorID     int            `json:"author_id"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug" gorm:"unique"`
	Content      string         `json:"content"`
	Kategori     pq.StringArray `gorm:"type:text[]" json:"kategori"`
	ThumbnailURL string         `json:"thumbnail_url"`
	Status       string         `json:"status"`
	PublishedAt  *time.Time     `json:"published_at"`
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
}

// RegisterBlog is used for creating a new blog
type RegisterBlog struct {
	ExistingID  []*int     `form:"existingID_image"`
	Title       string     `form:"title" binding:"required"`
	Slug        string     `form:"slug" binding:"required"`
	Content     string     `form:"content" binding:"required"`
	Kategori    []string   `form:"kategori" binding:"required,dive,kategori"`
	PublishedAt *time.Time `form:"published_at"`
	Status      string     `form:"status" default:"draft"`
}

// BlogPatch is used for updating a blog
type BlogPatch struct {
	ExistingID  []*int     `form:"existingID_image" binding:"omitempty"`
	Title       string     `form:"title" json:"title" binding:"omitempty"`
	Slug        string     `form:"slug" json:"slug" binding:"omitempty"`
	Content     string     `form:"content" json:"content" binding:"omitempty"`
	Kategori    []string   `form:"kategori" json:"kategori" binding:"omitempty,dive,kategori"`
	Status      string     `form:"status" json:"status" binding:"omitempty"`
	PublishedAt *time.Time `form:"published_at" json:"published_at" binding:"omitempty"`
}

type BlogResponse struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	AuthorID     int            `json:"author_id"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug" gorm:"unique"`
	Content      string         `json:"content"`
	Kategori     []string       `json:"kategori"`
	ThumbnailURL string         `json:"thumbnail_url"`
	PublishedAt  *time.Time     `json:"published_at"`
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
			val := strings.ToLower(fl.Field().String())
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

	if len(kategori) > 0 {
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

	return &updateBlog, nil
}

// DeleteBlog deletes a blog record by ID
func (m *BlogModel) DeleteBlog(ctx context.Context, tx *gorm.DB, id int) error {
	if err := tx.WithContext(ctx).Delete(&Blog{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetBlogsByKategori returns all blogs by kategori
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
}
