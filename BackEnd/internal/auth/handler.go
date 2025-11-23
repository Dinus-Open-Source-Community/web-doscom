package auth

import (
	"net/http"
	"slices"
	"time"

	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Model *model.UserModel
}

func NewUserauth(m *model.UserModel) *AuthHandler {
	return &AuthHandler{Model: m}
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
	user, err := h.Model.FindByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email ",
		})
		return
	}

	// veerify the password
	if !verifyPassword(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid  password",
		})
		return
	}

	// check the role
	allowedRoles := []string{"Super_Admin", "Kor_Pemro", "Kor_Jaringan", "Kor_Data", "Kor_Medcrev", "BPH", "pemro_ang", "jaringan_ang", "medcrev_ang", "data_ang", "BPH_ang"}

	if !slices.Contains(allowedRoles, user.Role) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Acces denied, users have no acces",
		})
		return
	}

	// generate token jwt
	token, err := Create_token(user.ID, user.Email, user.Username, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error: ": "Failed to create token",
		})

		return
	}

	// set cookie
	SetCustomCookie(c, Cookies{
		Name:     "AccessToken",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(1 * time.Hour),
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
		"message": "Login success, nasi padangnya sebungkus bolo",
		"token":   token,
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
	if err := h.Model.InsertUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
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
		Name:     "AccessToken",
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
