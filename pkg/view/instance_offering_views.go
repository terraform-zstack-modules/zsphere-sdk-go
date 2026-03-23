// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// InstanceOfferingInventoryView InstanceOffering
type InstanceOfferingInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	CpuSpeed int `json:"cpuSpeed,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	Type string `json:"type,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
	SortKey int `json:"sortKey,omitempty"`
	State string `json:"state,omitempty"`
}

// ChangeInstanceOfferingStateEventView ChangeInstanceOfferingStateEvent
type ChangeInstanceOfferingStateEventView struct {
	Inventory InstanceOfferingInventoryView `json:"inventory,omitempty"`
}

// QueryInstanceOfferingView QueryInstanceOffering
type QueryInstanceOfferingView struct {
	Inventories []InstanceOfferingInventoryView `json:"inventories,omitempty"`
}

// DeleteInstanceOfferingEventView DeleteInstanceOfferingEvent
type DeleteInstanceOfferingEventView struct {
	Success bool `json:"success,omitempty"`
}

