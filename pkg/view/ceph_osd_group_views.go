// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CephOsdGroupInventoryView CephOsdGroup
type CephOsdGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	Osds string `json:"osds,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
}

// QueryCephOsdGroupView QueryCephOsdGroup
type QueryCephOsdGroupView struct {
	Inventories []CephOsdGroupInventoryView `json:"inventories,omitempty"`
}

