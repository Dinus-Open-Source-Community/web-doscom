package dto

import (
	"time"
)

// use to get input from client
type CreateRequestBlog struct {
	ExistingID  []*int     `form:"existingID_image"`
	Title       string     `form:"title" binding:"required"`
	Slug        string     `form:"slug" binding:"required"`
	Content     string     `form:"content" binding:"required"`
	Kategori    []string   `form:"kategori" binding:"required,dive,kategori"`
	PublishedAt *time.Time `form:"published_at"`
	Status      string     `form:"status" binding:"required"`
}

// use to internal process -> insert to database
type BlogPayload struct {
	AuthorID     int        `json:"author_id"`
	Content      string     `json:"content"`
	ExistingID   []*int     `json:"existingID_image"`
	Kategori     []string   `json:"kategori"`
	PublishedAt  *time.Time `json:"published_at"`
	Slug         string     `json:"slug"`
	Status       string     `json:"status" default:"draft"`
	Title        string     `json:"title"`
	ThumbnailURL string     `json:"thumbnail_url"`
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
	ID           int                    `gorm:"primaryKey" json:"id"`
	AuthorID     int                    `json:"author_id"`
	Title        string                 `json:"title"`
	Slug         string                 `json:"slug" gorm:"unique"`
	Content      string                 `json:"content"`
	Kategori     []string               `json:"kategori"`
	ThumbnailURL string                 `json:"thumbnail_url"`
	PublishedAt  *time.Time             `json:"published_at"`
	BlogImage    []*BlogGalleryResponse `json:"blog_image"`
}

type BlogThumbnail struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Slug         string `json:"slug"`
	Kategori     string `json:"kategori"`
	ThumbnailURL string `json:"thumbnail_url"`
}
