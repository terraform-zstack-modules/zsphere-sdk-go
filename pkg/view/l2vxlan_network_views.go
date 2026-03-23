// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2VxlanNetworkInventoryView L2VxlanNetwork
type L2VxlanNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	Vni int `json:"vni,omitempty"`
	PoolUuid string `json:"poolUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
	AttachedHostRefs []L2NetworkHostRefInventoryView `json:"attachedHostRefs,omitempty"`
}

// CreateL2VxlanNetworkEventView CreateL2VxlanNetworkEvent
type CreateL2VxlanNetworkEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// QueryL2VxlanNetworkView QueryL2VxlanNetwork
type QueryL2VxlanNetworkView struct {
	Inventories []L2VxlanNetworkInventoryView `json:"inventories,omitempty"`
}

