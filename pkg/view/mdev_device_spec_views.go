// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MdevDeviceSpecInventoryView MdevDeviceSpec
type MdevDeviceSpecInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Specification string `json:"specification,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
}

// UpdateMdevDeviceSpecEventView UpdateMdevDeviceSpecEvent
type UpdateMdevDeviceSpecEventView struct {
	Inventory MdevDeviceSpecInventoryView `json:"inventory,omitempty"`
}

// QueryMdevDeviceSpecView QueryMdevDeviceSpec
type QueryMdevDeviceSpecView struct {
	Inventories []MdevDeviceSpecInventoryView `json:"inventories,omitempty"`
}

// GetMdevDeviceSpecCandidatesView GetMdevDeviceSpecCandidates
type GetMdevDeviceSpecCandidatesView struct {
	Inventories []MdevDeviceSpecInventoryView `json:"inventories,omitempty"`
}

