// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmNicInventoryView VmNic
type VmNicInventoryView struct {
	BaseInfoView
	BaseTimeView
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

// ChangeVmNicTypeEventView ChangeVmNicTypeEvent
type ChangeVmNicTypeEventView struct {
	Inventory VmNicInventoryView `json:"inventory,omitempty"`
}

// GetEipAttachableVmNicsView GetEipAttachableVmNics
type GetEipAttachableVmNicsView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
	Start int `json:"start,omitempty"`
	More bool `json:"more,omitempty"`
	Success bool `json:"success,omitempty"`
}

// UpdateVmNicMacEventView UpdateVmNicMacEvent
type UpdateVmNicMacEventView struct {
	Inventory VmNicInventoryView `json:"inventory,omitempty"`
}

// CreateVmNicEventView CreateVmNicEvent
type CreateVmNicEventView struct {
	Inventory VmNicInventoryView `json:"inventory,omitempty"`
}

// GetPortForwardingAttachableVmNicsView GetPortForwardingAttachableVmNics
type GetPortForwardingAttachableVmNicsView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
}

// AttachL3NetworkToVmNicEventView AttachL3NetworkToVmNicEvent
type AttachL3NetworkToVmNicEventView struct {
	Inventory VmNicInventoryView `json:"inventory,omitempty"`
}

// QueryVmNicView QueryVmNic
type QueryVmNicView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
}

// UpdateVmNicDriverEventView UpdateVmNicDriverEvent
type UpdateVmNicDriverEventView struct {
	Inventory VmNicInventoryView `json:"inventory,omitempty"`
}

// ChangeVmNicNetworkEventView ChangeVmNicNetworkEvent
type ChangeVmNicNetworkEventView struct {
	Inventory VmNicInventoryView `json:"inventory,omitempty"`
}

// GetCandidateVmNicsForLoadBalancerServerGroupView GetCandidateVmNicsForLoadBalancerServerGroup
type GetCandidateVmNicsForLoadBalancerServerGroupView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
}

// GetCandidateVmNicForSecurityGroupView GetCandidateVmNicForSecurityGroup
type GetCandidateVmNicForSecurityGroupView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
}

// GetCandidateVmNicsForPortMirrorView GetCandidateVmNicsForPortMirror
type GetCandidateVmNicsForPortMirrorView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
}

// DeleteVmNicEventView DeleteVmNicEvent
type DeleteVmNicEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetCandidateVmNicsForLoadBalancerView GetCandidateVmNicsForLoadBalancer
type GetCandidateVmNicsForLoadBalancerView struct {
	Inventories []VmNicInventoryView `json:"inventories,omitempty"`
}

