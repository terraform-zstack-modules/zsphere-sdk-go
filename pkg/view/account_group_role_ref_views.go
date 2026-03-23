// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountGroupRoleRefInventoryView AccountGroupRoleRef
type AccountGroupRoleRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	GroupUuid string `json:"groupUuid,omitempty"`
	RoleUuid string `json:"roleUuid,omitempty"`
}

