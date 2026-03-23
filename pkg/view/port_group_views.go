// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PortGroupInventoryView PortGroup
type PortGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	VSwitchUuid string `json:"vSwitchUuid,omitempty"`
	VlanMode string `json:"vlanMode,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	VlanRanges string `json:"vlanRanges,omitempty"`
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

// DeletePortGroupEventView DeletePortGroupEvent
type DeletePortGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreatePortGroupEventView CreatePortGroupEvent
type CreatePortGroupEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// UpdatePortGroupEventView UpdatePortGroupEvent
type UpdatePortGroupEventView struct {
	Inventory L3NetworkInventoryView `json:"inventory,omitempty"`
}

// QueryPortGroupView QueryPortGroup
type QueryPortGroupView struct {
	Inventories []PortGroupInventoryView `json:"inventories,omitempty"`
}

