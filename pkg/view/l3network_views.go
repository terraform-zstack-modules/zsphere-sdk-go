// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L3NetworkInventoryView L3Network
type L3NetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Type string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	L2NetworkUuid string `json:"l2NetworkUuid,omitempty"`
	State string `json:"state,omitempty"`
	DnsDomain string `json:"dnsDomain,omitempty"`
	System bool `json:"system,omitempty"`
	Category string `json:"category,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	EnableIPAM bool `json:"enableIPAM,omitempty"`
	Dns []string `json:"dns,omitempty"`
	IpRanges []IpRangeInventoryView `json:"ipRanges,omitempty"`
	NetworkServices []NetworkServiceL3NetworkRefInventoryView `json:"networkServices,omitempty"`
	HostRoute []L3NetworkHostRouteInventoryView `json:"hostRoute,omitempty"`
	ReservedIpRanges []ReservedIpRangeInventoryView `json:"reservedIpRanges,omitempty"`
}

// GetCandidateL3NetworksForChangeVmNicNetworkView GetCandidateL3NetworksForChangeVmNicNetwork
type GetCandidateL3NetworksForChangeVmNicNetworkView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// AttachNetworkServiceToL3NetworkEventView AttachNetworkServiceToL3NetworkEvent
type AttachNetworkServiceToL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// AddHostRouteToL3NetworkEventView AddHostRouteToL3NetworkEvent
type AddHostRouteToL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// GetAttachablePublicL3ForVRouterView GetAttachablePublicL3ForVRouter
type GetAttachablePublicL3ForVRouterView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// RemoveDnsFromL3NetworkEventView RemoveDnsFromL3NetworkEvent
type RemoveDnsFromL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// ChangeL3NetworkStateEventView ChangeL3NetworkStateEvent
type ChangeL3NetworkStateEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// CreateL3NetworkEventView CreateL3NetworkEvent
type CreateL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// DeleteL3NetworkEventView DeleteL3NetworkEvent
type DeleteL3NetworkEventView struct {
	Success bool `json:"success,omitempty"`
}

// RemoveHostRouteFromL3NetworkEventView RemoveHostRouteFromL3NetworkEvent
type RemoveHostRouteFromL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// QueryL3NetworkView QueryL3Network
type QueryL3NetworkView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// GetAttachableVpcL3NetworkView GetAttachableVpcL3Network
type GetAttachableVpcL3NetworkView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// AddDnsToL3NetworkEventView AddDnsToL3NetworkEvent
type AddDnsToL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// GetCandidateL3NetworksForIpSecConnectionView GetCandidateL3NetworksForIpSecConnection
type GetCandidateL3NetworksForIpSecConnectionView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// GetVmAttachableL3NetworkView GetVmAttachableL3Network
type GetVmAttachableL3NetworkView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// GetCandidateL3NetworksForLoadBalancerView GetCandidateL3NetworksForLoadBalancer
type GetCandidateL3NetworksForLoadBalancerView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// GetCandidateL3NetworksForServerGroupView GetCandidateL3NetworksForServerGroup
type GetCandidateL3NetworksForServerGroupView struct {
	Inventories []L3NetworkInventoryView `json:"inventories,omitempty"`
}

// DetachNetworkServiceFromL3NetworkEventView DetachNetworkServiceFromL3NetworkEvent
type DetachNetworkServiceFromL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// UpdateL3NetworkEventView UpdateL3NetworkEvent
type UpdateL3NetworkEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

