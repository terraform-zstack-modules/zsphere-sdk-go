// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDevicePciDeviceOfferingRefInventoryView PciDevicePciDeviceOfferingRef
type PciDevicePciDeviceOfferingRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	PciDeviceUuid string `json:"pciDeviceUuid,omitempty"`
	PciDeviceOfferingUuid string `json:"pciDeviceOfferingUuid,omitempty"`
}

// QueryPciDevicePciDeviceOfferingView QueryPciDevicePciDeviceOffering
type QueryPciDevicePciDeviceOfferingView struct {
	Inventories []PciDevicePciDeviceOfferingRefInventoryView `json:"inventories,omitempty"`
}

