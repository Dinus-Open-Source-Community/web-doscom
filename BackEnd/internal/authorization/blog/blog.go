package blogAuthorization

import (
	"fmt"
	"log"
	"strings"
	"web_doscom/internal/authorization"
	"web_doscom/internal/constants"
)

type blogStatusGroup int

const (
	statusEditable blogStatusGroup = iota
	statusModerated
	statusAll
)

var blogRoleStatusGroup = map[string]blogStatusGroup{
	constants.RoleKeyKoorMedcrev: statusEditable,
	constants.RoleKeyBPH:         statusModerated,
	constants.RoleKeySuperAdmin:  statusAll,
}

var (
	blogStatusEditable = map[string]struct{}{
		constants.StatusDraft:     {},
		constants.StatusPending:   {},
		constants.StatusUnpublish: {},
	}

	// untuk user validasi
	blogStatusModerated = map[string]struct{}{
		constants.StatusPublished: {},
		constants.StatusUnpublish: {},
		constants.StatusRejected:  {},
	}

	blogStatusAll = map[string]struct{}{
		constants.StatusDraft:     {},
		constants.StatusPending:   {},
		constants.StatusPublished: {},
		constants.StatusUnpublish: {},
		constants.StatusRejected:  {},
	}
)

var blogStatusDefault = map[blogStatusGroup]string{
	statusEditable:  constants.StatusDraft,
	statusModerated: constants.StatusPending,
	statusAll:       constants.StatusPending,
}

var blogRoleKeyPermissionStatus = map[string]map[string]struct{}{
	constants.RoleKeyKoorMedcrev: blogStatusEditable,
	constants.RoleKeyBPH:         blogStatusModerated,
	constants.RoleKeySuperAdmin:  blogStatusAll,
}

func setDefaultStatus(userRole string) string {
	group, ok := blogRoleStatusGroup[userRole]
	if !ok {
		return constants.StatusDraft
	}

	return blogStatusDefault[group]
}

func CanSetStatusBlog(userRole, status string) (string, error) {
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return "", fmt.Errorf("invalid user role")
	}

	if _, ok := blogStatusAll[status]; !ok {
		return "", fmt.Errorf("invalid status")
	}

	if strings.TrimSpace(status) == "" {
		defaultStatus := setDefaultStatus(userRole)
		log.Printf("status is empty, use default value")
		return defaultStatus, nil
	}

	if allowedStatus, ok := blogRoleKeyPermissionStatus[userRole]; ok {
		if _, exist := allowedStatus[status]; exist {
			return status, nil
		}
		return "", fmt.Errorf("status not allowed")
	}

	return "", fmt.Errorf("permissions denied for this role %s", userRole)
}

func CheckRolePermission(userRole string) error {

	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("role not valid %w", err)
	}

	allowedRole := map[string]bool{
		constants.RoleKeySuperAdmin:  true,
		constants.RoleKeyKoorMedcrev: true,
	}

	if !allowedRole[userRole] {
		return fmt.Errorf("role have no permission")
	}

	return nil
}
