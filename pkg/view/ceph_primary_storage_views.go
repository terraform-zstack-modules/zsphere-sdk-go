// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CephPrimaryStorageInventoryView CephPrimaryStorage
type CephPrimaryStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Mons []CephPrimaryStorageMonInventoryView `json:"mons,omitempty"`
	Pools []CephPrimaryStoragePoolInventoryView `json:"pools,omitempty"`
	Fsid string `json:"fsid,omitempty"`
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

// RemoveMonFromCephPrimaryStorageEventView RemoveMonFromCephPrimaryStorageEvent
type RemoveMonFromCephPrimaryStorageEventView struct {
	Inventory CephPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

// AddMonToCephPrimaryStorageEventView AddMonToCephPrimaryStorageEvent
type AddMonToCephPrimaryStorageEventView struct {
	Inventory CephPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

