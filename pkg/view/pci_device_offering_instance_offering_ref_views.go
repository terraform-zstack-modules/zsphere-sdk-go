// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceOfferingInstanceOfferingRefInventoryView PciDeviceOfferingInstanceOfferingRef
type PciDeviceOfferingInstanceOfferingRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	InstanceOfferingUuid string `json:"instanceOfferingUuid,omitempty"`
	PciDeviceOfferingUuid string `json:"pciDeviceOfferingUuid,omitempty"`
	Metadata PciDeviceMetaDataView `json:"metadata,omitempty"`
	PciDeviceCount int `json:"pciDeviceCount,omitempty"`
}

