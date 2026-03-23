// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceInventoryView PciDevice
type PciDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	PciSpecUuid string `json:"pciSpecUuid,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	VirtStatus string `json:"virtStatus,omitempty"`
	PassThroughState string `json:"passThroughState,omitempty"`
	Chooser string `json:"chooser,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	Device string `json:"device,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	PciDeviceAddress string `json:"pciDeviceAddress,omitempty"`
	IommuGroup string `json:"iommuGroup,omitempty"`
	MetaData PciDeviceMetaDataView `json:"metaData,omitempty"`
	MatchedPciDeviceOfferingRef []PciDevicePciDeviceOfferingRefInventoryView `json:"matchedPciDeviceOfferingRef,omitempty"`
	MdevSpecRefs []PciDeviceMdevSpecRefInventoryView `json:"mdevSpecRefs,omitempty"`
}

// DeletePciDeviceEventView DeletePciDeviceEvent
type DeletePciDeviceEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdatePciDeviceEventView UpdatePciDeviceEvent
type UpdatePciDeviceEventView struct {
	Inventory PciDeviceInventoryView `json:"inventory,omitempty"`
}

// DetachPciDeviceFromVmEventView DetachPciDeviceFromVmEvent
type DetachPciDeviceFromVmEventView struct {
	Inventory PciDeviceInventoryView `json:"inventory,omitempty"`
}

// GetPciDeviceCandidatesForNewCreateVmView GetPciDeviceCandidatesForNewCreateVm
type GetPciDeviceCandidatesForNewCreateVmView struct {
	Inventories []PciDeviceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AttachPciDeviceToVmEventView AttachPciDeviceToVmEvent
type AttachPciDeviceToVmEventView struct {
	Inventory PciDeviceInventoryView `json:"inventory,omitempty"`
}

// QueryPciDeviceView QueryPciDevice
type QueryPciDeviceView struct {
	Inventories []PciDeviceInventoryView `json:"inventories,omitempty"`
}

// GetPciDeviceCandidatesForAttachingVmView GetPciDeviceCandidatesForAttachingVm
type GetPciDeviceCandidatesForAttachingVmView struct {
	Inventories []PciDeviceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

