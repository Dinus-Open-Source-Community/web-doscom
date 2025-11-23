package model

import (
	"fmt"
	"time"

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
	Password  string    `json:"-"` // stored hashed; omitted from JSON responses
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

	if u.Username != nil {
		result["username"] = *u.Username
	}

	if u.Email != nil {
		result["email"] = *u.Email
	}

	if u.Fullname != nil {
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
func (m *UserModel) GetAllUser(role string, userRole string) ([]User, error) {
	var users []User

	query := m.DB.Model(&User{})

	switch role {
	case "Super_Admin":
		query = query.Where("role != ?", role)
	case "BPH_ang":
		query = query.Where("role = ?", userRole)
	default:
		query = query.Where("role = ?", userRole)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// update
func (m *UserModel) UpdateUser(Id int, patch map[string]any) (*User, error) {
	var user User
	// find user by id
	if err := m.DB.First(&user, Id).Error; err != nil {
		return nil, err
	}

	// set allowed fields to update
	allowedField := map[string]bool{
		"username":  true,
		"email":     true,
		"full_name": true,
	}

	// compare the data and filter the empty value
	filteredUpdates := make(map[string]any)
	for field, value := range patch {
		if allowedField[field] && value != nil {
			if str, ok := value.(string); ok {
				if str != "" {
					filteredUpdates[field] = value
				}
			} else {
				filteredUpdates[field] = value
			}
		}
	}
	// debug bwang
	fmt.Println("Filtered updates:", filteredUpdates)

	if len(filteredUpdates) == 0 {
		return &user, nil
	}

	// update data
	if err := m.DB.Model(&user).Updates(filteredUpdates).Error; err != nil {
		return nil, err
	}

	// make sure the data is updated
	if err := m.DB.First(&user, Id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// delete
func (m *UserModel) Deleteuser(id int) error {
	var user User

	// take user by id
	if err := m.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return err
	}

	if err := m.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
