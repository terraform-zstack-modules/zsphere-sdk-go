// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RolePolicyInventoryView RolePolicy
type RolePolicyInventoryView struct {
	BaseInfoView
	BaseTimeView
	RoleUuid string `json:"roleUuid,omitempty"`
	Actions string `json:"actions,omitempty"`
	Effect string `json:"effect,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
}

