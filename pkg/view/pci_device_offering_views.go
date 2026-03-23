// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceOfferingInventoryView PciDeviceOffering
type PciDeviceOfferingInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	RamSize string `json:"ramSize,omitempty"`
	AttachedInstanceOfferings []PciDeviceOfferingInstanceOfferingRefInventoryView `json:"attachedInstanceOfferings,omitempty"`
	MatchedPciDevices []PciDevicePciDeviceOfferingRefInventoryView `json:"matchedPciDevices,omitempty"`
}

// DeletePciDeviceOfferingEventView DeletePciDeviceOfferingEvent
type DeletePciDeviceOfferingEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreatePciDeviceOfferingEventView CreatePciDeviceOfferingEvent
type CreatePciDeviceOfferingEventView struct {
	Inventory PciDeviceOfferingInventoryView `json:"inventory,omitempty"`
}

// QueryPciDeviceOfferingView QueryPciDeviceOffering
type QueryPciDeviceOfferingView struct {
	Inventories []PciDeviceOfferingInventoryView `json:"inventories,omitempty"`
}

