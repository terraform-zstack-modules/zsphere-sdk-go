// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CephPrimaryStoragePoolInventoryView CephPrimaryStoragePool
type CephPrimaryStoragePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryStorageUuid string `json:"primaryStorageUuid,omitempty"`
	PoolName string `json:"poolName,omitempty"`
	AliasName string `json:"aliasName,omitempty"`
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
	UsedCapacity int64 `json:"usedCapacity,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	SecurityPolicy string `json:"securityPolicy,omitempty"`
	ReplicatedSize int `json:"replicatedSize,omitempty"`
	DiskUtilization float32 `json:"diskUtilization,omitempty"`
}

// AddCephPrimaryStoragePoolEventView AddCephPrimaryStoragePoolEvent
type AddCephPrimaryStoragePoolEventView struct {
	Inventory CephPrimaryStoragePoolInventoryView `json:"inventory,omitempty"`
}

// QueryCephPrimaryStoragePoolView QueryCephPrimaryStoragePool
type QueryCephPrimaryStoragePoolView struct {
	Inventories []CephPrimaryStoragePoolInventoryView `json:"inventories,omitempty"`
}

// DeleteCephPrimaryStoragePoolEventView DeleteCephPrimaryStoragePoolEvent
type DeleteCephPrimaryStoragePoolEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateCephPrimaryStoragePoolEventView UpdateCephPrimaryStoragePoolEvent
type UpdateCephPrimaryStoragePoolEventView struct {
	Inventory CephPrimaryStoragePoolInventoryView `json:"inventory,omitempty"`
}

