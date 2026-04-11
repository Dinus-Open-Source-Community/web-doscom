package model

import (
	"time"

	"gorm.io/gorm"
)

type WorkModel struct {
	DB *gorm.DB
}

type Work struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	GalleryID   int       `gorm:"column:gallery_id" json:"gallery_id"`
	Description string    `json:"description"`
	ProjectDate time.Time `gorm:"column:project_date" json:"project_date"`
	TeamProject int       `gorm:"column:pengurus_id" json:"pengurus_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegisterWork struct {
	Title       string    `json:"title" binding:"required"`
	GalleryID   int       `json:"gallery_id" binding:"required"`
	Description string    `json:"description" binding:"required"`
	ProjectDate time.Time `json:"project_date" binding:"required"`
	TeamProject int       `json:"pengurus_id" binding:"required"`
}

type WorkResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	GalleryID   int       `json:"gallery_id"`
	Description string    `json:"description"`
	ProjectDate time.Time `json:"project_date"`
	TeamProject int       `json:"pengurus_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type WorkPatch struct {
	Title       *string    `json:"title" binding:"omitempty"`
	Description *string    `json:"description" binding:"omitempty"`
	StartDate   *time.Time `json:"start_date" binding:"omitempty"`
	EndDate     *time.Time `json:"end_date" binding:"omitempty"`
}

func (Work) TableName() string {
	return "work"
}

// CRUD implementation

func (m *WorkModel) InsertWork(work *Work) error {
	work.CreatedAt = time.Now()
	work.UpdatedAt = time.Now()
	return m.DB.Create(work).Error
}

func (m *WorkModel) GetWorkById(id int) (*Work, error) {
	var work Work
	if err := m.DB.First(&work, id).Error; err != nil {
		return nil, err
	}
	return &work, nil
}

func (m *WorkModel) GetAllWorks() ([]Work, error) {
	var works []Work
	if err := m.DB.Order("created_at DESC").Find(&works).Error; err != nil {
		return nil, err
	}
	return works, nil
}

func (m *WorkModel) UpdateWork(id int, patch map[string]any) (*Work, error) {
	var existingWork Work
	if err := m.DB.First(&existingWork, id).Error; err != nil {
		return nil, err
	}

	// set allowed fields to update
	allowedFields := map[string]bool{
		"title":       true,
		"description": true,
		"StartDate":   true,
		"EndDate":     true,
	}

	// compare the data and filter the empty value
	filteredUpdates := make(map[string]any)
	for field, value := range patch {
		if allowedFields[field] && value != nil {
			if str, ok := value.(string); ok {
				if str != "" {
					filteredUpdates[field] = value
				}
			} else {
				filteredUpdates[field] = value
			}
		}
	}

	// check if there is any update
	if len(filteredUpdates) == 0 {
		return &existingWork, nil
	}

	// update
	if err := m.DB.Model(&existingWork).Updates(filteredUpdates).Error; err != nil {
		return nil, err
	}

	// reload data -> make sure the data is updated
	if err := m.DB.First(&existingWork, id).Error; err != nil {
		return nil, err
	}

	return &existingWork, nil

}

func (m *WorkModel) DeleteWork(id int) error {
	var work Work
	if err := m.DB.First(&work, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
		return err
	}

	if err := m.DB.Delete(&work).Error; err != nil {
		return err
	}

	return nil
}
