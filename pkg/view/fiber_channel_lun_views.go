// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FiberChannelLunInventoryView FiberChannelLun
type FiberChannelLunInventoryView struct {
	BaseInfoView
	BaseTimeView
	FiberChannelStorageUuid string `json:"fiberChannelStorageUuid,omitempty"`
	ScsiLunHostRefs []ScsiLunHostRefInventoryView `json:"scsiLunHostRefs,omitempty"`
	ScsiLunVmInstanceRefs []ScsiLunVmInstanceRefInventoryView `json:"scsiLunVmInstanceRefs,omitempty"`
	Wwid string `json:"wwid,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	Model string `json:"model,omitempty"`
	Wwn string `json:"wwn,omitempty"`
	Serial string `json:"serial,omitempty"`
	Type string `json:"type,omitempty"`
	Hctl string `json:"hctl,omitempty"`
	Path string `json:"path,omitempty"`
	State string `json:"state,omitempty"`
	Size int64 `json:"size,omitempty"`
	MultipathDeviceUuid string `json:"multipathDeviceUuid,omitempty"`
	Source string `json:"source,omitempty"`
}

// QueryFiberChannelLunView QueryFiberChannelLun
type QueryFiberChannelLunView struct {
	Inventories []FiberChannelLunInventoryView `json:"inventories,omitempty"`
}

