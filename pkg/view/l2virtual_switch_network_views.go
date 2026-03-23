// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// L2VirtualSwitchNetworkInventoryView L2VirtualSwitchNetwork
type L2VirtualSwitchNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView
	IsDistributed bool `json:"isDistributed,omitempty"`
	VSwitchIndex int `json:"vSwitchIndex,omitempty"`
	PortGroups []PortGroupInventoryView `json:"portGroups,omitempty"`
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
	Type string `json:"type,omitempty"`
	VSwitchType string `json:"vSwitchType,omitempty"`
	VirtualNetworkId int `json:"virtualNetworkId,omitempty"`
	AttachedClusterUuids []string `json:"attachedClusterUuids,omitempty"`
	AttachedHostRefs []L2NetworkHostRefInventoryView `json:"attachedHostRefs,omitempty"`
}

// QueryL2VirtualSwitchNetworkView QueryL2VirtualSwitchNetwork
type QueryL2VirtualSwitchNetworkView struct {
	Inventories []L2VirtualSwitchNetworkInventoryView `json:"inventories,omitempty"`
}

