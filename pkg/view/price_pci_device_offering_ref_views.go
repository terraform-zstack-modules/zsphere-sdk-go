// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PricePciDeviceOfferingRefInventoryView PricePciDeviceOfferingRef
type PricePciDeviceOfferingRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	PriceUuid string `json:"priceUuid,omitempty"`
	PciDeviceOfferingUuid string `json:"pciDeviceOfferingUuid,omitempty"`
}

