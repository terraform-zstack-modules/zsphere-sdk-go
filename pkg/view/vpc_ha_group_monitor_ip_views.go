// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcHaGroupMonitorIpInventoryView VpcHaGroupMonitorIp
type VpcHaGroupMonitorIpInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VpcHaRouterUuid string `json:"vpcHaRouterUuid,omitempty"`
	MonitorIp string `json:"monitorIp,omitempty"`
}

