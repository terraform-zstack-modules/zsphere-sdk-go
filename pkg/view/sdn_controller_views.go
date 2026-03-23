// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SdnControllerInventoryView SdnController
type SdnControllerInventoryView struct {
	BaseInfoView
	BaseTimeView
	VendorType string `json:"vendorType,omitempty"`
	Description string `json:"description,omitempty"`
	Ip string `json:"ip,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	VniRanges []SdnVniRangeView `json:"vniRanges,omitempty"`
	VxlanPools []HardwareL2VxlanNetworkPoolInventoryView `json:"vxlanPools,omitempty"`
}

// RemoveSdnControllerEventView RemoveSdnControllerEvent
type RemoveSdnControllerEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddSdnControllerEventView AddSdnControllerEvent
type AddSdnControllerEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// UpdateSdnControllerEventView UpdateSdnControllerEvent
type UpdateSdnControllerEventView struct {
	Inventory SdnControllerInventoryView `json:"inventory,omitempty"`
}

// QuerySdnControllerView QuerySdnController
type QuerySdnControllerView struct {
	Inventories []SdnControllerInventoryView `json:"inventories,omitempty"`
}

