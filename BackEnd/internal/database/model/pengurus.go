package model

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Validasi sesuai SQL
var ValidDivisi = map[string]bool{
	"bph":      true,
	"pemro":    true,
	"jaringan": true,
	"medcrev":  true,
	"data":     true,
}

var ValidPosition = map[string]bool{
	"ketum":        true,
	"sdm":          true,
	"pr":           true,
	"pm":           true,
	"pmAng":        true,
	"KoorPemro":    true,
	"KoorJaringan": true,
	"KoorMedcrev":  true,
	"KoorData":     true,
	"sekum":        true,
	"bendum":       true,
	"sekAng":       true,
	"bendAng":      true,
	"PemroAng":     true,
	"JaringanAng":  true,
	"MedcrevAng":   true,
	"DataAng":      true,
}

// "BPH":          true,
var ValidSosmed = map[string]bool{
	"instagram": true,
	"linkedin":  true,
	"github":    true,
}

type PengurusModel struct {
	DB *gorm.DB
}

// Model GORM
type Pengurus struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `gorm:"column:id_user" json:"id_user"`
	PhotoURL  string    `gorm:"column:photo_url" json:"url_asset"`
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Divisi    string    `gorm:"column:divisi" json:"divisi"`
	Position  string    `json:"position"`
	Sosmed    string    `json:"sosmed"`
	Period    string    `json:"period"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// DTO untuk frontend
// UserID   int       `json:"id_user"`
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
	Name     string                `form:"position" binding:"omitempty,position"`
	Sosmed   string                `form:"sosmed" binding:"omitempty,sosmed"`
	Period   string                `form:"period" binding:"omitempty" validate:"max=50"`
	Position string                `form:"position" binding:"omitempty,position"` // condition just for koor and super_admin
	PhotoURL *multipart.FileHeader `form:"url_asset" binding:"omitempty"`
}

// Custom validator
func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("divisi", func(fl validator.FieldLevel) bool {
			return ValidDivisi[fl.Field().String()]
		})
		v.RegisterValidation("position", func(fl validator.FieldLevel) bool {
			return ValidPosition[fl.Field().String()]
		})
		v.RegisterValidation("sosmed", func(fl validator.FieldLevel) bool {
			val := fl.Field().String()
			return val == "" || ValidSosmed[val]
		})
	}
}

func (Pengurus) TableName() string {
	return "pengurus"
}

// User model stub for role validation (should be replaced with your User struct)
type UserGet struct {
	ID    int
	Role  string
	Email string
}

// Insert new pengurus data
func (m *PengurusModel) InsertPengurus(pengurus *Pengurus) error {
	// Validasi user_id ke tabel users (pseudo, harus ada model User)
	var user User
	if err := m.DB.First(&user, pengurus.UserID).Error; err != nil {
		return fmt.Errorf("user_id tidak ditemukan")
	}
	// Validasi email unik
	var count int64
	m.DB.Model(&Pengurus{}).Where("email = ?", pengurus.Email).Count(&count)
	if count > 0 {
		return fmt.Errorf("email sudah terdaftar di pengurus")
	}

	pengurus.CreatedAt = time.Now()
	result := m.DB.Create(pengurus)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no row inserted")
	}
	return nil
}

// Get pengurus by email
func (m *PengurusModel) FindByEmail(email string) (*Pengurus, error) {
	var pengurus Pengurus
	result := m.DB.First(&pengurus, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pengurus, nil
}

// Get pengurus by id
func (m *PengurusModel) GetPengurusById(id int) (*Pengurus, error) {
	var pengurus Pengurus
	if err := m.DB.First(&pengurus, id).Error; err != nil {
		return nil, err
	}
	return &pengurus, nil
}

// Get all pengurus data
func (m *PengurusModel) GetAllPengurusByDivisi(divisi string) ([]Pengurus, error) {
	var pengurus []Pengurus
	db := m.DB
	if divisi != "" {
		db = db.Where("divisi = ?", divisi)
	}
	if err := db.Find(&pengurus).Error; err != nil {
		return nil, err
	}
	return pengurus, nil
}

// Update pengurus
func (m *PengurusModel) UpdatePengurus(Id int, patch PengurusPatch) (*Pengurus, error) {
	var pengurus Pengurus
	// find pengurus by id
	if err := m.DB.First(&pengurus, Id).Error; err != nil {
		return nil, err
	}
	if err := m.DB.Model(&pengurus).Updates(patch).Error; err != nil {
		return nil, err
	}
	// reload data
	if err := m.DB.First(&pengurus, Id).Error; err != nil {
		return nil, err
	}
	return &pengurus, nil
}

func (m *PengurusModel) UpdatePengurusPartial(id int, data map[string]any) (*PengurusResponse, error) {
	var updatePengurus Pengurus

	if err := m.DB.Model(&updatePengurus).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return nil, err
	}

	return &PengurusResponse{
		ID:       updatePengurus.ID,
		PhotoURL: updatePengurus.PhotoURL,
		Email:    updatePengurus.Email,
		Divisi:   updatePengurus.Divisi,
		Name:     updatePengurus.Name,
		Position: updatePengurus.Position,
		Sosmed:   updatePengurus.Sosmed,
		Period:   updatePengurus.Period,
	}, nil
}

// Delete pengurus
func (m *PengurusModel) DeletePengurus(ctx context.Context, id int) error {
	result := m.DB.WithContext(ctx).Delete(&Pengurus{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete data %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (m *PengurusModel) GetPengurusByDivisi(ctx context.Context, division string) ([]PengurusResponse, error) {
	var dataPengurus []PengurusResponse

	if division == "" {
		return nil, fmt.Errorf("division required")
	}

	if err := m.DB.WithContext(ctx).
		Model(&Pengurus{}).
		Select("id, photo_url, email, divisi, name, position, sosmed,  period").
		Where("divisi = ?", division).
		Order("name ASC").
		Scan(&dataPengurus).
		Error; err != nil {
		return nil, fmt.Errorf("terjadi error ketika ambil data wlee %w", err)
	}

	return dataPengurus, nil
}
