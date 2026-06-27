package pengurus

import (
	"context"
	"fmt"
	"slices"

	"web_doscom/internal/authorization"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/entity"
)

func RolePositionAuthorization(
	ctx context.Context,
	idParams, currentUserID int,
	userRole string,
	pengurusData entity.Pengurus,
) (string, error) {
	// white list role
	actor, ok := constants.RoleGroup[userRole]
	if !ok {
		return "", fmt.Errorf("role not valid")
	}

	actorRole := actor.Role
	// SuperAdmin bypass
	if actorRole == constants.RoleAdmin {
		return actorRole, nil
	}

	// cek role pengurus
	if actorRole == constants.RolePengurus && currentUserID != idParams {
		return "", fmt.Errorf("You are not allowed to update this data")
	}

	// whitelist position
	targetPositionRole := constants.ValidPosition[pengurusData.Position]

	// ckeck user level role
	actorLevel := constants.RoleLevel[actorRole]
	targetLevel := constants.RoleLevel[targetPositionRole]

	if actorLevel >= targetLevel && currentUserID != idParams {
		return "", fmt.Errorf("You are not allowed to update this data")
	}

	return actorRole, nil
}

func CanEditPengurusField(userRole string, field string) (bool, error) {
	validRole, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return false, fmt.Errorf("invalid user role")
	}

	allowedFields, exists := constants.RoleFieldPermission[validRole.Role]
	if !exists {
		return false, fmt.Errorf("no permissions defined for this role")
	}

	// for _, allowfield := range allowedFields {
	// 	if field == allowfield {
	// 		return true, nil
	// 	}
	// }
	if slices.Contains(allowedFields, field) {
		return true, nil
	}

	return false, fmt.Errorf("field %s not allowed to edit or insert by u", field)
}
