// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2VlanNetworkInventoryView L2VlanNetwork
type L2VlanNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	Vlan int `json:"vlan,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
	AttachedHostRefs []L2NetworkHostRefInventoryView `json:"attachedHostRefs,omitempty"`
}

// QueryL2VlanNetworkView QueryL2VlanNetwork
type QueryL2VlanNetworkView struct {
	Inventories []L2VlanNetworkInventoryView `json:"inventories,omitempty"`
}

// CreateL2VlanNetworkEventView CreateL2VlanNetworkEvent
type CreateL2VlanNetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

