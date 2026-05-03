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
	statusEditable  workStatusGroup = iota // 0
	statusModerated                        // 1
	statusAll                              // 2
)

var workRoleStatusGroup = map[string]workStatusGroup{
	constants.RoleKeyKoorData:     statusEditable,
	constants.RoleKeyKoorJaringan: statusEditable,
	constants.RoleKeyKoorPemro:    statusEditable,
	constants.RoleKeyKoorMedcrev:  statusEditable,
	constants.RoleKeyBPH:          statusModerated,
	constants.RoleKeySuperAdmin:   statusAll,
}

var workGroupAllowedStatus = map[workStatusGroup]map[string]struct{}{
	statusEditable: {
		constants.StatusDraft:   {},
		constants.StatusPending: {},
	},

	// untuk user validasi
	statusModerated: {
		constants.StatusPublished: {},
		constants.StatusScheduled: {},
		constants.StatusUnpublish: {},
		constants.StatusRejected:  {},
	},

	statusAll: {
		constants.StatusDraft:     {},
		constants.StatusPending:   {},
		constants.StatusPublished: {},
		constants.StatusScheduled: {},
		constants.StatusUnpublish: {},
		constants.StatusRejected:  {},
	},
}

var workStatusDefault = map[workStatusGroup]string{
	statusEditable:  constants.StatusDraft,
	statusModerated: constants.StatusPending,
	statusAll:       constants.StatusPending,
}

var workStatusToGroup = map[string]workStatusGroup{
	constants.StatusDraft:     statusEditable,
	constants.StatusPending:   statusEditable,
	constants.StatusPublished: statusModerated,
	constants.StatusScheduled: statusModerated,
	constants.StatusUnpublish: statusModerated,
	constants.StatusRejected:  statusModerated,
}

func getRoleGroup(userRole string) (workStatusGroup, bool) {
	group, ok := workRoleStatusGroup[userRole]

	return group, ok
}

func setDefaultStatus(userRole string) string {
	group, ok := getRoleGroup(userRole)
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

	group, ok := getRoleGroup(userRole)
	if !ok {
		return "", fmt.Errorf("permissions denied for this role %s", userRole)
	}

	allowedStatus := workGroupAllowedStatus[group]

	if strings.TrimSpace(status) == "" {
		log.Printf("status is empty, use default value")
		return workStatusDefault[group], nil
	}

	// cek apakah status yang diberikan itu ada di daftar atau tidak
	if _, validGlobal := workGroupAllowedStatus[statusAll][status]; !validGlobal {
		return "", fmt.Errorf("status not valid")
	}

	if _, allowed := allowedStatus[status]; !allowed {
		return "", fmt.Errorf("status not allowed for this role %s", userRole)
	}

	return status, nil
}

func CanDeleteWork(userRole, status string) (bool, error) {
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return false, fmt.Errorf("invalid userRole")
	}

	roleStatusGroup, ok := workRoleStatusGroup[userRole]
	if !ok {
		return false, fmt.Errorf("permissions denied for this role %s", userRole)
	}

	statusGroup, ok := workStatusToGroup[status]
	if !ok {
		return false, fmt.Errorf("status not valid")
	}

	// aturan delete work
	// bisa hapus work jika role group nya itu >= status yang mau di hapus
	// contoh
	// 	koordinator (0) -> hapus work status draft (0) -> (0) >= (0) -> true
	// 	koordinator (0) -> hapus work status published (1) -> (0) >= (1) -> false
	// 	bph (1) -> hapus work status publishe (1)	 -> (1) >= (1) -> true
	// 	super admin (2) -> hapus work status apapun -> (2) >= (*) -> true
	if roleStatusGroup >= statusGroup {
		return true, nil
	}
	return false, fmt.Errorf("denied: role %s cannot delete work with status %s", userRole, status)
}
