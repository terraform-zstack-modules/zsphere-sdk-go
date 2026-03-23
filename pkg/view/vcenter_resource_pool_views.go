// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VCenterResourcePoolInventoryView VCenterResourcePool
type VCenterResourcePoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	VCenterClusterUuid string `json:"vCenterClusterUuid,omitempty"`
	MorVal string `json:"morVal,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	CPULimit int64 `json:"CPULimit,omitempty"`
	CPUOverheadLimit int64 `json:"CPUOverheadLimit,omitempty"`
	CPUReservation int64 `json:"CPUReservation,omitempty"`
	CPUShares int64 `json:"CPUShares,omitempty"`
	CPULevel string `json:"CPULevel,omitempty"`
	MemoryLimit int64 `json:"memoryLimit,omitempty"`
	MemoryOverheadLimit int64 `json:"memoryOverheadLimit,omitempty"`
	MemoryReservation int64 `json:"memoryReservation,omitempty"`
	MemoryShares int64 `json:"memoryShares,omitempty"`
	MemoryLevel string `json:"memoryLevel,omitempty"`
	SubResources []VCenterResourcePoolUsageInventoryView `json:"subResources,omitempty"`
}

// QueryVCenterResourcePoolView QueryVCenterResourcePool
type QueryVCenterResourcePoolView struct {
	Inventories []VCenterResourcePoolInventoryView `json:"inventories,omitempty"`
}

