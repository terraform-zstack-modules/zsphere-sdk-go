// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmInstancePciDeviceSpecRefInventoryView VmInstancePciDeviceSpecRef
type VmInstancePciDeviceSpecRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	PciSpecUuid string `json:"pciSpecUuid,omitempty"`
	PciDeviceNumber int `json:"pciDeviceNumber,omitempty"`
}

// QueryVmInstancePciDeviceSpecRefView QueryVmInstancePciDeviceSpecRef
type QueryVmInstancePciDeviceSpecRefView struct {
	Inventories []VmInstancePciDeviceSpecRefInventoryView `json:"inventories,omitempty"`
}

// AddPciDeviceSpecToVmInstanceEventView AddPciDeviceSpecToVmInstanceEvent
type AddPciDeviceSpecToVmInstanceEventView struct {
	Inventory VmInstancePciDeviceSpecRefInventoryView `json:"inventory,omitempty"`
}

