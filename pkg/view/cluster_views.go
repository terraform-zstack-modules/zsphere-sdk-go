// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ClusterInventoryView Cluster
type ClusterInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	HypervisorType string `json:"hypervisorType,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

// ChangeClusterStateEventView ChangeClusterStateEvent
type ChangeClusterStateEventView struct {
	Inventory ClusterInventoryView `json:"inventory,omitempty"`
}

// GetCandidateClustersForAttachingL2NetworkView GetCandidateClustersForAttachingL2Network
type GetCandidateClustersForAttachingL2NetworkView struct {
	Inventories []ClusterInventoryView `json:"inventories,omitempty"`
	Success bool `json:"success,omitempty"`
}

// CreateMiniClusterEventView CreateMiniClusterEvent
type CreateMiniClusterEventView struct {
	Inventory ClusterInventoryView `json:"inventory,omitempty"`
}

// DeleteClusterEventView DeleteClusterEvent
type DeleteClusterEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateClusterEventView UpdateClusterEvent
type UpdateClusterEventView struct {
	Inventory ClusterInventoryView `json:"inventory,omitempty"`
}

// QueryClusterView QueryCluster
type QueryClusterView struct {
	Inventories []ClusterInventoryView `json:"inventories,omitempty"`
}

// CreateClusterEventView CreateClusterEvent
type CreateClusterEventView struct {
	Inventory ClusterInventoryView `json:"inventory,omitempty"`
}

