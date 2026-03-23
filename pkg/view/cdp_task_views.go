// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// CdpTaskInventoryView CdpTask
type CdpTaskInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	PolicyUuid string `json:"policyUuid,omitempty"`
	BackupStorageUuid string `json:"backupStorageUuid,omitempty"`
	Status string `json:"status,omitempty"`
	State string `json:"state,omitempty"`
	TaskType string `json:"taskType,omitempty"`
	BackupBandwidth int64 `json:"backupBandwidth,omitempty"`
	MaxCapacity int64 `json:"maxCapacity,omitempty"`
	UsedCapacity int64 `json:"usedCapacity,omitempty"`
	MaxLatency int64 `json:"maxLatency,omitempty"`
	LastLatency int64 `json:"lastLatency,omitempty"`
	ResourceRefs []CdpTaskResourceRefInventoryView `json:"resourceRefs,omitempty"`
}

// CreateCdpTaskEventView CreateCdpTaskEvent
type CreateCdpTaskEventView struct {
	Inventory CdpTaskInventoryView `json:"inventory,omitempty"`
}

// QueryCdpTaskView QueryCdpTask
type QueryCdpTaskView struct {
	Inventories []CdpTaskInventoryView `json:"inventories,omitempty"`
}

// DeleteCdpTaskEventView DeleteCdpTaskEvent
type DeleteCdpTaskEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateCdpTaskEventView UpdateCdpTaskEvent
type UpdateCdpTaskEventView struct {
	Inventory CdpTaskInventoryView `json:"inventory,omitempty"`
}

