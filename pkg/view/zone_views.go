// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZoneInventoryView Zone
type ZoneInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
}

// DeleteZoneEventView DeleteZoneEvent
type DeleteZoneEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeZoneStateEventView ChangeZoneStateEvent
type ChangeZoneStateEventView struct {
	Inventory ZoneInventoryView `json:"inventory,omitempty"`
}

// UpdateZoneEventView UpdateZoneEvent
type UpdateZoneEventView struct {
	Inventory ZoneInventoryView `json:"inventory,omitempty"`
}

// QueryZoneView QueryZone
type QueryZoneView struct {
	Inventories []ZoneInventoryView `json:"inventories,omitempty"`
}

// GetZoneView GetZone
type GetZoneView struct {
	Inventories []ZoneInventoryView `json:"inventories,omitempty"`
}

// CreateZoneEventView CreateZoneEvent
type CreateZoneEventView struct {
	Inventory ZoneInventoryView `json:"inventory,omitempty"`
}

