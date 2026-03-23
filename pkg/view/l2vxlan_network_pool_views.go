// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2VxlanNetworkPoolInventoryView L2VxlanNetworkPool
type L2VxlanNetworkPoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	AttachedVtepRefs []VtepInventoryView `json:"attachedVtepRefs,omitempty"`
	RemoteVteps []RemoteVtepInventoryView `json:"remoteVteps,omitempty"`
	AttachedVxlanNetworkRefs []L2VxlanNetworkInventoryView `json:"attachedVxlanNetworkRefs,omitempty"`
	AttachedVniRanges []VniRangeInventoryView `json:"attachedVniRanges,omitempty"`
	AttachedCidrs map[string]string `json:"attachedCidrs,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
	AttachedHostRefs []L2NetworkHostRefInventoryView `json:"attachedHostRefs,omitempty"`
}

// CreateL2VxlanNetworkPoolEventView CreateL2VxlanNetworkPoolEvent
type CreateL2VxlanNetworkPoolEventView struct {
	Inventory L2NetworkInventoryView `json:"inventory,omitempty"`
}

// QueryL2VxlanNetworkPoolView QueryL2VxlanNetworkPool
type QueryL2VxlanNetworkPoolView struct {
	Inventories []L2VxlanNetworkPoolInventoryView `json:"inventories,omitempty"`
}

