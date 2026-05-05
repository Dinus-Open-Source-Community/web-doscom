package dto

import "time"

type GalleryInsert struct {
	IDUsers      int       `form:"id_users"`
	FileUploadID int       `form:"file_upload_id"`
	GalleryName  string    `form:"gallery_name"`
	GalleryType  string    `form:"gallery_type"`
	Description  string    `form:"description"`
	EventDate    time.Time `form:"event_date" time_format:"2006-01-02"`
}

type CreateGallery struct {
	GalleryName string    `form:"gallery_name" binding:"required"`
	GalleryType string    `form:"gallery_type" binding:"required"`
	Description string    `form:"description" binding:"required"`
	EventDate   time.Time `form:"event_date" binding:"required" time_format:"2006-01-02"`
}

type GalleryResponse struct {
	ID           int       `json:"id"`
	IDUsers      int       `json:"id_users"`
	FileUploadID int       `json:"file_upload_id"`
	GalleryName  string    `json:"gallery_name"`
	GalleryType  string    `json:"gallery_type"`
	Description  string    `json:"description"`
	EventDate    time.Time `json:"event_date"`
	FileURL      string    `json:"file_url"`
}
