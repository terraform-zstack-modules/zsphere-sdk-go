// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2NetworkInventoryView L2Network
type L2NetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
	AttachedHostRefs []L2NetworkHostRefInventoryView `json:"attachedHostRefs,omitempty"`
}

// AttachL2NetworkToClusterEventView AttachL2NetworkToClusterEvent
type AttachL2NetworkToClusterEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// DeleteL2NetworkEventView DeleteL2NetworkEvent
type DeleteL2NetworkEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateL2NetworkVirtualNetworkIdEventView UpdateL2NetworkVirtualNetworkIdEvent
type UpdateL2NetworkVirtualNetworkIdEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// AttachL2NetworkToHostEventView AttachL2NetworkToHostEvent
type AttachL2NetworkToHostEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// DetachL2NetworkFromHostEventView DetachL2NetworkFromHostEvent
type DetachL2NetworkFromHostEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// UpdateL2NetworkEventView UpdateL2NetworkEvent
type UpdateL2NetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// CreateL2NetworkEventView CreateL2NetworkEvent
type CreateL2NetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// DetachL2NetworkFromClusterEventView DetachL2NetworkFromClusterEvent
type DetachL2NetworkFromClusterEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// QueryL2NetworkView QueryL2Network
type QueryL2NetworkView struct {
	Inventories []L2NetworkInventoryView `json:"inventories,omitempty"`
}

