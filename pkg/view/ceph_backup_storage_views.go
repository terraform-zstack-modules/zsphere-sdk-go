// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CephBackupStorageInventoryView CephBackupStorage
type CephBackupStorageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Mons []CephBackupStorageMonInventoryView `json:"mons,omitempty"`
	Fsid string `json:"fsid,omitempty"`
	PoolName string `json:"poolName,omitempty"`
	PoolAvailableCapacity int64 `json:"poolAvailableCapacity,omitempty"`
	PoolUsedCapacity int64 `json:"poolUsedCapacity,omitempty"`
	PoolReplicatedSize int `json:"poolReplicatedSize,omitempty"`
	PoolDiskUtilization float32 `json:"poolDiskUtilization,omitempty"`
	PoolSecurityPolicy string `json:"poolSecurityPolicy,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	AttachedZoneUuids []string `json:"attachedZoneUuids,omitempty"`
}

// RemoveMonFromCephBackupStorageEventView RemoveMonFromCephBackupStorageEvent
type RemoveMonFromCephBackupStorageEventView struct {
	Inventory CephBackupStorageInventoryView `json:"inventory,omitempty"`
}

// AddMonToCephBackupStorageEventView AddMonToCephBackupStorageEvent
type AddMonToCephBackupStorageEventView struct {
	Inventory CephBackupStorageInventoryView `json:"inventory,omitempty"`
}

