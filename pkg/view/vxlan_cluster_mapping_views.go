// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VxlanClusterMappingInventoryView VxlanClusterMapping
type VxlanClusterMappingInventoryView struct {
	BaseInfoView
	BaseTimeView
	VxlanUuid string `json:"vxlanUuid,omitempty"`
	ClusterUuid string `json:"clusterUuid,omitempty"`
	VlanId int `json:"vlanId,omitempty"`
	PhysicalInterface string `json:"physicalInterface,omitempty"`
}

