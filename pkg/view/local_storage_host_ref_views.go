// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LocalStorageHostRefInventoryView LocalStorageHostRef
type LocalStorageHostRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	SystemUsedCapacity int64 `json:"systemUsedCapacity,omitempty"`
}

