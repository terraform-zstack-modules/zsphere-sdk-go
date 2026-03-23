// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmVfNicInventoryView VmVfNic
type VmVfNicInventoryView struct {
	BaseInfoView
	BaseTimeView
	PciDeviceUuid string `json:"pciDeviceUuid,omitempty"`
	HaState string `json:"haState,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Ip string `json:"ip,omitempty"`
	Mac string `json:"mac,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	MetaData string `json:"metaData,omitempty"`
	DriverType string `json:"driverType,omitempty"`
	UsedIps []UsedIpInventoryView `json:"usedIps,omitempty"`
	InternalName string `json:"internalName,omitempty"`
	DeviceId int `json:"deviceId,omitempty"`
	Type string `json:"type,omitempty"`
	State string `json:"state,omitempty"`
}

// ChangeVfNicHaStateEventView ChangeVfNicHaStateEvent
type ChangeVfNicHaStateEventView struct {
	Inventory VmVfNicInventoryView `json:"inventory,omitempty"`
}

