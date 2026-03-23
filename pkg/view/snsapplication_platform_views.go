// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SNSApplicationPlatformInventoryView SNSApplicationPlatform
type SNSApplicationPlatformInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
}

// DeleteSNSApplicationPlatformEventView DeleteSNSApplicationPlatformEvent
type DeleteSNSApplicationPlatformEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeSNSApplicationPlatformStateEventView ChangeSNSApplicationPlatformStateEvent
type ChangeSNSApplicationPlatformStateEventView struct {
	Inventory SNSApplicationPlatformInventoryView `json:"inventory,omitempty"`
}

// CreateSNSApplicationPlatformEventView CreateSNSApplicationPlatformEvent
type CreateSNSApplicationPlatformEventView struct {
	Inventory SNSApplicationPlatformInventoryView `json:"inventory,omitempty"`
}

// UpdateSNSApplicationPlatformEventView UpdateSNSApplicationPlatformEvent
type UpdateSNSApplicationPlatformEventView struct {
	Inventory SNSApplicationPlatformInventoryView `json:"inventory,omitempty"`
}

// QuerySNSApplicationPlatformView QuerySNSApplicationPlatform
type QuerySNSApplicationPlatformView struct {
	Inventories []SNSApplicationPlatformInventoryView `json:"inventories,omitempty"`
}

