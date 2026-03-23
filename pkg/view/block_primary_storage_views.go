// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// BlockPrimaryStorageInventoryView BlockPrimaryStorage
type BlockPrimaryStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	VendorName string `json:"vendorName,omitempty"`
	Metadata string `json:"metadata,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	TotalPhysicalCapacity int64 `json:"totalPhysicalCapacity,omitempty"`
	AvailablePhysicalCapacity int64 `json:"availablePhysicalCapacity,omitempty"`
	SystemUsedCapacity int64 `json:"systemUsedCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
}

// UpdateBlockPrimaryStorageEventView UpdateBlockPrimaryStorageEvent
type UpdateBlockPrimaryStorageEventView struct {
	Inventory BlockPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// QueryBlockPrimaryStorageView QueryBlockPrimaryStorage
type QueryBlockPrimaryStorageView struct {
	Inventories []BlockPrimaryStorageInventoryView `json:"inventories,omitempty"`
}

