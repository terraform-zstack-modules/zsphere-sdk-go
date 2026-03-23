// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SlbOfferingInventoryView SlbOffering
type SlbOfferingInventoryView struct {
	BaseInfoView
	BaseTimeView
	ManagementNetworkUuid string `json:"managementNetworkUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
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

// QuerySlbOfferingView QuerySlbOffering
type QuerySlbOfferingView struct {
	Inventories []SlbOfferingInventoryView `json:"inventories,omitempty"`
}

