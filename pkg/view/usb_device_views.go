// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// UsbDeviceInventoryView UsbDevice
type UsbDeviceInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	State string `json:"state,omitempty"`
	BusNum string `json:"busNum,omitempty"`
	DevNum string `json:"devNum,omitempty"`
	IdVendor string `json:"idVendor,omitempty"`
	IdProduct string `json:"idProduct,omitempty"`
	IManufacturer string `json:"iManufacturer,omitempty"`
	IProduct string `json:"iProduct,omitempty"`
	ISerial string `json:"iSerial,omitempty"`
	UsbVersion string `json:"usbVersion,omitempty"`
	AttachType string `json:"attachType,omitempty"`
}

// UpdateUsbDeviceEventView UpdateUsbDeviceEvent
type UpdateUsbDeviceEventView struct {
	Inventory UsbDeviceInventoryView `json:"inventory,omitempty"`
}

// QueryUsbDeviceView QueryUsbDevice
type QueryUsbDeviceView struct {
	Inventories []UsbDeviceInventoryView `json:"inventories,omitempty"`
}

// GetUsbDeviceCandidatesForAttachingVmView GetUsbDeviceCandidatesForAttachingVm
type GetUsbDeviceCandidatesForAttachingVmView struct {
	Inventories []UsbDeviceInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AttachUsbDeviceToVmEventView AttachUsbDeviceToVmEvent
type AttachUsbDeviceToVmEventView struct {
	Inventory UsbDeviceInventoryView `json:"inventory,omitempty"`
}

// DetachUsbDeviceFromVmEventView DetachUsbDeviceFromVmEvent
type DetachUsbDeviceFromVmEventView struct {
	Inventory UsbDeviceInventoryView `json:"inventory,omitempty"`
}

