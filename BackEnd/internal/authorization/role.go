package authorization

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/utils"
)

func GetRoleInfo(role string) (constants.Divisioninfo, error) {
	validRole, exist := constants.RoleGroup[role]
	if !exist {
		return constants.Divisioninfo{}, fmt.Errorf("role tidak valid")
	}
	return validRole, nil
}

func SetDivitionAndPositionByRole(position, requestedDivisi, userRole string) (string, string, error) {
	// Clean inputs
	position = strings.TrimSpace(position)
	requestedDivisi = strings.ToLower(strings.TrimSpace(requestedDivisi))

	// Get actor role information
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return "", "", fmt.Errorf("invalid user role")
	}

	// SUPERADMIN PRIVILEGE: If the actor has Admin role, allow all divisions and positions
	if role.Role == constants.RoleAdmin {
		return requestedDivisi, position, nil
	}

	// FOR NON-ADMIN (e.g., Coordinators)
	// Auto-assign division from actor's profile
	divisi := strings.ToLower(strings.TrimSpace(role.Divisi))

	// Check if the requested position exists in system
	_, ok = constants.ValidPosition[position]
	if !ok {
		return "", "", fmt.Errorf("position not recognized by system")
	}

	// Get allowed positions for the specific division
	allowedPositions, ok := constants.PositionGroup[divisi]
	if !ok {
		return "", "", fmt.Errorf("division not valid or not found")
	}

	// Verify if the position belongs to the division
	if !slices.Contains(allowedPositions, position) {
		return "", "", fmt.Errorf("position '%s' is not allowed for division '%s'", position, divisi)
	}

	return divisi, position, nil
}

func FilterRoleFieldPermission(userRole string, data *dto.PengurusPayload) (map[string]any, error) {
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return nil, fmt.Errorf("invalid user role")
	}

	allFields := utils.StructToMap(data)

	// SUPERADMIN PRIVILEGE: Admin can edit all fields
	if role.Role == constants.RoleAdmin {
		return allFields, nil
	}

	// Role-based filtering for others
	allowedFields, exists := constants.RoleFieldPermission[role.Role]
	if !exists {
		return nil, fmt.Errorf("no permissions defined for this role")
	}
	log.Printf("allowed fields: %#v", allowedFields)

	editableFields := make(map[string]any)
	for _, field := range allowedFields {
		if val, ok := allFields[field]; ok {
			editableFields[field] = val
		}
	}

	return editableFields, nil
}

func CheckRolePermission(userRole string, grantedRole ...string) error {

	if len(grantedRole) == 0 {
		return fmt.Errorf("role is required not empty")
	}

	validRole, err := GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("invalid user role")
	}

	for _, role := range grantedRole {
		// check by role level
		if validRole.Role == role {
			return nil
		}

		// check by role key
		if userRole == role {
			return nil
		}
	}

	return fmt.Errorf("permissions denied for this role %s", userRole)
}
