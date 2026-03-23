// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccountGroupInventoryView AccountGroup
type AccountGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	RootGroupUuid string `json:"rootGroupUuid,omitempty"`
}

// CreateAccountGroupEventView CreateAccountGroupEvent
type CreateAccountGroupEventView struct {
	Inventory AccountGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateAccountGroupEventView UpdateAccountGroupEvent
type UpdateAccountGroupEventView struct {
	Inventory AccountGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteAccountGroupEventView DeleteAccountGroupEvent
type DeleteAccountGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryAccountGroupView QueryAccountGroup
type QueryAccountGroupView struct {
	Inventories []AccountGroupInventoryView `json:"inventories,omitempty"`
}

