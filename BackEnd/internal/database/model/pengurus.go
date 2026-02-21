package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
	"pm_ang":       true,
	"sekum":        true,
	"bendum":       true,
	"sek_ang":      true,
	"ben_ang":      true,
	"kor_pemro":    true,
	"kor_jaringan": true,
	"kor_medcrev":  true,
	"kor_data":     true,
	"anggota":      true,
	"pemro_ang":    true,
	"jaringan_ang": true,
	"medcrev_ang":  true,
	"data_ang":     true,
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
	URLAsset  string    `gorm:"column:urlasset" json:"url_asset"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Divisi    string    `gorm:"column:divisi" json:"divisi"`
	Name      string    `json:"name"`
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
	URLAsset string `json:"url_asset"`
	Email    string `json:"email"`
	Divisi   string `json:"divisi"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Sosmed   string `json:"sosmed"`
	Period   string `json:"period"`
}

// Untuk create/register pengurus
// type RegisterPengurusRequest struct {
// 	UserID   int    `form:"id_user"`
// 	URLAsset string `form:"url_asset"`
// 	Email    string `form:"email"`
// 	Divisi   string `form:"divisi" binding:"required,divisi"`
// 	Name     string `form:"name" binding:"required" validate:"min=2,max=150"`
// 	Position string `form:"position" binding:"required,position"`
// 	Sosmed   string `form:"sosmed" binding:"omitempty,socialurl"`
// 	Period   string `form:"period" binding:"required" validate:"max=50"`
// }

type RegisterPengurusRequest struct {
	UserID   int    `form:"id_user"`
	URLAsset string `form:"url_asset"`
	Email    string `form:"email"`
	Divisi   string `form:"divisi" binding:"required,divisi"`
	Name     string `form:"name" binding:"required" validate:"min=2,max=150"`
	Position string `form:"position" binding:"required,position"`
	Sosmed   string `form:"sosmed" binding:"omitempty,sosmed"`
	Period   string `form:"period" binding:"required" validate:"max=50"`
}

// Untuk update/patch pengurus
type PengurusPatch struct {
	URLAsset *string `json:"url_asset" binding:"omitempty"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Divisi   *string `json:"divisi" binding:"omitempty,divisi"`
	Name     *string `json:"name" binding:"omitempty" validate:"min=2,max=150"`
	Position *string `json:"position" binding:"omitempty,position"`
	Sosmed   *string `json:"sosmed" binding:"omitempty,sosmed"`
	Period   *string `json:"period" binding:"omitempty" validate:"max=50"`
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

// Delete pengurus
func (m *PengurusModel) DeletePengurus(id int) error {
	var pengurus Pengurus
	if err := m.DB.First(&pengurus, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}
	if err := m.DB.Delete(&pengurus).Error; err != nil {
		return err
	}
	return nil
}
