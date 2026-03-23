// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ZStoneClusterViewView ZStoneClusterView
type ZStoneClusterViewView struct {
	Uuid string `json:"uuid,omitempty"`
	First bool `json:"first,omitempty"`
	ManagementNetworkCidr string `json:"managementNetworkCidr,omitempty"`
	PublicNetworkCidr string `json:"publicNetworkCidr,omitempty"`
	ClusterNetworkCidr string `json:"clusterNetworkCidr,omitempty"`
	ChronyIp string `json:"chronyIp,omitempty"`
	Type string `json:"type,omitempty"`
	Hosts ZStoneHostSummaryViewView `json:"hosts,omitempty"`
	Pools ZStonePoolSummaryViewView `json:"pools,omitempty"`
}

