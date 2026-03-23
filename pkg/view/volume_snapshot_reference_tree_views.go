// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeSnapshotReferenceTreeInventoryView VolumeSnapshotReferenceTree
type VolumeSnapshotReferenceTreeInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	RootImageUuid string `json:"rootImageUuid,omitempty"`
	RootVolumeUuid string `json:"rootVolumeUuid,omitempty"`
	RootVolumeSnapshotUuid string `json:"rootVolumeSnapshotUuid,omitempty"`
	RootVolumeSnapshotTreeUuid string `json:"rootVolumeSnapshotTreeUuid,omitempty"`
	RootInstallUrl string `json:"rootInstallUrl,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
}

