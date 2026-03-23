// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ScsiLunVmInstanceRefInventoryView ScsiLunVmInstanceRef
type ScsiLunVmInstanceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ScsiLunUuid string `json:"scsiLunUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	AttachMultipath bool `json:"attachMultipath,omitempty"`
}

