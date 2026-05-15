package pengurus

import (
	"context"
	"fmt"
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
