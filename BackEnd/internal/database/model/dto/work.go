package dto

import (
	"encoding/json"
	"time"
)

type CreateRequestWork struct {
	ExistingID   []*int    `form:"existingID_image"`
	PengurusID   int       `form:"pengurus_id" binding:"required"`
	Title        string    `form:"title" binding:"required"`
	Tagline      string    `form:"tagline" binding:"required"`
	Description  string    `form:"description" binding:"required"`
	Slug         string    `form:"slug" binding:"required"`
	ProjectType  string    `form:"project_type" binding:"required"`
	Technologies []string  `form:"technologies[]" binding:"required"`
	ProjectDate  time.Time `form:"project_date" binding:"required"`
	Status       string    `form:"status" binding:"required"`
}

type WorkResponseClient struct {
	ID           int             `json:"id"`
	Title        string          `json:"title"`
	Tagline      string          `json:"tagline"`
	Description  string          `json:"description"`
	Slug         string          `json:"slug"`
	ProjectType  string          `json:"project_type"`
	Technologies []string        `json:"technologies"`
	ProjectDate  time.Time       `json:"project_date"`
	ImageURL     string          `json:"image_url"`
	Gallery      json.RawMessage `json:"gallery"`
}

type WorkResponseInternal struct {
	ID           int             `json:"id"`
	Title        string          `json:"title"`
	Tagline      string          `json:"tagline"`
	Description  string          `json:"description"`
	Slug         string          `json:"slug"`
	ProjectType  string          `json:"project_type"`
	Status       string          `json:"status"`
	Technologies []string        `json:"technologies"`
	ProjectDate  time.Time       `json:"project_date"`
	ImageURL     string          `json:"image_url"`
	Gallery      json.RawMessage `json:"gallery"`
}
type WorkPatch struct {
	PengurusID   int       `json:"pengurus_id"`
	Title        string    `json:"title"`
	Tagline      string    `json:"tagline"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ProjectType  string    `json:"project_type"`
	ProjectDate  time.Time `json:"project_date"`
	Status       string    `json:"status"`
	Technologies []string  `json:"technologies"`
	ExistingID   []*int    `json:"existingID_image"`
}

type WorkPayloadUpdate struct {
	PengurusID   int       `json:"pengurus_id"`
	Title        string    `json:"title"`
	Tagline      string    `json:"tagline"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ProjectType  string    `json:"project_type"`
	ProjectDate  time.Time `json:"project_date"`
	Status       string    `json:"status"`
	Technologies []string  `json:"technologies"`
	UpdatedAt    time.Time `json:"updated_at"`
	ImageURL     string    `json:"image_url"`
}

type WorkUpdateResponse struct {
	ID           int                    `json:"id"`
	PengurusID   int                    `json:"pengurus_id"`
	Title        string                 `json:"title"`
	Tagline      string                 `json:"tagline"`
	Description  string                 `json:"description"`
	Slug         string                 `json:"slug"`
	ProjectType  string                 `json:"project_type"`
	Technologies []string               `json:"technologies"`
	ProjectDate  time.Time              `json:"project_date"`
	ImageURL     string                 `json:"image_url"`
	Status       string                 `json:"status"`
	UpdatedAt    time.Time              `json:"updated_at"`
	CreatedAt    time.Time              `json:"created_at"`
	WorkGallery  []*WorkGalleryResponse `json:"work_gallery"`
}
