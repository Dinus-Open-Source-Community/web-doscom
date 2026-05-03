package dto

import (
	"mime/multipart"
)

type FileUploadResponse struct {
	ID               uint   `json:"id"`
	UserID           uint   `json:"user_id"`
	Category         string `json:"category"`
	OriginalFilename string `json:"original_filename"`
	StoredFilename   string `json:"stored_filename"`
	FileSize         int64  `json:"file_size"`
	ContentType      string `json:"content_type"`
	FileURL          string `json:"file_url"`
}

type UploadFileRequest struct {
	FileHeader *multipart.FileHeader
	File       multipart.File
	Folder     string
	UserID     uint
}
