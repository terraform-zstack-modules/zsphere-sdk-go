// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SnapshotUsageInventoryView SnapshotUsage
type SnapshotUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	SnapshotUuid string `json:"SnapshotUuid,omitempty"`
	SnapshotStatus string `json:"SnapshotStatus,omitempty"`
	SnapshotName string `json:"SnapshotName,omitempty"`
	SnapshotSize int64 `json:"SnapshotSize,omitempty"`
	Inventory string `json:"inventory,omitempty"`
}

