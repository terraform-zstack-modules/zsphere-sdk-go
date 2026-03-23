// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostIpmiInventoryView HostIpmi
type HostIpmiInventoryView struct {
	BaseInfoView
	BaseTimeView
	IpmiAddress string `json:"ipmiAddress,omitempty"`
	IpmiUsername string `json:"ipmiUsername,omitempty"`
	IpmiPort int `json:"ipmiPort,omitempty"`
	IpmiPowerStatus string `json:"ipmiPowerStatus,omitempty"`
}

// GetHostPowerStatusEventView GetHostPowerStatusEvent
type GetHostPowerStatusEventView struct {
	Inventory HostIpmiInventoryView `json:"inventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UpdateHostIpmiEventView UpdateHostIpmiEvent
type UpdateHostIpmiEventView struct {
	HostIpmiInventory HostIpmiInventoryView `json:"hostIpmiInventory,omitempty"`
	Success bool `json:"success,omitempty"`
}

