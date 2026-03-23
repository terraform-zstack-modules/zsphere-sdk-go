// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountResourceRefInventoryView AccountResourceRef
type AccountResourceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	AccountUuid string `json:"accountUuid,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	AccountPermissionFrom string `json:"accountPermissionFrom,omitempty"`
	ResourcePermissionFrom string `json:"resourcePermissionFrom,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryAccountResourceRefView QueryAccountResourceRef
type QueryAccountResourceRefView struct {
	Inventories []AccountResourceRefInventoryView `json:"inventories,omitempty"`
}

// ChangeResourceOwnerEventView ChangeResourceOwnerEvent
type ChangeResourceOwnerEventView struct {
	Inventory AccountResourceRefInventoryView `json:"inventory,omitempty"`
}

