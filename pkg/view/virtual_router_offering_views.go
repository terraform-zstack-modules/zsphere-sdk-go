// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VirtualRouterOfferingInventoryView VirtualRouterOffering
type VirtualRouterOfferingInventoryView struct {
	BaseInfoView
	BaseTimeView
	ManagementNetworkUuid string `json:"managementNetworkUuid,omitempty"`
	PublicNetworkUuid string `json:"publicNetworkUuid,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
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

// QueryVirtualRouterOfferingView QueryVirtualRouterOffering
type QueryVirtualRouterOfferingView struct {
	Inventories []VirtualRouterOfferingInventoryView `json:"inventories,omitempty"`
}

// CreateInstanceOfferingEventView CreateInstanceOfferingEvent
type CreateInstanceOfferingEventView struct {
	Inventory InstanceOfferingInventoryView `json:"inventory,omitempty"`
}

// UpdateInstanceOfferingEventView UpdateInstanceOfferingEvent
type UpdateInstanceOfferingEventView struct {
	Inventory InstanceOfferingInventoryView `json:"inventory,omitempty"`
}

