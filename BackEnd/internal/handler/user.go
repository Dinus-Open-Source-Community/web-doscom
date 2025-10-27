package handler

import (
	"log"
	"net/http"
	"strconv"
	"web_doscom/internal/auth"
	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	Model *model.UserModel
}

func NewUserHandler(m *model.UserModel) *UserHandler {
	return &UserHandler{Model: m}
}

// CreateUserHandler untuk POST /api/user
// CreateUser godoc
// @Summary      Create new user
// @Description  Membuat user baru
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body  model.RegisterRequest  true  "User info"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/users [post]
func (m *UserHandler) CreateUser(c *gin.Context) {

	var input model.RegisterRequest

	// get the req body
	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})

		return
	}

	// hash the password
	passwordHash := auth.HashPassword(input.Password)
	log.Printf("Role value: '%s'", input.Role)

	// mapping the user
	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		Role:      input.Role,
		Password:  passwordHash,
		Full_name: input.Fullname,
	}

	// insert to database
	if err := m.Model.InsertUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})

}

// CreateSuperAdmin godoc
// @Summary      Create superadmin user
// @Description  Membuat user dengan role superadmin (hanya untuk admin/superadmin)
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body  model.RegisterRequest  true  "User info"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      409  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/users/superadmin [post]
func (m *UserHandler) CreateSuperAdmin(c *gin.Context) {
	var input model.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read req body"})
		return
	}
	if input.Email == "" || input.Password == "" || input.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, email, and password are required"})
		return
	}
	if len(input.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
		return
	}

	// Check email uniqueness
	if _, err := m.Model.FindByEmail(input.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	// Hash password
	passwordHash := auth.HashPassword(input.Password)

	user := &model.User{
		Username: input.Fullname,
		Email:    input.Email,
		Password: passwordHash,
		Role:     "superadmin",
	}

	if err := m.Model.InsertUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Superadmin created successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	})
}

// GetUserHandler untuk GET /api/user/:id
// GetUser godoc
// @Summary      Get user by ID
// @Description  Mendapatkan user berdasarkan ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.UserResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Router      /api/v1/users/{id} [get]
func (m *UserHandler) GetUser(c *gin.Context) {
	idParams := c.Param("id")

	// konversi id dari string ke int
	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	user, err := m.Model.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User Not Found",
		})
	}

	userConsum := model.UserResponse{
		Id:        int(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Full_name: user.Full_name,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get user",
		"user":    userConsum,
	})
}

// get all user
// GetAllUser godoc
// @Summary      Get all users
// @Description  Mendapatkan daftar semua user
// @Tags         Users
// @Produce      json
// @Success      200  {array}   model.UserResponse
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/users [get]
func (m *UserHandler) GetAllUser(c *gin.Context) {
	// get all user
	users, err := m.Model.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users data",
		})
	}

	var userresponse []model.UserResponse
	for _, u := range users {
		userresponse = append(userresponse, model.UserResponse{
			Id:        int(u.ID),
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
			Full_name: u.Full_name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "List of users (excluding superadmin)",
		"users":   userresponse,
	})
}

// update admin by id
// UpdateUser godoc
// @Summary      Update user
// @Description  Mengupdate data user berdasarkan ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "User ID"
// @Param        user  body      model.UserPatch  true  "User info"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/users/{id} [put]
func (m *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}

	// req body data new user
	var patchUser model.UserPatch
	if err := c.ShouldBindJSON(&patchUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	// update the data
	updateuser, err := m.Model.UpdateUser(id, patchUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update data user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update user data",
		"user":    updateuser,
	})
}

// method delete user
// DeleteUser godoc
// @Summary      Delete user
// @Description  Menghapus user berdasarkan ID
// @Tags         Users
// @Produce      json
// @Param        id   path  int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Router       /api/v1/users/{id} [delete]
func (m *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalild id",
		})
		return
	}
	// take the user by id
	if err := m.Model.Deleteuser(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted info kopi dan gorengan bolo",
	})
}
