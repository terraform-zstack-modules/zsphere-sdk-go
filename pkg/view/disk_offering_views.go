// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// DiskOfferingInventoryView DiskOffering
type DiskOfferingInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	DiskSize int64 `json:"diskSize,omitempty"`
	SortKey int `json:"sortKey,omitempty"`
	State string `json:"state,omitempty"`
	Type string `json:"type,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
}

// DeleteDiskOfferingEventView DeleteDiskOfferingEvent
type DeleteDiskOfferingEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeDiskOfferingStateEventView ChangeDiskOfferingStateEvent
type ChangeDiskOfferingStateEventView struct {
	Inventory DiskOfferingInventoryView `json:"inventory,omitempty"`
}

// UpdateDiskOfferingEventView UpdateDiskOfferingEvent
type UpdateDiskOfferingEventView struct {
	Inventory DiskOfferingInventoryView `json:"inventory,omitempty"`
}

// QueryDiskOfferingView QueryDiskOffering
type QueryDiskOfferingView struct {
	Inventories []DiskOfferingInventoryView `json:"inventories,omitempty"`
}

// CreateDiskOfferingEventView CreateDiskOfferingEvent
type CreateDiskOfferingEventView struct {
	Inventory DiskOfferingInventoryView `json:"inventory,omitempty"`
}

