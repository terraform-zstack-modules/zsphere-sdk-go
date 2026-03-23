// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// NetworkRouterAreaRefInventoryView NetworkRouterAreaRef
type NetworkRouterAreaRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	VRouterUuid string `json:"vRouterUuid,omitempty"`
	ApplianceVmType string `json:"applianceVmType,omitempty"`
	RouterAreaUuid string `json:"routerAreaUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
}

// AddVRouterNetworksToOspfAreaEventView AddVRouterNetworksToOspfAreaEvent
type AddVRouterNetworksToOspfAreaEventView struct {
	Inventories []NetworkRouterAreaRefInventoryView `json:"inventories,omitempty"`
}

// GetVpcAttachedOspfView GetVpcAttachedOspf
type GetVpcAttachedOspfView struct {
	Inventories []NetworkRouterAreaRefInventoryView `json:"inventories,omitempty"`
}

// QueryVRouterOspfNetworkView QueryVRouterOspfNetwork
type QueryVRouterOspfNetworkView struct {
	Inventories []NetworkRouterAreaRefInventoryView `json:"inventories,omitempty"`
}

