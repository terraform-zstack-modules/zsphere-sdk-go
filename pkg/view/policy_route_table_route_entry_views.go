// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PolicyRouteTableRouteEntryInventoryView PolicyRouteTableRouteEntry
type PolicyRouteTableRouteEntryInventoryView struct {
	BaseInfoView
	BaseTimeView
	TableUuid string `json:"tableUuid,omitempty"`
	DestinationCidr string `json:"destinationCidr,omitempty"`
	NextHopIp string `json:"nextHopIp,omitempty"`
	Distance int `json:"distance,omitempty"`
}

// DeletePolicyRouteTableRouteEntryEventView DeletePolicyRouteTableRouteEntryEvent
type DeletePolicyRouteTableRouteEntryEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryPolicyRouteTableRouteEntryView QueryPolicyRouteTableRouteEntry
type QueryPolicyRouteTableRouteEntryView struct {
	Inventories []PolicyRouteTableRouteEntryInventoryView `json:"inventories,omitempty"`
}

// CreatePolicyRouteTableRouteEntryEventView CreatePolicyRouteTableRouteEntryEvent
type CreatePolicyRouteTableRouteEntryEventView struct {
	Inventory PolicyRouteTableRouteEntryInventoryView `json:"inventory,omitempty"`
}

