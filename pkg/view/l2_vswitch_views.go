// Copyright (c) ZStack.io, Inc.

package view

type L2VirtualSwitchNetworkInventoryView struct {
	BaseInfoView
	BaseTimeView

	IsDistributed        bool                     `json:"isDistributed"`
	ZoneUuid             string                   `json:"zoneUuid"` // datacenter UUID. If specified, the cloud host will be created in the specified datacenter.
	PhysicalInterface    string                   `json:"physicalInterface"`
	Type                 string                   `json:"type"`
	VSwitchType          string                   `json:"vSwitchType"`
	VirtualNetworkId     int16                    `json:"virtualNetworkId"`
	AttachedClusterUuids []string                 `json:"attachedClusterUuids"`
	PortGroups           []PortGroupInventoryView `json:"portGroups"`
}
