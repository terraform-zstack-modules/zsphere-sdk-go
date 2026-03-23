// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceSpecInventoryView PciDeviceSpec
type PciDeviceSpecInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	Device string `json:"device,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	RamSize string `json:"ramSize,omitempty"`
	MaxPartNum int `json:"maxPartNum,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	IsVirtual bool `json:"isVirtual,omitempty"`
	RomVersion string `json:"romVersion,omitempty"`
	RomMd5sum string `json:"romMd5sum,omitempty"`
}

// QueryPciDeviceSpecView QueryPciDeviceSpec
type QueryPciDeviceSpecView struct {
	Inventories []PciDeviceSpecInventoryView `json:"inventories,omitempty"`
}

// GetPciDeviceSpecCandidatesView GetPciDeviceSpecCandidates
type GetPciDeviceSpecCandidatesView struct {
	Inventories []PciDeviceSpecInventoryView `json:"inventories,omitempty"`
}

// UpdatePciDeviceSpecEventView UpdatePciDeviceSpecEvent
type UpdatePciDeviceSpecEventView struct {
	Inventory PciDeviceSpecInventoryView `json:"inventory,omitempty"`
}

