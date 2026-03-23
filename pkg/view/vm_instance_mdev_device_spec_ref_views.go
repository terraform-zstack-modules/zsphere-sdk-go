// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmInstanceMdevDeviceSpecRefInventoryView VmInstanceMdevDeviceSpecRef
type VmInstanceMdevDeviceSpecRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	MdevSpecUuid string `json:"mdevSpecUuid,omitempty"`
	MdevDeviceNumber int `json:"mdevDeviceNumber,omitempty"`
}

// QueryVmInstanceMdevDeviceSpecRefView QueryVmInstanceMdevDeviceSpecRef
type QueryVmInstanceMdevDeviceSpecRefView struct {
	Inventories []VmInstanceMdevDeviceSpecRefInventoryView `json:"inventories,omitempty"`
}

// AddMdevDeviceSpecToVmInstanceEventView AddMdevDeviceSpecToVmInstanceEvent
type AddMdevDeviceSpecToVmInstanceEventView struct {
	Inventory VmInstanceMdevDeviceSpecRefInventoryView `json:"inventory,omitempty"`
}

