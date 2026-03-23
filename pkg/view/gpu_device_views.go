// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// GpuDeviceInventoryView GpuDevice
type GpuDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	SerialNumber string `json:"serialNumber,omitempty"`
	Memory int64 `json:"memory,omitempty"`
	Power int64 `json:"power,omitempty"`
	IsDriverLoaded bool `json:"isDriverLoaded,omitempty"`
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

// QueryGpuDeviceView QueryGpuDevice
type QueryGpuDeviceView struct {
	Inventories []GpuDeviceInventoryView `json:"inventories,omitempty"`
}

