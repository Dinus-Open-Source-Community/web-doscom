package handler

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"web_doscom/internal/auth"
	env "web_doscom/internal/config"
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

func (m *UserHandler) SetDefaultValue(email, fullname, role string) (*model.DefaultValue, error) {
	env.LoadEnv()
	secret := os.Getenv("PASSWORD_SECRET")

	var assignedRole, roles string
	if strings.HasPrefix(role, "Kor_") {
		roles = strings.TrimPrefix(role, "Kor_")
		assignedRole = strings.ToLower(roles) + "_ang"
	} else if role == "BPH" {
		assignedRole = "BPH_ang"
	} else {
		log.Println(">>> CALLED handler.SetDefaultValue <<<", role)
		assignedRole = ""
	}

	atIndex := strings.Index(email, "@")
	partEmail := email
	if atIndex != -1 {
		partEmail = email[:atIndex]
	} else {
		partEmail = "user_" + roles
	}

	fullnamePart := strings.Split(fullname, " ")[0]

	username := fullnamePart + "_" + partEmail
	defaultPassword := partEmail + secret + fullnamePart

	return &model.DefaultValue{
		Username: username,
		Password: defaultPassword,
		Role:     assignedRole,
	}, nil
}

func (m *UserHandler) CreateUser(c *gin.Context) {
	// get the user role
	creatorRole := c.MustGet("role").(string)

	var input model.RegisterRequest
	// get the req body
	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	if input.Email == "" || input.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing fields, all fields are required",
		})

		return
	}

	// set default value for user
	var (
		defaultValue *model.DefaultValue
		Err          error
	)
	switch {
	case strings.HasPrefix(creatorRole, "Kor_"):
		defaultValue, Err = m.SetDefaultValue(input.Email, input.Fullname, creatorRole)
		if Err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": Err.Error(),
			})
		}
	case creatorRole == "BPH":
		defaultValue, Err = m.SetDefaultValue(input.Email, input.Fullname, creatorRole)
		if Err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": Err.Error(),
			})
		}
	case creatorRole == "Super_Admin":
		if input.Role == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "role must be given",
			})
			return
		}
		defaultValue, Err = m.SetDefaultValue(input.Email, input.Fullname, creatorRole)
		defaultValue.Role = input.Role
		if Err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": Err.Error(),
			})
		}
	default:
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "Invalid role",
			"message": "heyy jangan kau pakai route ini, bapak kau 3",
		})
	}

	// hash the password
	passwordHash := auth.HashPassword(defaultValue.Password)
	log.Printf("Role value: '%s'", defaultValue.Role)

	// mapping the user
	user := &model.User{
		Username:  defaultValue.Username,
		Email:     input.Email,
		Role:      defaultValue.Role,
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

func (m *UserHandler) CreateSuperAdmin(c *gin.Context) {
	role := c.MustGet("role").(string)
	if role != "Super_Admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Heyy tinggi kali kau bah pake route ini",
		})
		return
	}
	// get the req body
	var input model.RegisterRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read req body"})
		return
	}

	if input.Email == "" || input.Password == "" || input.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, email, and password are required"})
		return
	}

	// Check email uniqueness
	if _, err := m.Model.FindByEmail(input.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
		return
	}

	// Hash password
	passwordHash := auth.HashPassword(input.Password)

	user := model.User{
		Username:  input.Fullname,
		Email:     input.Email,
		Password:  passwordHash,
		Role:      "Super_Admin",
		Full_name: input.Fullname,
	}

	if err := m.Model.InsertUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Superadmin created successfully",
	})
}

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

func (m *UserHandler) GetAllUser(c *gin.Context) {
	role := c.GetString("role")

	// condition to get all user by role
	var userRole string
	if strings.HasPrefix(role, "Super_Admin") {
		userRole = "Super_Admin"
	} else if strings.HasPrefix(role, "Kor_") {
		divisi := strings.TrimPrefix(role, "Kor_")
		userRole = strings.ToLower(divisi) + "_ang"
	} else {
		userRole = "BPH_ang"
	}

	// get all user
	users, err := m.Model.GetAllUser(role, userRole)
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
	userUpdate := patchUser.ToMap()
	updateuser, err := m.Model.UpdateUser(id, userUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update data user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully update user data",
		"user":    updateuser,
	})
}

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

// wrapper function for documentation and testing

// wrapper SuperAdmin

// CreateUserHandler untuk POST /api/user
// CreateUser godoc
// @Summary      Create new user
// @Description  Membuat user baru
// @Accept       json
// @Produce      json
// @Param        user  body  model.RegisterRequest  true  "Email, FullName, Role"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	       Super_Admin
// @Router       /api/v1/admin/ [post]
func (m *UserHandler) SuperAdminCreateUser(c *gin.Context) {
	m.CreateUser(c)
}

// CreateSuperAdmin godoc
// @Summary      Create superadmin user
// @Description  Membuat user dengan role superadmin (hanya untuk admin/superadmin)
// @Accept       json
// @Produce      json
// @Param        user  body  model.RegisterRequest  true  "User info"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Failure      409  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	       Super_Admin
// @Router       /api/v1/admin/superadmin [post]
func (m *UserHandler) SuperAdminCreateSuperAdmin(c *gin.Context) {
	m.CreateSuperAdmin(c)
}

// GetUserHandler untuk GET /api/user/:id
// GetUser godoc
// @Summary      Get user by ID
// @Description  Mendapatkan user berdasarkan ID
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.UserResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	       Super_Admin
// @Router       /api/v1/admin/{id} [get]
func (m *UserHandler) SuperAdminGetUser(c *gin.Context) {
	m.GetUser(c)
}

// get all user
// GetAllUser godoc
// @Summary      Get all users
// @Description  Mendapatkan daftar semua user
// @Produce      json
// @Success      200  {array}   model.UserResponse
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	       Super_Admin
// @Router       /api/v1/admin/ [get]
func (m *UserHandler) SuperAdminGetAllUser(c *gin.Context) {
	m.GetAllUser(c)
}

// update admin by id
// UpdateUser godoc
// @Summary      Update user
// @Description  Mengupdate data user berdasarkan ID
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "User ID"
// @Param        user  body      model.UserPatch  true  "User info"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	       Super_Admin
// @Router       /api/v1/admin/{id} [put]
func (m *UserHandler) SuperAdminUpdateUser(c *gin.Context) {
	m.UpdateUser(c)
}

// update admin by id
// UpdateUser godoc
// @Summary      Update user
// @Description  Mengupdate data user berdasarkan ID
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "User ID"
// @Param        user  body      model.UserPatch  true  "User info"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	       Super_Admin
// @Router       /api/v1/admin/{id} [put]
func (m *UserHandler) SuperAdminDeleteUser(c *gin.Context) {
	m.DeleteUser(c)
}

// wrapper para koor dan BPH

// CreateUserHandler untuk POST /api/user
// CreateUser godoc
// @Summary      Create new user
// @Description  Membuat user baru
// @Accept       json
// @Produce      json
// @Param        user  body  model.RegisterRequest  true  "Email, FullName"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	      Koor
// @Router       /api/v1/koor/ [post]
func (m *UserHandler) KoorCreateUser(c *gin.Context) {
	m.CreateUser(c)
}

// GetUserHandler untuk GET /api/user/:id
// GetUser godoc
// @Summary      Get user by ID
// @Description  Mendapatkan user berdasarkan ID
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.UserResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	      Koor
// @Router       /api/v1/koor/{id} [get]
func (m *UserHandler) KoorGetUser(c *gin.Context) {
	m.GetUser(c)
}

// get all user
// GetAllUser godoc
// @Summary      Get all users
// @Description  Mendapatkan daftar semua user
// @Produce      json
// @Success      200  {array}   model.UserResponse
// @Failure      500  {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	      Koor
// @Router       /api/v1/koor/ [get]
func (m *UserHandler) KoorGetAllUser(c *gin.Context) {
	m.GetAllUser(c)
}

// update admin by id
// UpdateUser godoc
// @Summary      Update user
// @Description  Mengupdate data user berdasarkan ID
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "User ID"
// @Param        user  body      model.UserPatch  true  "User info"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	      Koor
// @Router       /api/v1/koor/{id} [put]
func (m *UserHandler) KoorUpdateUser(c *gin.Context) {
	m.UpdateUser(c)
}

// update admin by id
// UpdateUser godoc
// @Summary      Update user
// @Description  Mengupdate data user berdasarkan ID
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "User ID"
// @Param        user  body      model.UserPatch  true  "User info"
// @Success      200   {object}  model.UserResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Security ApiKeyAuth
// @Tags 	      Koor
// @Router       /api/v1/koor/{id} [put]
func (m *UserHandler) KoorDeleteUser(c *gin.Context) {
	m.DeleteUser(c)
}
