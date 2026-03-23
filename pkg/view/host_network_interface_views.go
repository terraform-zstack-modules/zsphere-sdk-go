// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostNetworkInterfaceInventoryView HostNetworkInterface
type HostNetworkInterfaceInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostUuid string `json:"hostUuid,omitempty"`
	BondingUuid string `json:"bondingUuid,omitempty"`
	InterfaceModel string `json:"interfaceModel,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	VendorName string `json:"vendorName,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	SubvendorName string `json:"subvendorName,omitempty"`
	InterfaceName string `json:"interfaceName,omitempty"`
	InterfaceType string `json:"interfaceType,omitempty"`
	Speed int64 `json:"speed,omitempty"`
	SlaveActive bool `json:"slaveActive,omitempty"`
	CarrierActive bool `json:"carrierActive,omitempty"`
	IpAddresses []string `json:"ipAddresses,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Mac string `json:"mac,omitempty"`
	CallBackIp string `json:"callBackIp,omitempty"`
	PciDeviceAddress string `json:"pciDeviceAddress,omitempty"`
	OffloadStatus string `json:"offloadStatus,omitempty"`
	VirtStatus string `json:"virtStatus,omitempty"`
	Description string `json:"description,omitempty"`
}

// SetIpOnHostNetworkInterfaceEventView SetIpOnHostNetworkInterfaceEvent
type SetIpOnHostNetworkInterfaceEventView struct {
	Inventory HostNetworkInterfaceInventoryView `json:"inventory,omitempty"`
}

// UpdateHostNetworkInterfaceEventView UpdateHostNetworkInterfaceEvent
type UpdateHostNetworkInterfaceEventView struct {
	Inventory HostNetworkInterfaceInventoryView `json:"inventory,omitempty"`
}

// QueryHostNetworkInterfaceView QueryHostNetworkInterface
type QueryHostNetworkInterfaceView struct {
	Inventories []HostNetworkInterfaceInventoryView `json:"inventories,omitempty"`
}

// LocateHostNetworkInterfaceEventView LocateHostNetworkInterfaceEvent
type LocateHostNetworkInterfaceEventView struct {
	Success bool `json:"success,omitempty"`
}

