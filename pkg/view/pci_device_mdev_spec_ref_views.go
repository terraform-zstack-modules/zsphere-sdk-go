// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PciDeviceMdevSpecRefInventoryView PciDeviceMdevSpecRef
type PciDeviceMdevSpecRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	PciDeviceUuid string `json:"pciDeviceUuid,omitempty"`
	MdevSpecUuid string `json:"mdevSpecUuid,omitempty"`
	Effective bool `json:"effective,omitempty"`
}

