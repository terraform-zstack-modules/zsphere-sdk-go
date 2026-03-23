// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeSnapshotReferenceInventoryView VolumeSnapshotReference
type VolumeSnapshotReferenceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	ParentId int64 `json:"parentId,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	VolumeSnapshotUuid string `json:"volumeSnapshotUuid,omitempty"`
	VolumeSnapshotInstallUrl string `json:"volumeSnapshotInstallUrl,omitempty"`
	DirectSnapshotUuid string `json:"directSnapshotUuid,omitempty"`
	DirectSnapshotInstallUrl string `json:"directSnapshotInstallUrl,omitempty"`
	ReferenceUuid string `json:"referenceUuid,omitempty"`
	ReferenceType string `json:"referenceType,omitempty"`
	ReferenceInstallUrl string `json:"referenceInstallUrl,omitempty"`
	ReferenceVolumeUuid string `json:"referenceVolumeUuid,omitempty"`
}

