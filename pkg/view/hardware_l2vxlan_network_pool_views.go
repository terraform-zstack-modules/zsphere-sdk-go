// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HardwareL2VxlanNetworkPoolInventoryView HardwareL2VxlanNetworkPool
type HardwareL2VxlanNetworkPoolInventoryView struct {
	BaseInfoView
	BaseTimeView
	SdnControllerUuid string `json:"sdnControllerUuid,omitempty"`
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

