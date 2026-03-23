// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcHaGroupInventoryView VpcHaGroup
type VpcHaGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Monitors []VpcHaGroupMonitorIpInventoryView `json:"monitors,omitempty"`
	VrRefs []VpcHaGroupApplianceVmRefInventoryView `json:"vrRefs,omitempty"`
	Services []VpcHaGroupNetworkServiceRefInventoryView `json:"services,omitempty"`
	UsedIps []VpcHaGroupVipRefInventoryView `json:"usedIps,omitempty"`
}

// DeleteVpcHaGroupEventView DeleteVpcHaGroupEvent
type DeleteVpcHaGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateVpcHaGroupEventView CreateVpcHaGroupEvent
type CreateVpcHaGroupEventView struct {
	Inventory VpcHaGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateVpcHaGroupEventView UpdateVpcHaGroupEvent
type UpdateVpcHaGroupEventView struct {
	Inventory VpcHaGroupInventoryView `json:"inventory,omitempty"`
}

// ChangeVpcHaGroupMonitorIpsEventView ChangeVpcHaGroupMonitorIpsEvent
type ChangeVpcHaGroupMonitorIpsEventView struct {
	Inventory VpcHaGroupInventoryView `json:"inventory,omitempty"`
}

// QueryVpcHaGroupView QueryVpcHaGroup
type QueryVpcHaGroupView struct {
	Inventories []VpcHaGroupInventoryView `json:"inventories,omitempty"`
}

