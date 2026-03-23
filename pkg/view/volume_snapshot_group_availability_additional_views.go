// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VolumeSnapshotGroupAvailabilityView VolumeSnapshotGroupAvailability
type VolumeSnapshotGroupAvailabilityView struct {
	Uuid string `json:"uuid,omitempty"`
	Available bool `json:"available,omitempty"`
	Reason string `json:"reason,omitempty"`
}

