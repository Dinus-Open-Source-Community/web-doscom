package model

import (
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

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

func (m *UserModel) InsertUser(user *User) error {
	return m.DB.Create(user).Error
}
