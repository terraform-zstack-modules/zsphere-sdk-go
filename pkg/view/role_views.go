// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RoleInventoryView Role
type RoleInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	Policies []string `json:"policies,omitempty"`
}

// CreateRoleEventView CreateRoleEvent
type CreateRoleEventView struct {
	Inventory RoleInventoryView `json:"inventory,omitempty"`
}

// DeleteRoleEventView DeleteRoleEvent
type DeleteRoleEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryRoleView QueryRole
type QueryRoleView struct {
	Inventories []RoleInventoryView `json:"inventories,omitempty"`
}

// UpdateRoleEventView UpdateRoleEvent
type UpdateRoleEventView struct {
	Inventory RoleInventoryView `json:"inventory,omitempty"`
}

