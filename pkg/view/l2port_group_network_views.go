// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2PortGroupNetworkInventoryView L2PortGroupNetwork
type L2PortGroupNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	VSwitchUuid string `json:"vSwitchUuid,omitempty"`
	VlanMode string `json:"vlanMode,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	VlanRanges string `json:"vlanRanges,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
	AttachedHostRefs []L2NetworkHostRefInventoryView `json:"attachedHostRefs,omitempty"`
}

// QueryL2PortGroupNetworkView QueryL2PortGroupNetwork
type QueryL2PortGroupNetworkView struct {
	Inventories []L2PortGroupNetworkInventoryView `json:"inventories,omitempty"`
}

