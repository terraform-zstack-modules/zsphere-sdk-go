// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MirrorNetworkUsedIpInventoryView MirrorNetworkUsedIp
type MirrorNetworkUsedIpInventoryView struct {
	BaseInfoView
	BaseTimeView
	UsedIpInventory UsedIpInventoryView `json:"usedIpInventory,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	HostUuid string `json:"hostUuid,omitempty"`
	Description string `json:"description,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
}

// QueryPortMirrorNetworkUsedIpView QueryPortMirrorNetworkUsedIp
type QueryPortMirrorNetworkUsedIpView struct {
	Inventories []MirrorNetworkUsedIpInventoryView `json:"inventories,omitempty"`
}

