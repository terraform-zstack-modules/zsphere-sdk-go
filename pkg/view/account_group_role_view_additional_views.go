// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountGroupRoleViewView AccountGroupRoleView
type AccountGroupRoleViewView struct {
	GroupUuid string `json:"groupUuid,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	Roles []RoleInventoryView `json:"roles,omitempty"`
}

