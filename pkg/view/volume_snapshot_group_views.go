// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeSnapshotGroupInventoryView VolumeSnapshotGroup
type VolumeSnapshotGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	SnapshotCount int `json:"snapshotCount,omitempty"`
	Description string `json:"description,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	VolumeSnapshotRefs []VolumeSnapshotGroupRefInventoryView `json:"volumeSnapshotRefs,omitempty"`
}

// CreateVolumeSnapshotGroupEventView CreateVolumeSnapshotGroupEvent
type CreateVolumeSnapshotGroupEventView struct {
	Inventory VolumeSnapshotGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateVolumeSnapshotGroupEventView UpdateVolumeSnapshotGroupEvent
type UpdateVolumeSnapshotGroupEventView struct {
	Inventory VolumeSnapshotGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteVolumeSnapshotGroupEventView DeleteVolumeSnapshotGroupEvent
type DeleteVolumeSnapshotGroupEventView struct {
	Results []DeleteSnapshotGroupResultView `json:"results,omitempty"`
}

// QueryVolumeSnapshotGroupView QueryVolumeSnapshotGroup
type QueryVolumeSnapshotGroupView struct {
	Inventories []VolumeSnapshotGroupInventoryView `json:"inventories,omitempty"`
}

// GetMemorySnapshotGroupReferenceView GetMemorySnapshotGroupReference
type GetMemorySnapshotGroupReferenceView struct {
	Inventories []VolumeSnapshotGroupInventoryView `json:"inventories,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

