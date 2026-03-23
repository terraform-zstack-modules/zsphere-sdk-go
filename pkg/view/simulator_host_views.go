// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SimulatorHostInventoryView SimulatorHost
type SimulatorHostInventoryView struct {
	BaseInfoView
	BaseTimeView
	MemoryCapacity int64 `json:"memoryCapacity,omitempty"`
	CpuCapacity int64 `json:"cpuCapacity,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ManagementIp string `json:"managementIp,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	TotalCpuCapacity int64 `json:"totalCpuCapacity,omitempty"`
	AvailableCpuCapacity int64 `json:"availableCpuCapacity,omitempty"`
	CpuSockets int `json:"cpuSockets,omitempty"`
	TotalMemoryCapacity int64 `json:"totalMemoryCapacity,omitempty"`
	AvailableMemoryCapacity int64 `json:"availableMemoryCapacity,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiPowerStatus string `json:"ipmiPowerStatus,omitempty"`
	CpuStatus string `json:"cpuStatus,omitempty"`
	MemoryStatus string `json:"memoryStatus,omitempty"`
	DiskStatus string `json:"diskStatus,omitempty"`
	NicStatus string `json:"nicStatus,omitempty"`
	GpuStatus string `json:"gpuStatus,omitempty"`
	PowerSupplyStatus string `json:"powerSupplyStatus,omitempty"`
	FanStatus string `json:"fanStatus,omitempty"`
	RaidStatus string `json:"raidStatus,omitempty"`
	TemperatureStatus string `json:"temperatureStatus,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Nqn string `json:"nqn,omitempty"`
	Hostname string `json:"hostname,omitempty"`
}

