// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SharedBlockGroupPrimaryStorageInventoryView SharedBlockGroupPrimaryStorage
type SharedBlockGroupPrimaryStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	SharedBlocks []SharedBlockInventoryView `json:"sharedBlocks,omitempty"`
	SharedBlockGroupType string `json:"sharedBlockGroupType,omitempty"`
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

// AddSharedBlockToSharedBlockGroupEventView AddSharedBlockToSharedBlockGroupEvent
type AddSharedBlockToSharedBlockGroupEventView struct {
	Inventory SharedBlockGroupPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// RefreshSharedBlockDeviceCapacityEventView RefreshSharedBlockDeviceCapacityEvent
type RefreshSharedBlockDeviceCapacityEventView struct {
	Inventory SharedBlockGroupPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// QuerySharedBlockGroupPrimaryStorageView QuerySharedBlockGroupPrimaryStorage
type QuerySharedBlockGroupPrimaryStorageView struct {
	Inventories []SharedBlockGroupPrimaryStorageInventoryView `json:"inventories,omitempty"`
}

