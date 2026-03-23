// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ClusterDRSInventoryView ClusterDRS
type ClusterDRSInventoryView struct {
	BaseInfoView
	BaseTimeView
	ClusterUuid string `json:"clusterUuid,omitempty"`
	State string `json:"state,omitempty"`
	BalancedState string `json:"balancedState,omitempty"`
	LastAdviceGroupUuid string `json:"lastAdviceGroupUuid,omitempty"`
	AutomationLevel string `json:"automationLevel,omitempty"`
	Thresholds []ThresholdView `json:"thresholds,omitempty"`
	ThresholdDuration int `json:"thresholdDuration,omitempty"`
	Description string `json:"description,omitempty"`
}

// CreateClusterDRSEventView CreateClusterDRSEvent
type CreateClusterDRSEventView struct {
	Inventory ClusterDRSInventoryView `json:"inventory,omitempty"`
}

// UpdateClusterDRSEventView UpdateClusterDRSEvent
type UpdateClusterDRSEventView struct {
	Inventory ClusterDRSInventoryView `json:"inventory,omitempty"`
}

// QueryClusterDRSView QueryClusterDRS
type QueryClusterDRSView struct {
	Inventories []ClusterDRSInventoryView `json:"inventories,omitempty"`
}

// DeleteClusterDRSEventView DeleteClusterDRSEvent
type DeleteClusterDRSEventView struct {
	Success bool `json:"success,omitempty"`
}

