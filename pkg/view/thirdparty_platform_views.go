// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ThirdpartyPlatformInventoryView ThirdpartyPlatform
type ThirdpartyPlatformInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Template string `json:"template,omitempty"`
	State string `json:"state,omitempty"`
	Description string `json:"description,omitempty"`
	LastSyncDate time.Time `json:"lastSyncDate,omitempty"`
}

// UpdateThirdpartyPlatformEventView UpdateThirdpartyPlatformEvent
type UpdateThirdpartyPlatformEventView struct {
	Inventory ThirdpartyPlatformInventoryView `json:"inventory,omitempty"`
}

// AddThirdpartyPlatformEventView AddThirdpartyPlatformEvent
type AddThirdpartyPlatformEventView struct {
	Inventory ThirdpartyPlatformInventoryView `json:"inventory,omitempty"`
}

// QueryThirdpartyPlatformView QueryThirdpartyPlatform
type QueryThirdpartyPlatformView struct {
	Inventories []ThirdpartyPlatformInventoryView `json:"inventories,omitempty"`
}

// DeleteThirdpartyPlatformEventView DeleteThirdpartyPlatformEvent
type DeleteThirdpartyPlatformEventView struct {
	Success bool `json:"success,omitempty"`
}

