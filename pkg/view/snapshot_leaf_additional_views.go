// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SnapshotLeafInventoryView SnapshotLeaf
type SnapshotLeafInventoryView struct {
	BaseInfoView
	BaseTimeView
	Inventory VolumeSnapshotInventoryView `json:"inventory,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	Children []*SnapshotLeafInventoryView `json:"children,omitempty"`
}

