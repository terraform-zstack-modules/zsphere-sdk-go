// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MdevDeviceInventoryView MdevDevice
type MdevDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	MttyUuid string `json:"mttyUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	MdevSpecUuid string `json:"mdevSpecUuid,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Chooser string `json:"chooser,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

// GetMdevDeviceCandidatesView GetMdevDeviceCandidates
type GetMdevDeviceCandidatesView struct {
	Inventories []MdevDeviceInventoryView `json:"inventories,omitempty"`
}

// AttachMdevDeviceToVmEventView AttachMdevDeviceToVmEvent
type AttachMdevDeviceToVmEventView struct {
	Inventory MdevDeviceInventoryView `json:"inventory,omitempty"`
}

// UpdateMdevDeviceEventView UpdateMdevDeviceEvent
type UpdateMdevDeviceEventView struct {
	Inventory MdevDeviceInventoryView `json:"inventory,omitempty"`
}

// DeleteMdevDeviceEventView DeleteMdevDeviceEvent
type DeleteMdevDeviceEventView struct {
	Success bool `json:"success,omitempty"`
}

// DetachMdevDeviceFromVmEventView DetachMdevDeviceFromVmEvent
type DetachMdevDeviceFromVmEventView struct {
	Inventory MdevDeviceInventoryView `json:"inventory,omitempty"`
}

// QueryMdevDeviceView QueryMdevDevice
type QueryMdevDeviceView struct {
	Inventories []MdevDeviceInventoryView `json:"inventories,omitempty"`
}

