// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeSnapshotInventoryView VolumeSnapshot
type VolumeSnapshotInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	TreeUuid string `json:"treeUuid,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	PrimaryStorageInstallPath string `json:"primaryStorageInstallPath,omitempty"`
	VolumeType string `json:"volumeType,omitempty"`
	Format string `json:"format,omitempty"`
	Latest bool `json:"latest,omitempty"`
	Size int64 `json:"size,omitempty"`
	Distance int `json:"distance,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	BackupStorageRefs []VolumeSnapshotBackupStorageRefInventoryView `json:"backupStorageRefs,omitempty"`
	GroupUuid string `json:"groupUuid,omitempty"`
}

// DeleteVolumeSnapshotEventView DeleteVolumeSnapshotEvent
type DeleteVolumeSnapshotEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateVolumesSnapshotEventView CreateVolumesSnapshotEvent
type CreateVolumesSnapshotEventView struct {
	Inventories []VolumeSnapshotInventoryView `json:"inventories,omitempty"`
}

// CreateVolumeSnapshotEventView CreateVolumeSnapshotEvent
type CreateVolumeSnapshotEventView struct {
	Inventory VolumeSnapshotInventoryView `json:"inventory,omitempty"`
}

// UpdateVolumeSnapshotEventView UpdateVolumeSnapshotEvent
type UpdateVolumeSnapshotEventView struct {
	Inventory VolumeSnapshotInventoryView `json:"inventory,omitempty"`
}

// QueryVolumeSnapshotView QueryVolumeSnapshot
type QueryVolumeSnapshotView struct {
	Inventories []VolumeSnapshotInventoryView `json:"inventories,omitempty"`
}

