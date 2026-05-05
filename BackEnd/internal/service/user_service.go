package service

import (
	"fmt"
	"os"
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
	switch {
	case validRoleCreator.Role == constants.RoleKoordinator:
		assignedRole = constants.AutoAsignRole[creatorRole]
	case validRoleCreator.Role == constants.RoleAdmin:
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
		creatorRole.Role == constants.RoleKoordinator {
		return fmt.Errorf("invalid role")
	}

	defaultValue, err := s.SetDefaultValue(
		userData.Email,
		userData.Username,
		userRole,
		userData.Role,
	)
	if err != nil {
		return fmt.Errorf("failed to set defaultValue to new user")
	}

	// hash password
	passowordHash := auth.HashPassword(defaultValue.Password)

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

func (s *UserService) GetUserById(id int) (*entity.User, error) {
	return s.UserModel.GetUserById(id)
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
		return fmt.Errorf("user tidak di temukan %w", err)
	}
	userToDeleteRoleGroup := constants.RoleGroup[userWantToDeleteData.Role]

	userRoleLevel := constants.RoleLevel[userValidRoleGroup.Role]
	userToDeleteRoleLevel := constants.RoleLevel[userToDeleteRoleGroup.Role]

	if userRoleLevel <= userToDeleteRoleLevel && id != userWantToDeleteData.ID {
		return fmt.Errorf("you are not allowed to delete this data")
	}

	if userValidRoleGroup.Divisi != userToDeleteRoleGroup.Divisi {
		return fmt.Errorf("you cannot delete data from another divison")
	}

	if err := s.UserModel.DeleteUser(id); err != nil {
		return fmt.Errorf("terjadi kesalahan %w", err)
	}

	return nil
}
