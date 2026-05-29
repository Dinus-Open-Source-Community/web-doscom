package entity

import (
	"context"
	"fmt"
	"time"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"

	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Full_name string    `json:"full_name"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// insert new user data
func (m *UserModel) InsertUser(user *User) error {
	result := m.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no row inserted")
	}

	return nil
}

// get user by email
func (m *UserModel) FindByEmail(email string) (*User, error) {
	var user User
	result := m.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, fmt.Errorf("terjadi kesalahan ketika ambil data %w", result.Error)
	}

	return &user, nil

}

// get user by id
func (m *UserModel) GetUserById(id int) (*User, error) {
	var user User
	if err := m.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// get all user data
func (m *UserModel) GetAllUserBaseOnRole(currentUserRole string) ([]dto.UserResponse, error) {
	var usersData []dto.UserResponse

	query := m.DB.Model(&User{}).Select("id, username, email, role, full_name")

	switch currentUserRole {
	case constants.RoleKeySuperAdmin:
		query = query.Where("role != ?", currentUserRole)
	default:
		query = query.Where("role = ?", currentUserRole)
	}

	if err := query.Scan(&usersData).Error; err != nil {
		return nil, err
	}

	return usersData, nil
}

func (m *UserModel) GetSuperAdmin(ctx context.Context, userRole string, userID int) ([]dto.UserResponse, error) {
	var user []dto.UserResponse

	if err := m.DB.Model(&User{}).
		Select("id, username, email, role, full_name").
		Where("role = ?", userRole).
		Where("id <> ?", userID).
		Scan(&user).
		Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *UserModel) UpdateUser(Id int, dataToUpdate map[string]any) (*dto.UserResponse, error) {
	var user User
	if err := m.DB.First(&user, Id).Error; err != nil {
		return nil, fmt.Errorf("user tidak di temukan %w", err)
	}

	if err := m.DB.Model(&user).Updates(dataToUpdate).Error; err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		Id:        Id,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Full_name: user.Full_name,
	}, nil
}

// delete
func (m *UserModel) DeleteUser(id int) error {

	result := m.DB.Delete(&User{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete data %w", result.Error)
	}
	if result.RowsAffected == 0 {
		// log.Printf("[entity] error nya disini")
		return fmt.Errorf("user nya ngga ada wlee :|>  %w", gorm.ErrRecordNotFound)
	}

	return nil
}
