// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeSnapshotTreeInventoryView VolumeSnapshotTree
type VolumeSnapshotTreeInventoryView struct {
	BaseInfoView
	BaseTimeView
	VolumeUuid string `json:"volumeUuid,omitempty"`
	Current bool `json:"current,omitempty"`
	Status string `json:"status,omitempty"`
	Tree SnapshotLeafInventoryView `json:"tree,omitempty"`
}

// QueryVolumeSnapshotTreeView QueryVolumeSnapshotTree
type QueryVolumeSnapshotTreeView struct {
	Inventories []VolumeSnapshotTreeInventoryView `json:"inventories,omitempty"`
}

