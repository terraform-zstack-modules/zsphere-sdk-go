// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// FreeIpInventoryView FreeIp
type FreeIpInventoryView struct {
	BaseInfoView
	BaseTimeView
	IpRangeUuid string `json:"ipRangeUuid,omitempty"`
	Ip string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

