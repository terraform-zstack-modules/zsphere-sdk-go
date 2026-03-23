// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostInventoryView Host
type HostInventoryView struct {
	BaseInfoView
	BaseTimeView
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

// UpdateHostEventView UpdateHostEvent
type UpdateHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
}

// ChangeHostStateEventView ChangeHostStateEvent
type ChangeHostStateEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
}

// AddHostEventView AddHostEvent
type AddHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
}

// UpdateHostnameEventView UpdateHostnameEvent
type UpdateHostnameEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UpdateHostNqnEventView UpdateHostNqnEvent
type UpdateHostNqnEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// ReconnectHostEventView ReconnectHostEvent
type ReconnectHostEventView struct {
	Inventory HostInventoryView `json:"inventory,omitempty"`
}

// DeleteHostEventView DeleteHostEvent
type DeleteHostEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryHostView QueryHost
type QueryHostView struct {
	Inventories []HostInventoryView `json:"inventories,omitempty"`
}

// GetVmMigrationCandidateHostsView GetVmMigrationCandidateHosts
type GetVmMigrationCandidateHostsView struct {
	Inventories []HostInventoryView `json:"inventories,omitempty"`
}

// GetHostCandidatesForVmMigrationView GetHostCandidatesForVmMigration
type GetHostCandidatesForVmMigrationView struct {
	Inventories []HostInventoryView `json:"inventories,omitempty"`
}

