// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceUsageInventoryView PciDeviceUsage
type PciDeviceUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	PciDeviceUuid string `json:"pciDeviceUuid,omitempty"`
	VendorId string `json:"vendorId,omitempty"`
	DeviceId string `json:"deviceId,omitempty"`
	SubvendorId string `json:"subvendorId,omitempty"`
	SubdeviceId string `json:"subdeviceId,omitempty"`
	Description string `json:"description,omitempty"`
	VmUuid string `json:"vmUuid,omitempty"`
	VmName string `json:"vmName,omitempty"`
	Status string `json:"status,omitempty"`
	Inventory string `json:"inventory,omitempty"`
}

