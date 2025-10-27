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
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Password   string    `json:"password"`
	Full_name  string    `json:"full_name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
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
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required"`
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
func (m *UserModel) GetAllUser() ([]User, error) {
	var users []User
	if err := m.DB.Where("role != ?", "Super_Admin").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// update
func (m *UserModel) UpdateUser(Id int, patch UserPatch) (*User, error) {
	var user User
	// find user by id
	if err := m.DB.First(&user, Id).Error; err != nil {
		return nil, err // user tidak ditemukan
	}

	// compare the data
	pathData := UserPatch{
		Username: patch.Username,
		Email:    patch.Email,
		Fullname: patch.Fullname,
	}

	if err := m.DB.Model(&user).Updates(pathData).Error; err != nil {
		return nil, err
	}

	// reload data -> memastikan data terbarui
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
