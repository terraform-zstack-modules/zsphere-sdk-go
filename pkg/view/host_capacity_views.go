// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostCapacityInventoryView HostCapacity
type HostCapacityInventoryView struct {
	BaseInfoView
	BaseTimeView
	TotalMemory int64 `json:"totalMemory,omitempty"`
	TotalCpu int64 `json:"totalCpu,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	CpuSockets int `json:"cpuSockets,omitempty"`
	CpuCoreNum int `json:"cpuCoreNum,omitempty"`
	AvailableMemory int64 `json:"availableMemory,omitempty"`
	AvailableCpu int64 `json:"availableCpu,omitempty"`
	TotalPhysicalMemory int64 `json:"totalPhysicalMemory,omitempty"`
	AvailablePhysicalMemory int64 `json:"availablePhysicalMemory,omitempty"`
}

