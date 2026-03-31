package utils

import (
	"fmt"
	"slices"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"
)

func SetDivitionAndPositionByRole(position, userRole string) (string, string, error) {
	var divisi, validPosition string

	// cek role user
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return "", "", fmt.Errorf("role not valid")
	}
	// auto assign division
	divisi = role.Divisi

	// check position valid for divisi or not
	positionRole, ok := constants.ValidPosition[position]
	if !ok {
		return "", "", fmt.Errorf("position not valid")
	}

	// take the position valid for divisi
	allowedPosition, ok := constants.PositionGroup[divisi]
	if !ok {
		return "", "", fmt.Errorf("divisi not valid")
	}

	// check position from param is can be assign to divisi or not
	if slices.Contains(allowedPosition, position) {
		return "", "", fmt.Errorf("position not allowed for this divisi")
	}

	if positionRole != role.Role {
		return "", "", fmt.Errorf("role not allowed ")
	}

	return divisi, validPosition, nil
}

func FilterRoleFieldPermission(userRole string, data *model.PengurusPatch) (map[string]any, error) {
	role, ok := constants.RoleGroup[userRole]
	if !ok {
		return nil, fmt.Errorf("role not valid")
	}

	actorRole := role.Role
	allowedFields, exists := constants.RoleFieldPermission[actorRole]
	if !exists {
		return nil, fmt.Errorf("this role has no defined permissions")
	}

	allField := StructToMap(data)

	editableFieds := make(map[string]any)
	for _, field := range allowedFields {
		if val, ok := allField[field]; ok {
			editableFieds[field] = val
		}
	}

	return editableFieds, nil
}
