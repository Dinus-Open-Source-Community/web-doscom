package model

import (
	"fmt"
	"strings"
	"time"
	"web_doscom/internal/constants"

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

// dto for front end -> data that send to frontend
type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Full_name string `json:"full_name"`
}

// untuk register/create user method
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Fullname string `json:"fullname" binding:"required"`
}

// untuk login method
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// untuk kebutuhan update data
type UserPatch struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Fullname *string `json:"fullname"`
}

type DefaultValue struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// method struct pada UserPatch
func (u *UserPatch) ToMap() map[string]any {
	result := make(map[string]any)

	if u.Username != nil && strings.TrimSpace(*u.Username) != "" {
		result["username"] = *u.Username
	}

	if u.Email != nil && strings.TrimSpace(*u.Email) != "" {
		result["email"] = *u.Email
	}

	if u.Fullname != nil && strings.TrimSpace(*u.Fullname) != "" {
		result["full_name"] = *u.Fullname
	}

	return result
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
		return nil, result.Error
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
func (m *UserModel) GetAllUserBaseOnRole(userRoleToget string) ([]UserResponse, error) {
	var usersData []UserResponse

	query := m.DB.Model(&User{}).Select("id, username, email, role, full_name")

	switch {
	case userRoleToget == constants.RoleKeySuperAdmin:
		query = query.Where("role != ?", userRoleToget)
	default:
		query = query.Where("role = ?", userRoleToget)
	}

	if err := query.Scan(&usersData).Error; err != nil {
		return nil, err
	}

	return usersData, nil
}

func (m *UserModel) UpdateUser(Id int, dataToUpdate map[string]any) (*UserResponse, error) {
	var user User
	if err := m.DB.First(&user, Id).Error; err != nil {
		return nil, fmt.Errorf("user tidak di temukan %w", err)
	}

	if err := m.DB.Model(&user).Updates(dataToUpdate).Error; err != nil {
		return nil, err
	}

	return &UserResponse{
		Id,
		user.Username,
		user.Email,
		user.Role,
		user.Full_name,
	}, nil
}

// delete
func (m *UserModel) DeleteUser(id int) error {

	result := m.DB.Delete(&User{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete data %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found %w", gorm.ErrRecordNotFound)
	}

	return nil
}
