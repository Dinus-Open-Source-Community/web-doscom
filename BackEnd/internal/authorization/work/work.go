package workAuthorization

import (
	"fmt"
	"log"
	"strings"
	"web_doscom/internal/authorization"
	"web_doscom/internal/constants"
)

type workStatusGroup int

const (
	statusEditable workStatusGroup = iota
	statusModerated
	statusAll
)

var workRoleStatusGroup = map[string]workStatusGroup{
	constants.RoleKeyKoorData:     statusEditable,
	constants.RoleKeyKoorJaringan: statusEditable,
	constants.RoleKeyKoorPemro:    statusEditable,
	constants.RoleKeyKoorMedcrev:  statusEditable,
	constants.RoleKeyBPH:          statusModerated,
	constants.RoleKeySuperAdmin:   statusAll,
}

var (
	workStatusEditable = map[string]struct{}{
		constants.StatusDraft:   {},
		constants.StatusPending: {},
	}

	// untuk user validasi
	workStatusModerated = map[string]struct{}{
		constants.StatusPublished: {},
		constants.StatusScheduled: {},
		constants.StatusUnpublish: {},
		constants.StatusRejected:  {},
	}

	workStatusAll = map[string]struct{}{
		constants.StatusDraft:     {},
		constants.StatusPending:   {},
		constants.StatusPublished: {},
		constants.StatusScheduled: {},
		constants.StatusUnpublish: {},
		constants.StatusRejected:  {},
	}
)

var workStatusDefault = map[workStatusGroup]string{
	statusEditable:  constants.StatusDraft,
	statusModerated: constants.StatusPending,
	statusAll:       constants.StatusPending,
}

var workRoleKeyPermissionStatus = map[string]map[string]struct{}{
	constants.RoleKeyKoorMedcrev:  workStatusEditable,
	constants.RoleKeyKoorPemro:    workStatusEditable,
	constants.RoleKeyKoorJaringan: workStatusEditable,
	constants.RoleKeyKoorData:     workStatusEditable,
	constants.RoleKeyBPH:          workStatusModerated,
	constants.RoleKeySuperAdmin:   workStatusAll,
}

func setDefaultStatus(userRole string) string {
	group, ok := workRoleStatusGroup[userRole]
	if !ok {
		return constants.StatusDraft
	}

	return workStatusDefault[group]
}

func CanSetStatusWork(userRole, status string) (string, error) {
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return "", fmt.Errorf("invalid user role")
	}

	if _, ok := workStatusAll[status]; !ok {
		return "", fmt.Errorf("invalid status")
	}

	if strings.TrimSpace(status) == "" {
		defaultStatus := setDefaultStatus(userRole)
		log.Printf("status is empty, use default value")
		return defaultStatus, nil
	}

	if allowedStatus, ok := workRoleKeyPermissionStatus[userRole]; ok {
		if _, exist := allowedStatus[status]; exist {
			return status, nil
		}
		return "", fmt.Errorf("status not allowed")
	}

	return "", fmt.Errorf("permissions denied for this role %s", userRole)
}
