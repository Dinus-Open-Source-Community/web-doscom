package entity

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"web_doscom/internal/database/model/dto"

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
	ID               int       `gorm:"primaryKey" json:"id"`
	IDUser           int       `gorm:"column:id_user" json:"id_user"`
	PhotoURL         string    `gorm:"column:photo_url" json:"url_asset"`
	Name             string    `json:"name"`
	Email            string    `gorm:"uniqueIndex" json:"email"`
	Divisi           string    `gorm:"column:divisi" json:"divisi"`
	Position         string    `json:"position"`
	StartPeriodeYear int       `json:"start_periode_year"`
	EndPeriodeYear   int       `json:"end_periode_year"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
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
	log.Printf("user id %d", pengurus.IDUser)
	if err := m.DB.First(&user, pengurus.IDUser).Error; err != nil {
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

// Get pengurus by user id
func (m *PengurusModel) GetPengurusByUserID(ctx context.Context, id int) (*dto.PengurusResponse, error) {
	var pengurusRow dto.PengurusRow

	query := `
		SELECT
			p.id,
			p.id_user,
			p.photo_url,
			p.name,
			p.email,
			p.divisi,
			p.position,
			p.start_periode_year,
			p.end_periode_year,
			p.created_at,
			p.updated_at,
			COALESCE(sosmed.urls, '[]'::json) AS sosmed
		FROM pengurus p
		LEFT JOIN LATERAL (
			SELECT json_agg(
				json_build_object(
					'platform', ps.platform,
					'username', ps.username,
					'url', ps.url,
					'is_primary', ps.is_primary
				)
			) AS urls FROM pengurus_sosmed ps WHERE ps.pengurus_id = p.id
		)sosmed ON true WHERE p.id_user = $1
	`

	if err := m.DB.WithContext(ctx).Raw(query, id).Scan(&pengurusRow).Error; err != nil {
		return nil, fmt.Errorf("error while getting the data: %w", err)
	}

	var sosmed []dto.PengurusSosmedResponse
	if err := json.Unmarshal(pengurusRow.Sosmed, &sosmed); err != nil {
		return nil, fmt.Errorf("error while unmarshalling sosmed: %w", err)
	}

	return &dto.PengurusResponse{
		ID:               pengurusRow.ID,
		IDUser:           pengurusRow.IDUser,
		PhotoURL:         pengurusRow.PhotoURL,
		Email:            pengurusRow.Email,
		Divisi:           pengurusRow.Divisi,
		Name:             pengurusRow.Name,
		Position:         pengurusRow.Position,
		Sosmed:           sosmed,
		StartPeriodeYear: pengurusRow.StartPeriodeYear,
		EndPeriodeYear:   pengurusRow.EndPeriodeYear,
		CreatedAt:        pengurusRow.CreatedAt,
		UpdatedAt:        pengurusRow.UpdatedAt,
	}, nil
}

// Get all pengurus data
func (m *PengurusModel) GetAllPengurusByDivisi(ctx context.Context, divisi string) ([]Pengurus, error) {
	pengurus := []Pengurus{}
	db := m.DB.WithContext(ctx)
	if divisi != "" {
		db = db.Where("divisi = ?", divisi)
	}
	if err := db.Find(&pengurus).Error; err != nil {
		return nil, err
	}
	return pengurus, nil
}

// Update pengurus
func (m *PengurusModel) UpdatePengurus(Id int, patch dto.PengurusPatch) (*Pengurus, error) {
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

func (m *PengurusModel) UpdatePengurusPartial(id int, data map[string]any) (*dto.PengurusResponse, error) {
	var updatePengurus Pengurus

	result := m.DB.Model(&updatePengurus).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(data)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &dto.PengurusResponse{
		ID:               updatePengurus.ID,
		PhotoURL:         updatePengurus.PhotoURL,
		Email:            updatePengurus.Email,
		Divisi:           updatePengurus.Divisi,
		Name:             updatePengurus.Name,
		Position:         updatePengurus.Position,
		StartPeriodeYear: updatePengurus.StartPeriodeYear,
		EndPeriodeYear:   updatePengurus.EndPeriodeYear,
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

func (m *PengurusModel) GetPengurusByDivisi(ctx context.Context, division string) ([]dto.PengurusResponse, error) {
	dataPengurus := []dto.PengurusResponse{}

	if division == "" {
		return nil, fmt.Errorf("division required")
	}

	if err := m.DB.WithContext(ctx).
		Model(&Pengurus{}).
		Select("id, photo_url, email, divisi, name, position,  start_periode_year, end_periode_year").
		Where("divisi = ?", division).
		Order("name ASC").
		Scan(&dataPengurus).
		Error; err != nil {
		return nil, fmt.Errorf("terjadi error ketika ambil data wlee %w", err)
	}

	return dataPengurus, nil
}
