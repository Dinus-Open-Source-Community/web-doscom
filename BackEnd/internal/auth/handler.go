package auth

import (
	"log"
	"net/http"

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
			"error":    "Invalid  password",
			"pw":       user.Password,
			"pw input": input.Password,
		})
		return
	}

	// check the role
	if user.Role != "Admin" && user.Role != "Super_Admin" {
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

	// send back
	c.JSON(http.StatusOK, gin.H{
		"message:": "login success bolo, nasi padang satu bungkus",
		"token:":   token,
	})
}

// RegisterHandler godoc
// @Summary      Admin register
// @Description  register untuk user dengan role Admin atau Super_Admin
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login  body  model.LoginRequest  true  "Login credentials"
// @Success      200  {object}  map[string]string  "Login berhasil, kembalikan token JWT"
// @Failure      400  {object}  map[string]string  "Request body invalid"
// @Failure      401  {object}  map[string]string  "Email/password salah atau akses ditolak"
// @Failure      500  {object}  map[string]string  "Error saat membuat token"
// @Router       /api/v1/register [post]
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
	log.Printf("Role value: '%s'", input.Role)

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
