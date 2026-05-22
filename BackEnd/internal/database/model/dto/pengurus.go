package dto

import (
	"encoding/json"
	"mime/multipart"
	"time"
)

// admin consume
type PengurusResponse struct {
	ID               int                      `json:"id"`
	IDUser           int                      `json:"id_user"`
	PhotoURL         string                   `json:"photo_url"`
	Email            string                   `json:"email"`
	Divisi           string                   `json:"divisi"`
	Name             string                   `json:"name"`
	Position         string                   `json:"position"`
	Sosmed           []PengurusSosmedResponse `json:"sosmed" gorm:"-"`
	StartPeriodeYear int                      `json:"start_periode_year"`
	EndPeriodeYear   int                      `json:"end_periode_year"`
	CreatedAt        time.Time                `json:"created_at"`
	UpdatedAt        time.Time                `json:"updated_at"`
}

// public consume
type PengurusPublicResponse struct {
	ID               int                      `json:"id"`
	PhotoURL         string                   `json:"photo_url"`
	Divisi           string                   `json:"divisi"`
	Name             string                   `json:"name"`
	Position         string                   `json:"position"`
	Sosmed           []PengurusSosmedResponse `json:"sosmed"`
	StartPeriodeYear int                      `json:"start_periode_year"`
	EndPeriodeYear   int                      `json:"end_periode_year"`
}

// Untuk create/register pengurus
type RegisterPengurusRequest struct {
	UserID           int                   `form:"id_user"`
	PhotoURL         *multipart.FileHeader `form:"photo_url"`
	Email            string                `form:"email"`
	Divisi           string                `form:"divisi" binding:"required,divisi"`
	Name             string                `form:"name" binding:"required" validate:"min=2,max=150"`
	Position         string                `form:"position" binding:"required,position"`
	Sosmed           []string              `form:"sosmed" binding:"omitempty,max=4,dive,socialurl"`
	StartPeriodeYear int                   `form:"start_periode_year" binding:"required" validate:"max=50"`
	EndPeriodeYear   int                   `form:"end_periode_year" binding:"required" validate:"max=50"`
}

type PengurusPayload struct {
	Email            string                `form:"email" binding:"omitempty,email"`
	Divisi           string                `form:"divisi" binding:"omitempty,divisi"`
	Name             string                `form:"name" binding:"omitempty"`
	StartPeriodeYear int                   `form:"start_periode_year" binding:"omitempty"`
	EndPeriodeYear   int                   `form:"end_periode_year" binding:"omitempty"`
	Position         string                `form:"position" binding:"omitempty,position"`
	PhotoURL         *multipart.FileHeader `form:"file" binding:"omitempty"`
}

// Untuk update/patch pengurus
type PengurusPatch struct {
	Email            string                `form:"email" binding:"omitempty,email"`
	Divisi           string                `form:"divisi" binding:"omitempty,divisi"`
	Name             string                `form:"name" binding:"omitempty"`
	Sosmed           []string              `form:"sosmed" binding:"omitempty,max=4,dive,socialurl"`
	StartPeriodeYear int                   `form:"start_periode_year" binding:"omitempty"`
	EndPeriodeYear   int                   `form:"end_periode_year" binding:"omitempty"`
	Position         string                `form:"position" binding:"omitempty,position"`
	PhotoURL         *multipart.FileHeader `form:"file" binding:"omitempty"`
}

// struct to scan data from database
type PengurusRow struct {
	ID               int             `json:"id"`
	IDUser           int             `json:"id_user"`
	PhotoURL         string          `json:"photo_url"`
	Email            string          `json:"email"`
	Divisi           string          `json:"divisi"`
	Name             string          `json:"name"`
	Position         string          `json:"position"`
	Sosmed           json.RawMessage `json:"sosmed"`
	StartPeriodeYear int             `form:"start_periode_year"`
	EndPeriodeYear   int             `form:"end_periode_year"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}
