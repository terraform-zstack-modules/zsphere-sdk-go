// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ManagementNodeStatusViewView ManagementNodeStatusView
type ManagementNodeStatusViewView struct {
	Ip string `json:"ip,omitempty"`
	GatewayIp string `json:"gatewayIp,omitempty"`
	OwnsVip bool `json:"ownsVip,omitempty"`
	PeerReachable bool `json:"peerReachable,omitempty"`
	GatewayReachable bool `json:"gatewayReachable,omitempty"`
	VipReachable bool `json:"vipReachable,omitempty"`
	KeepalivedStatus string `json:"keepalivedStatus,omitempty"`
	HaMonitorStatus string `json:"haMonitorStatus,omitempty"`
	DatabaseStatus string `json:"databaseStatus,omitempty"`
	UiStatus string `json:"uiStatus,omitempty"`
	ManagementsNodeStatus string `json:"managementsNodeStatus,omitempty"`
	SlaveIoRunning bool `json:"slaveIoRunning,omitempty"`
	SlaveSqlRunning bool `json:"slaveSqlRunning,omitempty"`
}

