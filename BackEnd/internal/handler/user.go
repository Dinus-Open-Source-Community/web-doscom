package handler

import (
	"net/http"
	"strconv"
	"web_doscom/internal/auth"
	"web_doscom/internal/authorization"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"

	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (m *UserHandler) CreateUser(c *gin.Context) {
	// get the user role
	creatorRole := c.MustGet("role").(string)

	validRole, err := authorization.GetRoleInfo(creatorRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Role not valid, who are u??",
		})
		return
	}

	if validRole.Role == constants.RolePengurus {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are not allowed to access this!!",
		})
		return
	}

	var input dto.RegisterRequest
	// get the req body
	if c.ShouldBind(&input) != nil {
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

	if err := m.Service.InsertUserWithDefaultValue(&input, creatorRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed while create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})

}

func (m *UserHandler) CreateSuperAdmin(c *gin.Context) {
	creatorRole := c.MustGet("role").(string)

	validRole, err := authorization.GetRoleInfo(creatorRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "terjadi kesalahan",
		})
		return
	}

	if validRole.Role != constants.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Not allowed to access this route, nakal yaa!!",
		})
		return
	}

	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	if input.Email == "" || input.Password == "" || input.Fullname == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "some field are missing",
		})
		return
	}

	// Check email uniqueness
	if _, err := m.Service.FindByEmail(input.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":   "email already registered",
			"message": "use another email, i know you have a lot of email",
		})
		return
	}

	passwordHash := auth.HashPassword(input.Password)

	user := dto.RegisterRequest{
		Username: input.Fullname,
		Email:    input.Email,
		Password: passwordHash,
		Role:     constants.RoleKeySuperAdmin,
		Fullname: input.Fullname,
	}

	if err := m.Service.InsertUserWithDefaultValue(&user, creatorRole); err != nil {
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
	CurrentUserRole := c.MustGet("role").(string)
	CurrentUserID := c.MustGet("user_id").(int)
	validRole, err := authorization.GetRoleInfo(CurrentUserRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "role not valid",
		})
		return
	}
	if validRole.Role == constants.RolePengurus {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "your role does not have access to this resource",
		})
		return
	}

	targetUserID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	userData, err := m.Service.GetUserById(CurrentUserRole, targetUserID, CurrentUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "User Not Found, err: " + err.Error(),
			"message": "ihh cari siapa sii, kok id ngga ada di cari cari uhmmmm ;)",
		})
		return
	}

	TargetUserData := dto.UserResponse{
		Id:        userData.ID,
		Username:  userData.Username,
		Email:     userData.Email,
		Role:      userData.Role,
		Full_name: userData.Full_name,
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Get user",
		"userData": TargetUserData,
	})
}

func (m *UserHandler) GetAllUserBasedOnRole(c *gin.Context) {
	userRole := c.MustGet("role").(string)

	validRole, error := authorization.GetRoleInfo(userRole)
	if error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": error.Error(),
		})
		return
	}

	if validRole.Role == constants.RolePengurus {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "unable to process the request",
			"message": "your role does not have access to this resource",
		})
		return
	}

	// get all user
	usersData, err := m.Service.GetAllUserBaseOnRole(userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to fetch users data",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "List of users data",
		"usersData": usersData,
	})
}

func (m *UserHandler) UpdateUser(c *gin.Context) {
	userRole := c.MustGet("role").(string)
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "role not valid to proced this action",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
	}

	var userUpdateData dto.UserPatch
	if err := c.ShouldBindJSON(&userUpdateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	userDataToUpdate := userUpdateData.ToMap()
	updatedDataUser, err := m.Service.UpdateUser(id, userDataToUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to update data user",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Successfully update user data",
		"userUpdatedData": updatedDataUser,
	})
}

func (m *UserHandler) DeleteUser(c *gin.Context) {
	userRole := c.MustGet("role").(string)
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "cannot proced",
		})
		return
	}

	if validRole.Role == constants.RolePengurus {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "you are not allowed",
			"message": "cannot proced this action",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalild id",
		})
		return
	}

	if err := m.Service.DeleteUserBaseOnRole(id, userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "error while delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted info kopi dan gorengan bolo",
	})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	currentUserID := c.MustGet("user_id").(int)
	userRole := c.MustGet("role").(string)
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	var changePasswordRequest dto.ChangePasswordRequest
	if err := c.ShouldBind(&changePasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	if err := h.Service.ChangePassword(
		currentUserID,
		changePasswordRequest,
		userRole,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to change password",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "password changed successfully, anjayy aku tau pw mu :)",
	})
}

func (h *UserHandler) ChangePasswordAdmin(c *gin.Context) {
	currentUserRole := c.MustGet("role").(string)
	targetUserID, err := strconv.Atoi(c.Param("id"))
	if err != nil || targetUserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id",
		})
	}

	validRole, err := authorization.GetRoleInfo(currentUserRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "role not valid",
		})
		return
	}

	if validRole.Role != constants.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are not allowed to access this route,, hayo ngapain masuk sini",
		})
		return
	}

	var changePasswordRequest dto.AdminChangePasswordRequest
	if err := c.ShouldBind(&changePasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read req body",
		})
		return
	}

	passwordUpdated, err := h.Service.ChangePasswordAdmin(targetUserID, currentUserRole, changePasswordRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "failed to change password, nakal ganti pw wong liyo, mentang mentang superadmin",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "password changed successfully",
		"passwordUpdated": passwordUpdated,
	})
}

func (m *UserHandler) GetSuperAdmin(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   err.Error(),
			"message": "role not valid",
		})
		return
	}

	if validRole.Role != constants.RoleAdmin {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are not allowed to access this route, nakal yaa!!",
		})
		return
	}
	superAdminData, err := m.Service.GetSuperAdmin(ctx, userRole, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "terjadi kesalahan ketika ambil data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "Get super admin data",
		"superAdmin Data": superAdminData,
	})
}
