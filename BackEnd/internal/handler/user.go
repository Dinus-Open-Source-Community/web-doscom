package handler

import (
	"net/http"
	"strconv"
	"web_doscom/internal/auth"
	"web_doscom/internal/authorization"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/response"

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
		response.Error(
			c,
			http.StatusForbidden,
			"Role not valid, who are u??",
			err,
		)
		return
	}

	if validRole.Role == constants.RolePengurus {
		response.Error(
			c,
			http.StatusForbidden,
			"you are not allowed to access this!!",
			err,
		)
		return
	}

	var input dto.RegisterRequest
	// get the req body
	if c.ShouldBind(&input) != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Failed to read req body",
			err,
		)
		return
	}

	if input.Email == "" || input.Fullname == "" {
		response.Error(
			c,
			http.StatusBadRequest,
			"Missing fields, all fields are required",
			err,
		)
		return
	}

	if err := m.Service.InsertUserWithDefaultValue(&input, creatorRole); err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed while create user",
			err,
		)
		return
	}

	response.Success(
		c,
		"User created successfully",
		http.StatusCreated,
		nil,
	)

}

func (m *UserHandler) CreateSuperAdmin(c *gin.Context) {
	creatorRole := c.MustGet("role").(string)

	validRole, err := authorization.GetRoleInfo(creatorRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid",
			err,
		)
		return
	}

	if validRole.Role != constants.RoleAdmin {
		response.Error(
			c,
			http.StatusForbidden,
			"Not allowed to access this route, nakal yaa!!",
			err,
		)
		return
	}

	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Failed to read req body",
			err,
		)
		return
	}

	if input.Email == "" || input.Password == "" || input.Fullname == "" {
		response.Error(
			c,
			http.StatusBadRequest,
			"some field are missing",
			err,
		)
		return
	}

	// Check email uniqueness
	if _, err := m.Service.FindByEmail(input.Email); err == nil {
		response.Error(
			c,
			http.StatusConflict,
			"email already registered",
			err,
		)
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
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to create user",
			err,
		)
		return
	}

	response.Success(
		c,
		"Superadmin created successfully",
		http.StatusCreated,
		nil,
	)
}

func (m *UserHandler) GetUserByID(c *gin.Context) {
	CurrentUserRole := c.MustGet("role").(string)
	CurrentUserID := c.MustGet("user_id").(int)
	_, err := authorization.GetRoleInfo(CurrentUserRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid",
			err,
		)
		return
	}

	targetUserID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid user id",
			err,
		)
		return
	}

	userData, err := m.Service.GetUserById(CurrentUserRole, targetUserID, CurrentUserID)
	if err != nil {
		response.Error(
			c,
			http.StatusNotFound,
			"User Not Found",
			err,
		)
		return
	}

	TargetUserData := dto.UserResponse{
		Id:        userData.ID,
		Username:  userData.Username,
		Email:     userData.Email,
		Role:      userData.Role,
		Full_name: userData.Full_name,
	}

	response.Success(
		c,
		"Get user",
		http.StatusOK,
		TargetUserData,
	)
}

func (m *UserHandler) GetAllUserBasedOnRole(c *gin.Context) {
	userRole := c.MustGet("role").(string)

	validRole, error := authorization.GetRoleInfo(userRole)
	if error != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid",
			error,
		)
		return
	}

	if validRole.Role == constants.RolePengurus {
		response.Error(
			c,
			http.StatusForbidden,
			"unable to process the request",
			nil,
		)
		return
	}

	// get all user
	usersData, err := m.Service.GetAllUserBaseOnRole(userRole)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to fetch users data",
			err,
		)
		return
	}

	response.Success(
		c,
		"List of users data",
		http.StatusOK,
		usersData,
	)
}

func (m *UserHandler) UpdateUserByID(c *gin.Context) {
	userRole := c.MustGet("role").(string)
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid to proced this action",
			err,
		)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalid id",
			err,
		)
	}

	var userUpdateData dto.UserPatch
	if err := c.ShouldBindJSON(&userUpdateData); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Failed to read req body",
			err,
		)
		return
	}

	userDataToUpdate := userUpdateData.ToMap()
	updatedDataUser, err := m.Service.UpdateUser(id, userDataToUpdate)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to update data user",
			err,
		)
		return

	}

	response.Success(
		c,
		"Successfully update user data",
		http.StatusOK,
		updatedDataUser,
	)
}

func (m *UserHandler) DeleteUser(c *gin.Context) {
	userRole := c.MustGet("role").(string)
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"cannot proced",
			err,
		)
		return
	}

	if validRole.Role == constants.RolePengurus {
		response.Error(
			c,
			http.StatusForbidden,
			"you are not allowed",
			err,
		)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"invalild id",
			err,
		)
		return
	}

	if err := m.Service.DeleteUserBaseOnRole(id, userRole); err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"error while delete user",
			err,
		)
		return
	}

	response.Success(
		c,
		"user deleted info kopi dan gorengan bolo",
		http.StatusOK,
		nil,
	)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	currentUserID := c.MustGet("user_id").(int)
	userRole := c.MustGet("role").(string)
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid",
			err,
		)
		return
	}

	var changePasswordRequest dto.ChangePasswordRequest
	if err := c.ShouldBind(&changePasswordRequest); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Failed to read req body",
			err,
		)
		return
	}

	if err := h.Service.ChangePassword(
		currentUserID,
		changePasswordRequest,
		userRole,
	); err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to change password",
			err,
		)
		return
	}

	response.Success(
		c,
		"password changed successfully, anjayy aku tau pw mu :)",
		http.StatusCreated,
		nil,
	)
}

func (h *UserHandler) ChangePasswordAdmin(c *gin.Context) {
	currentUserRole := c.MustGet("role").(string)
	targetUserID, err := strconv.Atoi(c.Param("id"))
	if err != nil || targetUserID <= 0 {
		response.Error(
			c,
			http.StatusBadRequest,
			"Invalid user id",
			err,
		)
		return
	}

	validRole, err := authorization.GetRoleInfo(currentUserRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid",
			err,
		)
		return
	}

	if validRole.Role != constants.RoleAdmin {
		response.Error(
			c,
			http.StatusForbidden,
			"you are not allowed to access this route,, hayo ngapain masuk sini",
			err,
		)
		return
	}

	var changePasswordRequest dto.AdminChangePasswordRequest
	if err := c.ShouldBind(&changePasswordRequest); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Failed to read req body",
			err,
		)
		return
	}

	_, err = h.Service.ChangePasswordAdmin(targetUserID, currentUserRole, changePasswordRequest)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"failed to change password, nakal ganti pw wong liyo, mentang mentang superadmin",
			err,
		)
		return
	}

	response.Success(
		c,
		"password changed successfully",
		http.StatusOK,
		nil,
	)
}

func (m *UserHandler) GetSuperAdmin(c *gin.Context) {
	ctx := c.Request.Context()
	userRole := c.MustGet("role").(string)
	userID := c.MustGet("user_id").(int)
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid",
			err,
		)
		return
	}

	if validRole.Role != constants.RoleAdmin {
		response.Error(
			c,
			http.StatusForbidden,
			"you are not allowed to access this route, nakal yaa!!",
			err,
		)
		return
	}
	superAdminData, err := m.Service.GetSuperAdmin(ctx, userRole, userID)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"terjadi kesalahan ketika ambil data",
			err,
		)
		return
	}

	response.Success(
		c,
		"Get super admin data",
		http.StatusOK,
		superAdminData,
	)
}

func (m *UserHandler) GetCurrentUser(c *gin.Context) {
	currentUserId := c.MustGet("user_id").(int)
	currentUserRole := c.MustGet("role").(string)

	_, err := authorization.GetRoleInfo(currentUserRole)
	if err != nil {
		response.Error(
			c,
			http.StatusForbidden,
			"role not valid",
			err,
		)
		return
	}

	userData, err := m.Service.GetCurrentUser(currentUserId, currentUserRole)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"terjadi kesalahan ketika ambil data",
			err,
		)
		return
	}

	response.Success(
		c,
		"current user data",
		http.StatusOK,
		userData,
	)
}

func (m *UserHandler) UpdateProfileUser(c *gin.Context) {
	userID := c.MustGet("user_id").(int)
	userRole := c.MustGet("role").(string)
	ctx := c.Request.Context()

	var input dto.UserPatch
	if err := c.ShouldBind(&input); err != nil {
		response.Error(
			c,
			http.StatusBadRequest,
			"Failed to read req body",
			err,
		)
		return
	}

	userDataToUpdate := input.ToMap()
	updatedDataUser, err := m.Service.UpdateUserProfile(ctx, userID, userRole, userDataToUpdate)
	if err != nil {
		response.Error(
			c,
			http.StatusInternalServerError,
			"Failed to update data user",
			err,
		)
		return
	}

	response.Success(
		c,
		"Successfully update user data",
		http.StatusOK,
		updatedDataUser,
	)
}
