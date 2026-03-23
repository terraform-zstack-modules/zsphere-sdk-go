// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FaultToleranceVmGroupInventoryView FaultToleranceVmGroup
type FaultToleranceVmGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	PrimaryVmInstanceUuid string `json:"primaryVmInstanceUuid,omitempty"`
	SecondaryVmInstanceUuid string `json:"secondaryVmInstanceUuid,omitempty"`
	Status string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	ImageUuid string `json:"imageUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	LastHostUuid string `json:"lastHostUuid,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	RootVolumeUuid string `json:"rootVolumeUuid,omitempty"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	DefaultL3NetworkUuid string `json:"defaultL3NetworkUuid,omitempty"`
	Type string `json:"type,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	MemorySize int64 `json:"memorySize,omitempty"`
	ReservedMemorySize int64 `json:"reservedMemorySize,omitempty"`
	CpuNum int `json:"cpuNum,omitempty"`
	CpuSpeed int64 `json:"cpuSpeed,omitempty"`
	AllocatorStrategy string `json:"allocatorStrategy,omitempty"`
	State string `json:"state,omitempty"`
	VmNics []VmNicInventoryView `json:"vmNics,omitempty"`
	AllVolumes []VolumeInventoryView `json:"allVolumes,omitempty"`
	VmCdRoms []VmCdRomInventoryView `json:"vmCdRoms,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
}

// QueryFaultToleranceVmView QueryFaultToleranceVm
type QueryFaultToleranceVmView struct {
	Inventories []FaultToleranceVmGroupInventoryView `json:"inventories,omitempty"`
}

