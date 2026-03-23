// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmInstanceInventoryView VmInstance
type VmInstanceInventoryView struct {
	BaseInfoView
	BaseTimeView
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

// CloneVmInstanceEventView CloneVmInstanceEvent
type CloneVmInstanceEventView struct {
	Result CloneVmInstanceResultsView `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CreateVmFromVmBackupEventView CreateVmFromVmBackupEvent
type CreateVmFromVmBackupEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmClockTrackEventView SetVmClockTrackEvent
type SetVmClockTrackEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// StopVmInstanceEventView StopVmInstanceEvent
type StopVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// AttachL3NetworkToVmEventView AttachL3NetworkToVmEvent
type AttachL3NetworkToVmEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// ResumeVmInstanceEventView ResumeVmInstanceEvent
type ResumeVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmBootOrderEventView SetVmBootOrderEvent
type SetVmBootOrderEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// QueryVmInstanceView QueryVmInstance
type QueryVmInstanceView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
}

// DetachIsoFromVmInstanceEventView DetachIsoFromVmInstanceEvent
type DetachIsoFromVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// ExpungeVmInstanceEventView ExpungeVmInstanceEvent
type ExpungeVmInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangeInstanceOfferingEventView ChangeInstanceOfferingEvent
type ChangeInstanceOfferingEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// RebootVmInstanceEventView RebootVmInstanceEvent
type RebootVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// DetachL3NetworkFromVmEventView DetachL3NetworkFromVmEvent
type DetachL3NetworkFromVmEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// UpdateVmInstanceEventView UpdateVmInstanceEvent
type UpdateVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// DeleteVmSshKeyEventView DeleteVmSshKeyEvent
type DeleteVmSshKeyEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// MigrateVmEventView MigrateVmEvent
type MigrateVmEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// DeleteVmConsolePasswordEventView DeleteVmConsolePasswordEvent
type DeleteVmConsolePasswordEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// DestroyVmInstanceEventView DestroyVmInstanceEvent
type DestroyVmInstanceEventView struct {
	Success bool `json:"success,omitempty"`
}

// StartVmInstanceEventView StartVmInstanceEvent
type StartVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// DeleteVmStaticIpEventView DeleteVmStaticIpEvent
type DeleteVmStaticIpEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotEventView CreateVmInstanceFromVolumeSnapshotEvent
type CreateVmInstanceFromVolumeSnapshotEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmSshKeyEventView SetVmSshKeyEvent
type SetVmSshKeyEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// AttachIsoToVmInstanceEventView AttachIsoToVmInstanceEvent
type AttachIsoToVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// AttachVmNicToVmEventView AttachVmNicToVmEvent
type AttachVmNicToVmEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// CreateVmFromVolumeBackupEventView CreateVmFromVolumeBackupEvent
type CreateVmFromVolumeBackupEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// ChangeVmNicStateEventView ChangeVmNicStateEvent
type ChangeVmNicStateEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// DeleteVmCdRomEventView DeleteVmCdRomEvent
type DeleteVmCdRomEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// SetVmConsolePasswordEventView SetVmConsolePasswordEvent
type SetVmConsolePasswordEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// CreateVpcVRouterEventView CreateVpcVRouterEvent
type CreateVpcVRouterEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// CreateVmInstanceFromVolumeSnapshotGroupEventView CreateVmInstanceFromVolumeSnapshotGroupEvent
type CreateVmInstanceFromVolumeSnapshotGroupEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// CreateVmInstanceEventView CreateVmInstanceEvent
type CreateVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmBootVolumeEventView SetVmBootVolumeEvent
type SetVmBootVolumeEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// SetVmConsoleModeEventView SetVmConsoleModeEvent
type SetVmConsoleModeEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// GetDataVolumeAttachableVmView GetDataVolumeAttachableVm
type GetDataVolumeAttachableVmView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
}

// CreateVmInstanceFromOvfEventView CreateVmInstanceFromOvfEvent
type CreateVmInstanceFromOvfEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// ChangeVmImageEventView ChangeVmImageEvent
type ChangeVmImageEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// CreateVmInstanceFromVolumeEventView CreateVmInstanceFromVolumeEvent
type CreateVmInstanceFromVolumeEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// CreateVmFromCdpBackupEventView CreateVmFromCdpBackupEvent
type CreateVmFromCdpBackupEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

// GetCandidateVMForAttachingAffinityGroupView GetCandidateVMForAttachingAffinityGroup
type GetCandidateVMForAttachingAffinityGroupView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// GetCandidateVmForAttachingIsoView GetCandidateVmForAttachingIso
type GetCandidateVmForAttachingIsoView struct {
	Inventories []VmInstanceInventoryView `json:"inventories,omitempty"`
}

// RecoverVmInstanceEventView RecoverVmInstanceEvent
type RecoverVmInstanceEventView struct {
	Inventory VmInstanceInventoryView `json:"inventory,omitempty"`
}

