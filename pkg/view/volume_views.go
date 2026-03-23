// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeInventoryView Volume
type VolumeInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	DiskOfferingUuid string `json:"diskOfferingUuid,omitempty"`
	RootImageUuid string `json:"rootImageUuid,omitempty"`
	InstallPath string `json:"installPath,omitempty"`
	Type string `json:"type,omitempty"`
	Format string `json:"format,omitempty"`
	Size int64 `json:"size,omitempty"`
	ActualSize int64 `json:"actualSize,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	IsShareable bool `json:"isShareable,omitempty"`
	VolumeQos string `json:"volumeQos,omitempty"`
	LastDetachDate time.Time `json:"lastDetachDate,omitempty"`
	LastVmInstanceUuid string `json:"lastVmInstanceUuid,omitempty"`
	LastAttachDate time.Time `json:"lastAttachDate,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

// QueryVolumeView QueryVolume
type QueryVolumeView struct {
	Inventories []VolumeInventoryView `json:"inventories,omitempty"`
}

// ResizeDataVolumeEventView ResizeDataVolumeEvent
type ResizeDataVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// GetVmAttachableDataVolumeView GetVmAttachableDataVolume
type GetVmAttachableDataVolumeView struct {
	Inventories []VolumeInventoryView `json:"inventories,omitempty"`
}

// CreateDataVolumeFromVolumeSnapshotEventView CreateDataVolumeFromVolumeSnapshotEvent
type CreateDataVolumeFromVolumeSnapshotEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// RecoverDataVolumeEventView RecoverDataVolumeEvent
type RecoverDataVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// CreateDataVolumeEventView CreateDataVolumeEvent
type CreateDataVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// DeleteVolumeQosEventView DeleteVolumeQosEvent
type DeleteVolumeQosEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// SyncVolumeSizeEventView SyncVolumeSizeEvent
type SyncVolumeSizeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// UpdateVolumeEventView UpdateVolumeEvent
type UpdateVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// ResizeRootVolumeEventView ResizeRootVolumeEvent
type ResizeRootVolumeEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// DetachDataVolumeFromVmEventView DetachDataVolumeFromVmEvent
type DetachDataVolumeFromVmEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// CreateDataVolumeFromVolumeBackupEventView CreateDataVolumeFromVolumeBackupEvent
type CreateDataVolumeFromVolumeBackupEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// CreateDataVolumeFromVolumeTemplateEventView CreateDataVolumeFromVolumeTemplateEvent
type CreateDataVolumeFromVolumeTemplateEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// ChangeVolumeStateEventView ChangeVolumeStateEvent
type ChangeVolumeStateEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// AttachDataVolumeToVmEventView AttachDataVolumeToVmEvent
type AttachDataVolumeToVmEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

// SetVolumeQosEventView SetVolumeQosEvent
type SetVolumeQosEventView struct {
	Inventory VolumeInventoryView `json:"inventory,omitempty"`
}

