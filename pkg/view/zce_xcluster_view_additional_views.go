// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZceXClusterViewView ZceXClusterView
type ZceXClusterViewView struct {
	ManagementNetworkCidr string `json:"managementNetworkCidr,omitempty"`
	GatewayNetworkCidr string `json:"gatewayNetworkCidr,omitempty"`
	PublicNetworkCidr string `json:"publicNetworkCidr,omitempty"`
	ClusterNetworkCidr string `json:"clusterNetworkCidr,omitempty"`
	Hosts ZceXHostSummaryViewView `json:"hosts,omitempty"`
	Pools ZceXPoolSummaryViewView `json:"pools,omitempty"`
}

