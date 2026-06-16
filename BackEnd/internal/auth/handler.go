package auth

import (
	"net/http"
	"time"

	"web_doscom/internal/authorization"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Auth *AuthService
}

func NewUserauth(a *AuthService) *AuthHandler {
	return &AuthHandler{Auth: a}
}

// LoginHandler godoc
// @Summary      Admin login
// @Description  Login untuk user dengan role Admin atau Super_Admin
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login  body  model.LoginRequest  true  "Login credentials"
// @Success      200  {object}  map[string]string  "Login berhasil, kembalikan token JWT"
// @Failure      400  {object}  map[string]string  "Request body invalid"
// @Failure      401  {object}  map[string]string  "Email/password salah atau akses ditolak"
// @Failure      500  {object}  map[string]string  "Error saat membuat token"
// @Router       /api/v1/login [post]
func (h *AuthHandler) LoginHandler(c *gin.Context) {

	// get the email and password from the req body
	var input dto.LoginRequest
	if c.Bind(&input) != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Failed to read req body",
			nil,
		)
		return
	}

	// look at the requested user
	userData, err := h.Auth.FindByEmail(input.Email)
	if err != nil {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid email or password",
			nil,
		)
		return
	}

	// veerify the password
	if !VerifyPassword(input.Password, userData.Password) {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Invalid email or password",
			nil,
		)
		return
	}

	// check the role
	_, err = authorization.GetRoleInfo(userData.Role)
	if err != nil {
		response.Error(
			c,
			http.StatusUnauthorized,
			"acces denied",
			nil,
		)
		return
	}
	// generate token jwt
	accesToken, err := Create_token(
		userData.ID,
		userData.Email,
		userData.Username,
		userData.Role,
	)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to create token",
			nil,
		)
		return
	}

	refreshToken, plainToken, err := generateRefreshToken(userData.ID)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to create refresh token",
			nil,
		)
		return
	}

	// Hapus token lama dulu agar tidak duplicate key error
	h.Auth.RefreshToken.DeleteRefreshTokenByUserId(userData.ID)

	// insert to database
	if err := h.Auth.CreateRefreshToken(refreshToken); err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to insert refresh token",
			nil,
		)
		return
	}

	// set refresh token cookie
	SetCustomCookie(c, Cookies{
		Name:     "RefreshToken",
		Value:    plainToken,
		Path:     "/api/v1/auth",
		Expires:  time.Now().Add(2 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	// set access token cookie
	SetCustomCookie(c, Cookies{
		Name:     "AccessToken",
		Value:    accesToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	response.Success(
		c,
		"login success bolo, nasi padang satu bungkus",
		http.StatusOK,
		nil,
	)

}

func (h *AuthHandler) RegisterUser(c *gin.Context) {

	// get the request body
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid request body or missing fields",
			nil,
		)
		return
	}

	if input.Username == "" || input.Email == "" || input.Password == "" || input.Role == "" || input.Fullname == "" {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid request body or missing fields",
			nil,
		)
		return
	}

	if len(input.Password) < 8 {
		response.Error(
			c,
			http.StatusBadRequest,
			"Password must be at least 8 characters",
			nil,
		)
		return
	}

	// hash the password
	passwordHash := HashPassword(input.Password)

	// mapping data user
	user := entity.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  passwordHash,
		Role:      input.Role,
		Full_name: input.Fullname,
	}

	// insert to database
	if err := h.Auth.InsertUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to register user",
		})
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to register user",
			nil,
		)
		return
	}

	response.Success(
		c,
		"user created successfully",
		http.StatusOK,
		nil,
	)

}

// Logout godoc
// @Summary      Logout user
// @Description  Logout user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {

	// delete refresh token from database
	refreshToken, err := c.Cookie("RefreshToken")
	if err != nil {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Cookie not found, what are you doing here????",
			nil,
		)
		return
	}
	if err := h.Auth.DeleteRefreshToken(refreshToken); err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"there was an error deleting the refresh token",
			nil,
		)
		return
	}
	// delete Cookie
	SetCustomCookie(c, Cookies{
		Name:     "RefreshToken",
		Value:    "",
		Path:     "/api/v1/auth",
		Expires:  time.Now().Add(-time.Hour),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	SetCustomCookie(c, Cookies{
		Name:     "AccessToken",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Minute),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	response.Success(
		c,
		"Logout Success, nasi padang satu bungkus",
		http.StatusOK,
		nil,
	)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Generate new access token using refresh token stored in HTTP-only cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} model.RefreshTokenSuccessResponse
// @Failure 401 {object} model.RefreshTokenErrorResponse
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshCookie, err := c.Cookie("RefreshToken")
	if err != nil {
		response.Error(
			c,
			http.StatusUnauthorized,
			"Refresh token not found, what are you doing heree????",
			nil,
		)
		return
	}

	newAccessToken, err := h.Auth.validateRefreshToken(refreshCookie)
	if err != nil {
		// clear expired Cookie
		SetCustomCookie(c, Cookies{
			Name:     "RefreshToken",
			Value:    "",
			Path:     "/api/v1/auth",
			Expires:  time.Now().Add(-time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "refresh token invalid or expired",
		})
		response.Error(
			c,
			http.StatusUnauthorized,
			"refresh token invalid or expired",
			nil,
		)
		return
	}
	// set new cookie access token
	SetCustomCookie(c, Cookies{
		Name:     "AccessToken",
		Value:    newAccessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	response.Success(
		c,
		"refresh token success",
		http.StatusOK,
		nil,
	)

}
