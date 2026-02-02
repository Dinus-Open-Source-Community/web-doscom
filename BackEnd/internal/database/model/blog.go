package model

import (
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"
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
	ID          int       `gorm:"primaryKey" json:"id"`
	WorkID      int       `json:"id_work"`
	PengurusID  int       `json:"id_pengurus"`
	Kategori    string    `json:"kategori"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug" gorm:"unique"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
	IsPublished bool      `json:"is_published" default:"FALSE"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RegisterBlog is used for creating a new blog
type RegisterBlog struct {
	ExistingID  []string  `form:"existingID_image" binding:"required"`
	Title       string    `form:"title" binding:"required"`
	GalleryID   int       `form:"id_asset" binding:"required"`
	Slug        string    `form:"slug" binding:"required"`
	Content     string    `form:"content" binding:"required"`
	Kategori    string    `form:"kategori" binding:"required,kategori"`
	PublishedAt time.Time `form:"published_at" binding:"required"`
	IsPublished bool      `form:"is_published" binding:"required"`
	WorkID      int       `form:"id_work" binding:"required"`
	PengurusID  int       `form:"id_pengurus" binding:"required"`
}

// BlogPatch is used for updating a blog
type BlogPatch struct {
	Title       *string    `json:"title" binding:"omitempty"`
	Content     *string    `json:"content" binding:"omitempty"`
	Slug        *string    `json:"slug" binding:"omitempty"`
	GalleryID   *int       `json:"id_asset" binding:"omitempty"`
	WorkID      *int       `json:"id_work" binding:"omitempty"`
	ActivityID  *int       `json:"id_activity" binding:"omitempty"`
	PengurusID  *int       `json:"id_pengurus" binding:"omitempty"`
	Kategori    *string    `json:"kategori" binding:"omitempty,kategori"`
	PublishedAt *time.Time `json:"published_at" binding:"omitempty"`
	IsPublished *bool      `json:"is_published" binding:"omitempty"`
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
func (m *BlogModel) GetBlogById(id int) (*Blog, error) {
	var blog Blog
	if err := m.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

// GetAllBlogs returns all blogs ordered by creation date
func (m *BlogModel) GetAllBlogs() ([]Blog, error) {
	var blogs []Blog
	if err := m.DB.Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

// pathValue returns the value pointed to by field if not nil, otherwise returns current.
func pathValue[T any](field *T, current T) T {
	if field != nil {
		return *field
	}
	return current
}

// UpdateBlog updates a blog record by ID
func (m *BlogModel) UpdateBlog(id int, input BlogPatch) (*Blog, error) {
	var blog Blog
	if err := m.DB.First(&blog, id).Error; err != nil {
		return nil, err
	}
	blog.Title = pathValue(input.Title, blog.Title)
	blog.Content = pathValue(input.Content, blog.Content)
	blog.Slug = pathValue(input.Slug, blog.Slug)
	blog.WorkID = pathValue(input.WorkID, blog.WorkID)
	blog.PengurusID = pathValue(input.PengurusID, blog.PengurusID)
	blog.Kategori = pathValue(input.Kategori, blog.Kategori)
	blog.PublishedAt = pathValue(input.PublishedAt, blog.PublishedAt)
	blog.IsPublished = pathValue(input.IsPublished, blog.IsPublished)
	blog.UpdatedAt = time.Now()
	if err := m.DB.Save(&blog).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

// DeleteBlog deletes a blog record by ID
func (m *BlogModel) DeleteBlog(id int) error {
	if err := m.DB.Delete(&Blog{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetBlogsByKategori returns all blogs by kategori
func (m *BlogModel) GetBlogsByKategori(kategori string) ([]Blog, error) {
	var blogs []Blog
	if err := m.DB.Where("kategori = ?", kategori).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}
