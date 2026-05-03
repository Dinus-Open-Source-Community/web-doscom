package dto

import "mime/multipart"

// DTO untuk frontend
type PengurusResponse struct {
	ID       int    `json:"id"`
	PhotoURL string `json:"photo_url"`
	Email    string `json:"email"`
	Divisi   string `json:"divisi"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Sosmed   string `json:"sosmed"`
	Period   string `json:"period"`
}

type PengurusPublicResponse struct {
	ID       int    `json:"id"`
	PhotoURL string `json:"photo_url"`
	Divisi   string `json:"divisi"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Sosmed   string `json:"sosmed"`
	Period   string `json:"period"`
}

// Untuk create/register pengurus
type RegisterPengurusRequest struct {
	UserID   int                   `form:"id_user"`
	PhotoURL *multipart.FileHeader `form:"photo_url"`
	Email    string                `form:"email"`
	Divisi   string                `form:"divisi" binding:"required,divisi"`
	Name     string                `form:"name" binding:"required" validate:"min=2,max=150"`
	Position string                `form:"position" binding:"required,position"`
	Sosmed   string                `form:"sosmed" binding:"omitempty,socialurl"`
	Period   string                `form:"period" binding:"required" validate:"max=50"`
}

// Untuk update/patch pengurus
type PengurusPatch struct {
	Email    string                `form:"email" binding:"omitempty,email"`
	Divisi   string                `form:"divisi" binding:"omitempty,divisi"`
	Name     string                `form:"name" binding:"omitempty"`
	Sosmed   string                `form:"sosmed" binding:"omitempty,socialurl"`
	Period   string                `form:"period" binding:"omitempty"`
	Position string                `form:"position" binding:"omitempty,position"`
	PhotoURL *multipart.FileHeader `form:"file" binding:"omitempty"`
}
