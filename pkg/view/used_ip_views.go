// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// UsedIpInventoryView UsedIp
type UsedIpInventoryView struct {
	BaseInfoView
	BaseTimeView
	IpRangeUuid string `json:"ipRangeUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	Ip string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	UsedFor string `json:"usedFor,omitempty"`
	IpInLong int64 `json:"ipInLong,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
}

// QueryIpAddressView QueryIpAddress
type QueryIpAddressView struct {
	Inventories []UsedIpInventoryView `json:"inventories,omitempty"`
}

