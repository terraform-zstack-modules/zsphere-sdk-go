// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccessControlListInventoryView AccessControlList
type AccessControlListInventoryView struct {
	BaseInfoView
	BaseTimeView
	IpVersion int `json:"ipVersion,omitempty"`
	Description string `json:"description,omitempty"`
	Entries []AccessControlListEntryInventoryView `json:"entries,omitempty"`
}

// CreateAccessControlListEventView CreateAccessControlListEvent
type CreateAccessControlListEventView struct {
	Inventory AccessControlListInventoryView `json:"inventory,omitempty"`
}

// QueryAccessControlListView QueryAccessControlList
type QueryAccessControlListView struct {
	Inventories []AccessControlListInventoryView `json:"inventories,omitempty"`
}

// DeleteAccessControlListEventView DeleteAccessControlListEvent
type DeleteAccessControlListEventView struct {
	Success bool `json:"success,omitempty"`
}

