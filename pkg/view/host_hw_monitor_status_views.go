// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostHwMonitorStatusInventoryView HostHwMonitorStatus
type HostHwMonitorStatusInventoryView struct {
	BaseInfoView
	BaseTimeView
	CpuStatus string `json:"cpuStatus,omitempty"`
	MemoryStatus string `json:"memoryStatus,omitempty"`
	DiskStatus string `json:"diskStatus,omitempty"`
	FanStatus string `json:"fanStatus,omitempty"`
	PowerSupplyStatus string `json:"powerSupplyStatus,omitempty"`
	RaidStatus string `json:"raidStatus,omitempty"`
	NetworkStatus string `json:"networkStatus,omitempty"`
	GpuStatus string `json:"gpuStatus,omitempty"`
	TemperatureStatus string `json:"temperatureStatus,omitempty"`
}

