// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// RoleAccountRefInventoryView RoleAccountRef
type RoleAccountRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	RoleUuid string `json:"roleUuid,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	AccountPermissionFrom string `json:"accountPermissionFrom,omitempty"`
}

// QueryRoleAccountRefView QueryRoleAccountRef
type QueryRoleAccountRefView struct {
	Inventories []RoleAccountRefInventoryView `json:"inventories,omitempty"`
}

