package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"web_doscom/internal/auth"
	"web_doscom/internal/authorization"
	env "web_doscom/internal/config"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
)

type UserService struct {
	UserModel *entity.UserModel
}

func NewUserService(m *entity.UserModel) *UserService {
	return &UserService{UserModel: m}
}

func (s *UserService) SetDefaultValue(email, fullname, creatorRole, reqRole string) (*dto.DefaultValue, error) {
	env.LoadEnv()
	secret := os.Getenv("PASSWORD_SECRET")

	// validasi creatorRole
	validRoleCreator, ok := constants.RoleGroup[creatorRole]
	if !ok {
		return nil, fmt.Errorf("role not valid")
	}

	var assignedRole string
	switch validRoleCreator.Role {
	case constants.RoleKoordinator:
		assignedRole = constants.AutoAsignRole[creatorRole]
	case constants.RoleAdmin:
		if reqRole == "" {
			return nil, fmt.Errorf("role is required for admin")
		}

		if _, ok := constants.RoleGroup[reqRole]; !ok {
			return nil, fmt.Errorf("role not valid")
		}
		assignedRole = reqRole
	default:
		return nil, fmt.Errorf("this role (%s) cannot create user", validRoleCreator.Role)
	}

	if assignedRole == "" {
		return nil, fmt.Errorf("failed to assign role")
	}

	// atIndex := strings.Index(email, "@")
	before, _, found := strings.Cut(email, "@")
	var partEmail string
	if found {
		partEmail = before
	} else {
		partEmail = "user_" + assignedRole
	}

	fullnamePart := strings.Split(fullname, " ")[0]

	username := fullnamePart + "_" + partEmail
	defaultPassword := partEmail + secret + fullnamePart

	return &dto.DefaultValue{
		Username: username,
		Password: defaultPassword,
		Role:     assignedRole,
	}, nil
}

// wrapper function for userdto
func (s *UserService) InsertUserWithDefaultValue(
	userData *dto.RegisterRequest,
	userRole string,
) error {
	// validasi userRole
	creatorRole, ok := constants.RoleGroup[userRole]
	if !ok {
		return fmt.Errorf("role not valid")
	}

	if creatorRole.Role != constants.RoleAdmin &&
		creatorRole.Role != constants.RoleKoordinator {
		return fmt.Errorf("invalid role")
	}

	defaultValue, err := s.SetDefaultValue(
		userData.Email,
		userData.Fullname,
		userRole,
		userData.Role,
	)
	if err != nil {
		return fmt.Errorf("failed to set defaultValue to new user: %w", err)
	}

	// check email

	// hash password
	passowordHash := auth.HashPassword(defaultValue.Password)

	log.Println("password", defaultValue.Password)
	user := &entity.User{
		Username:  defaultValue.Username,
		Email:     userData.Email,
		Role:      defaultValue.Role,
		Password:  passowordHash,
		Full_name: userData.Fullname,
	}

	// create user
	if err := s.UserModel.InsertUser(user); err != nil {
		return fmt.Errorf("failed to create user %w", err)
	}

	return nil
}

func (s *UserService) FindByEmail(email string) (*entity.User, error) {
	return s.UserModel.FindByEmail(email)
}

func (s *UserService) GetUserById(userRole string, targetUserID, currentUserID int) (*entity.User, error) {
	validCurrentUser, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return nil, err
	}

	dataUserTarget, err := s.UserModel.GetUserById(targetUserID)
	if err != nil {
		return nil, fmt.Errorf("terjadi kesalahan ketika ambil data %w", err)
	}
	validTargetUser, err := authorization.GetRoleInfo(dataUserTarget.Role)
	if err != nil {
		return nil, fmt.Errorf("something wongg wakk")
	}

	switch validCurrentUser.Role {
	case constants.RoleAdmin:
	case constants.RoleKoordinator:
		if validTargetUser.Divisi != validCurrentUser.Divisi {
			return nil, fmt.Errorf("you can not see other division data bro")
		}
	case constants.RolePengurus:
		if currentUserID != dataUserTarget.ID {
			return nil, fmt.Errorf("you can't see other user data bro, i wait you...")
		}
	default:
		return nil, fmt.Errorf("role not valid")
	}

	return dataUserTarget, nil
}

func (s *UserService) GetAllUserBaseOnRole(
	creatorRole string,
) ([]dto.UserResponse, error) {
	validRole, err := authorization.GetRoleInfo(creatorRole)
	if err != nil {
		return nil, err
	}

	var assignRole string
	switch validRole.Role {
	case constants.RoleAdmin:
		assignRole = constants.RoleKeySuperAdmin
	case constants.RoleKoordinator:
		assignRole = constants.AutoAsignRole[creatorRole]
	case constants.RolePengurus:
		return nil, fmt.Errorf("this role cannot access this data %s", validRole.Role)
	default:
		return nil, fmt.Errorf("need creator role to get data")
	}

	// userData, err := s.UserModel.GetAllUserBaseOnRole(assignRole)
	userData, err := s.UserModel.GetAllUserBaseOnRole(assignRole)
	if err != nil {
		return nil, fmt.Errorf("terjadi kesalahan di database: %w", err)
	}

	return userData, nil
}

func (s *UserService) UpdateUser(Id int, userDataToUpdate map[string]any) (*dto.UserResponse, error) {

	allowedFieldToUpdate := map[string]bool{
		"username":  true,
		"email":     true,
		"full_name": true,
	}

	filteredUpdates := make(map[string]any)
	for field, value := range userDataToUpdate {
		if allowedFieldToUpdate[field] && value != nil {
			if str, ok := value.(string); !ok || str != "" {
				filteredUpdates[field] = value
			}
		}
	}
	fmt.Println("Filtered updates:", filteredUpdates)

	updatedUserData, err := s.UserModel.UpdateUser(Id, userDataToUpdate)
	if err != nil {
		return nil, fmt.Errorf("terjadi kesalahan ketika update %w", err)
	}

	return updatedUserData, nil
}

func (s *UserService) DeleteUserBaseOnRole(id int, userRole string) error {
	userValidRoleGroup, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("role not valid, cannot proceed this action")
	}

	userWantToDeleteData, err := s.UserModel.GetUserById(id)
	if err != nil {
		log.Printf("[service] ini di execute wak")
		return fmt.Errorf("user nya ngga ada wlee :|>%w", err)
	}
	userToDeleteRoleGroup := constants.RoleGroup[userWantToDeleteData.Role]

	userRoleLevel := constants.RoleLevel[userValidRoleGroup.Role]
	userToDeleteRoleLevel := constants.RoleLevel[userToDeleteRoleGroup.Role]

	if userRoleLevel >= userToDeleteRoleLevel {
		return fmt.Errorf("you are not allowed to delete this data")
	}

	if userValidRoleGroup.Role == constants.RoleAdmin {
		log.Printf("[service] user id: %d", id)
		if err := s.UserModel.DeleteUser(id); err != nil {
			log.Printf("[service] ini di execute")
			return fmt.Errorf("terjadi kesalahan %w", err)
		}
	} else {
		if userValidRoleGroup.Divisi != userToDeleteRoleGroup.Divisi {
			return fmt.Errorf("you cannot delete data from another divison")
		}
		if err := s.UserModel.DeleteUser(id); err != nil {
			log.Printf("[service] ini ter execute")
			return fmt.Errorf("terjadi kesalahan %w", err)
		}
	}

	return nil
}

func checkPassword(password string) (bool, string) {
	score := 0

	check := []*regexp.Regexp{
		constants.AtLeastOneLowercase,
		constants.AtLeastOneUppercase,
		constants.AtLeastOneNumeric,
		constants.AtLeastOneSpecialChar,
		constants.EightCharsOrMore,
	}

	for _, regex := range check {
		if regex.MatchString(password) {
			score++
		}
	}

	switch {
	case score <= 2:
		return false, fmt.Sprintf("password too weak, like you weak: %d", score)
	case score <= 4:
		return true, fmt.Sprintf("password now little strong, make it stronger: %d", score)
	default:
		return true, "damnn your passord is strong like develeoper"
	}
}

func (s *UserService) ChangePassword(currentUserID int, passwordRequest dto.ChangePasswordRequest, userRole string) error {
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("role not valid")
	}

	currentUserData, err := s.UserModel.GetUserById(currentUserID)
	if err != nil {
		return fmt.Errorf("terjadi kesalahan ketika ambil data %w", err)
	}

	// check password old
	if !auth.VerifyPassword(passwordRequest.OldPassword, currentUserData.Password) {
		return fmt.Errorf("old password is not correct, hayoo kok lupa pw nya")
	}

	// check new password strength
	valid, message := checkPassword(passwordRequest.NewPassword)
	if !valid {
		return fmt.Errorf("%s", message)
	}

	// update password
	passwordUpdated := make(map[string]any)
	passwordUpdated["password"] = auth.HashPassword(passwordRequest.NewPassword)

	_, err = s.UserModel.UpdateUser(currentUserID, passwordUpdated)
	if err != nil {
		return fmt.Errorf("terjadi kesalahan ketika update password %w", err)
	}
	return nil
}

func (s *UserService) ChangePasswordAdmin(targetUserID int, userRole string, passwordRequest dto.AdminChangePasswordRequest) (string, error) {
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return "", fmt.Errorf("role not valid")
	}

	if validRole.Role != constants.RoleAdmin {
		return "", fmt.Errorf("you are not allowed to access this route, sopo kwe ganti pw wong liyo")
	}

	userData, err := s.UserModel.GetUserById(targetUserID)
	if err != nil {
		return "", fmt.Errorf("terjadi kesalahan ketika ambil data %w", err)
	}
	validTargetUser, err := authorization.GetRoleInfo(userData.Role)
	if err != nil {
		return "", fmt.Errorf("something wongg wakk role not valid")
	}

	roleLevelCurrentUser := constants.RoleLevel[validRole.Role]
	roleLevelTargetUser := constants.RoleLevel[validTargetUser.Role]

	if roleLevelCurrentUser >= roleLevelTargetUser {
		return "", fmt.Errorf("You cannot change the password for a role that is the same as your own")
	}

	passwordUpdated := make(map[string]any)
	passwordUpdated["password"] = auth.HashPassword(passwordRequest.NewPassword)
	_, err = s.UserModel.UpdateUser(targetUserID, passwordUpdated)
	if err != nil {
		return "", fmt.Errorf("terjadi kesalahan ketika update password %w", err)
	}

	return passwordRequest.NewPassword, nil
}

func (s *UserService) GetSuperAdmin(ctx context.Context, userRole string, userID int) ([]dto.UserResponse, error) {
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return nil, err
	}
	if validRole.Role != constants.RoleAdmin {
		return nil, fmt.Errorf("you are not allowed to access this route, nakal yaa!!")
	}

	superAdminData, err := s.UserModel.GetSuperAdmin(ctx, userRole, userID)
	if err != nil {
		return nil, fmt.Errorf("terjadi kesalahan ketika ambil data %w", err)
	}

	return superAdminData, nil
}
