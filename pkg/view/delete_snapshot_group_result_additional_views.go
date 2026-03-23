// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DeleteSnapshotGroupResultView DeleteSnapshotGroupResult
type DeleteSnapshotGroupResultView struct {
	SnapshotUuid string `json:"snapshotUuid,omitempty"`
	VolumeUuid string `json:"volumeUuid,omitempty"`
	Success bool `json:"success,omitempty"`
}

