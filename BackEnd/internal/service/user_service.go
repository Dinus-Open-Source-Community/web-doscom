package service

import (
	"context"
	"fmt"
	"hash"
	"log"
	"os"
	"strings"
	"web_doscom/internal/auth"
	env "web_doscom/internal/config"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"
)

type UserService struct {
	UserModel *model.UserModel
}

func NewUserService(m *model.UserModel) *UserService {
	return &UserService{UserModel: m}
}

func (s *UserService) SetDefaultValue(email, fullname, creatorRole, reqRole string) (*model.DefaultValue, error) {
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

	return &model.DefaultValue{
		Username: username,
		Password: defaultPassword,
		Role:     assignedRole,
	}, nil
}

// wrapper function for user model
func (s *UserService) InsertUser(
	userData *model.RegisterRequest,
	userRole string,
) error {
	// validasi userRole
	creatorRole, ok := constants.RoleGroup[userRole]
	if !ok {
		return fmt.Errorf("role not valid %w")
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

	user := &model.User{
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

func (s *UserService) FindByEmail(email string) (*model.User, error) {
	return s.UserModel.FindByEmail(email)
}

func (s *UserService) GetUserById(id int) (*model.User, error) {
	return s.UserModel.GetUserById(id)
}

func (s *UserService) GetAllUser(role string, userRole string) ([]model.User, error) {
	return s.UserModel.GetAllUser(role, userRole)
}

func (s *UserService) UpdateUser(Id int, patch map[string]any) (*model.User, error) {
	return s.UserModel.UpdateUser(Id, patch)
}

func (s *UserService) DeleteUser(id int) error {
	return s.UserModel.DeleteUser(id)
}
