// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SharedBlockInventoryView SharedBlock
type SharedBlockInventoryView struct {
	BaseInfoView
	BaseTimeView
	SharedBlockGroupUuid string `json:"sharedBlockGroupUuid,omitempty"`
	Type string `json:"type,omitempty"`
	DiskUuid string `json:"diskUuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	TotalCapacity int64 `json:"totalCapacity,omitempty"`
	AvailableCapacity int64 `json:"availableCapacity,omitempty"`
}

// QuerySharedBlockView QuerySharedBlock
type QuerySharedBlockView struct {
	Inventories []SharedBlockInventoryView `json:"inventories,omitempty"`
}

// UpdateSharedBlockEventView UpdateSharedBlockEvent
type UpdateSharedBlockEventView struct {
	Inventory SharedBlockGroupPrimaryStorageInventoryView `json:"inventory,omitempty"`
}

