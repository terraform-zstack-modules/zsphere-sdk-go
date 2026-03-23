// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ShareableVolumeVmInstanceRefInventoryView ShareableVolumeVmInstanceRef
type ShareableVolumeVmInstanceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
}

// QueryShareableVolumeVmInstanceRefView QueryShareableVolumeVmInstanceRef
type QueryShareableVolumeVmInstanceRefView struct {
	Inventories []ShareableVolumeVmInstanceRefInventoryView `json:"inventories,omitempty"`
}

