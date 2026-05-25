package dto

import "strings"

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

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type AdminChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

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
