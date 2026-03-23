// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BatchDeleteVolumeSnapshotStructView BatchDeleteVolumeSnapshotStruct
type BatchDeleteVolumeSnapshotStructView struct {
	SnapshotUuid string `json:"snapshotUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

