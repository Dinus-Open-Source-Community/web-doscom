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

var workStatusToGroup = map[string]workStatusGroup{
	constants.StatusDraft:     statusEditable,
	constants.StatusPending:   statusEditable,
	constants.StatusPublished: statusModerated,
	constants.StatusUnpublish: statusModerated,
	constants.StatusRejected:  statusModerated,
}
var workModeratedTransition = map[string]map[string]struct{}{
	constants.StatusPending: {
		constants.StatusRejected:  {},
		constants.StatusPublished: {},
	},
	constants.StatusPublished: {
		constants.StatusUnpublish: {},
	},
	constants.StatusUnpublish: {
		constants.StatusPublished: {},
	},
}

var deletableWorkStatusByGroup = map[workStatusGroup]map[string]struct{}{
	statusEditable: {
		constants.StatusDraft:    {},
		constants.StatusRejected: {},
	},
	statusModerated: {
		constants.StatusPublished: {},
		constants.StatusUnpublish: {},
	},
	statusAll: {
		constants.StatusPublished: {},
		constants.StatusUnpublish: {},
		constants.StatusRejected:  {},
	},
}

var workGroupAllowedStatus = map[workStatusGroup]map[string]struct{}{}
var workGroupViewStatus = map[workStatusGroup]map[string]struct{}{}

func copyMap(oldMap map[string]struct{}) map[string]struct{} {
	newMap := make(map[string]struct{}, len(oldMap))

	for key, value := range oldMap {
		newMap[key] = value
	}

	return newMap
}

func init() {
	workGroupAllowedStatus = map[workStatusGroup]map[string]struct{}{
		statusEditable:  {},
		statusModerated: {},
		statusAll:       {},
	}

	for status, group := range workStatusToGroup {
		workGroupAllowedStatus[group][status] = struct{}{}
		workGroupAllowedStatus[statusAll][status] = struct{}{}
	}

	workGroupViewStatus = map[workStatusGroup]map[string]struct{}{
		statusEditable: copyMap(workGroupAllowedStatus[statusAll]),
		statusModerated: {
			constants.StatusPending:   {},
			constants.StatusPublished: {},
			constants.StatusUnpublish: {},
		},
		statusAll: copyMap(workGroupAllowedStatus[statusAll]),
	}

}

var workStatusDefault = map[workStatusGroup]string{
	statusEditable:  constants.StatusDraft,
	statusModerated: constants.StatusPending,
	statusAll:       constants.StatusPending,
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

	status = strings.ToLower(strings.TrimSpace(status))
	if _, validGlobal := workGroupAllowedStatus[statusAll][status]; !validGlobal {
		return false, fmt.Errorf("status not valid")
	}

	allowedDelete, ok := deletableWorkStatusByGroup[roleStatusGroup]
	if !ok {
		return false, fmt.Errorf("delete rule not found for role %s", userRole)
	}
	if _, allowed := allowedDelete[status]; !allowed {
		return false, fmt.Errorf("role %s cannot delete work with status %s", userRole, status)
	}

	return true, nil
}

func GetViewableStatus(userRole string) ([]string, error) {
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return nil, fmt.Errorf("invalid userRole")
	}

	group, ok := getRoleGroup(userRole)
	if !ok {
		return nil, fmt.Errorf("permissions denied for this role %s", userRole)
	}

	viewStatus, ok := workGroupViewStatus[group]
	if !ok {
		return nil, fmt.Errorf("tidak ada status untuk role ini")
	}

	resultStatus := make([]string, 0, len(viewStatus))
	for status := range viewStatus {
		resultStatus = append(resultStatus, status)
	}

	return resultStatus, nil
}

func CanModerateWork(userRole, currentStatus, targetStatus string) error {
	_, err := authorization.GetRoleInfo(userRole)
	if err != nil {
		return fmt.Errorf("invalid userRole")
	}

	roleStatusGroup, ok := workRoleStatusGroup[userRole]
	if !ok {
		return fmt.Errorf("permissions denied for this role %s", userRole)
	}

	currentStatus = strings.ToLower(strings.TrimSpace(currentStatus))
	targetStatus = strings.ToLower(strings.TrimSpace(targetStatus))

	if _, validGlobal := workGroupAllowedStatus[statusAll][currentStatus]; !validGlobal {
		return fmt.Errorf("current status not valid")
	}
	if _, validGlobal := workGroupAllowedStatus[statusAll][targetStatus]; !validGlobal {
		return fmt.Errorf("target status not valid")
	}

	targetStatusGroup, ok := workStatusToGroup[targetStatus]
	if !ok {
		return fmt.Errorf("target status group not found/valid")
	}

	if roleStatusGroup < targetStatusGroup {
		return fmt.Errorf("status not allowed for this role %s", userRole)
	}

	// check if transition allowed or not
	allowedTargets, ok := workModeratedTransition[currentStatus]
	if !ok {
		return fmt.Errorf("status %s cannot be moderated/transition", currentStatus)
	}
	if _, allowed := allowedTargets[targetStatus]; !allowed {
		return fmt.Errorf("cannot change status from %s to %s", currentStatus, targetStatus)
	}
	return nil
}
