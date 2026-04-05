package auth

import (
	"net/http"
	"time"

	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"

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
	var input model.LoginRequest
	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})

		return
	}

	// look at the requested user
	userData, err := h.Auth.FindByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password ",
		})
		return
	}

	// veerify the password
	if !verifyPassword(input.Password, userData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// // check the role

	_, err = constants.GetRoleInfo(userData.Role)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "acces denied",
		})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error: ": "Failed to create token",
		})

		return
	}

	refreshToken, tokenHash, err := generateRefreshToken(userData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error: ": "Failed to create refresh token",
		})
		return
	}
	// insert to database
	if err := h.Auth.CreateRefreshToken(refreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error: ": "Failed to insert refresh token",
		})

		return
	}

	// set cookie
	SetCustomCookie(c, Cookies{
		Name:     "RefreshToken",
		Value:    refreshToken.Token,
		Path:     "/",
		Expires:  time.Now().Add(5 * 24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	// production response
	// send back
	// c.JSON(http.StatusOK, gin.H{
	// "message:": "login success bolo, nasi padang satu bungkus",
	// })

	// development response
	c.JSON(http.StatusOK, gin.H{
		"message":       "Login success, nasi padangnya sebungkus bolo",
		"token":         accesToken,
		"refresh_token": tokenHash,
	})
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {

	// get the request body
	var input model.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body or missing fields"})
		return
	}

	if input.Username == "" || input.Email == "" || input.Password == "" || input.Role == "" || input.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields (username, email, password, role, fullname) are required"})
		return
	}

	if len(input.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters"})
		return
	}

	// hash the password
	passwordHash := HashPassword(input.Password)

	// mapping data user
	user := model.User{
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
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})

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

	// delete cookie
	SetCustomCookie(c, Cookies{
		Name:     "RefreshToken",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Success, nasi padang satu bungkus",
	})
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "Refresh token not found",
		})
		return
	}

	newAccessToken, err := h.Auth.validateRefreshToken(refreshCookie)
	if err != nil {
		// clear expired Cookie
		SetCustomCookie(c, Cookies{
			Name:     "RefreshToken",
			Value:    "",
			Path:     "/",
			Expires:  time.Now().Add(-time.Hour),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})

		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "refresh token invalid or expired",
		})

		return
	}

	// production response
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "refresh token success created",
	// })

	// development response
	c.JSON(http.StatusOK, gin.H{
		"message": "refresh token success",
		"token":   newAccessToken,
	})

}
