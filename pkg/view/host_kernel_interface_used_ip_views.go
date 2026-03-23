// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostKernelInterfaceUsedIpInventoryView HostKernelInterfaceUsedIp
type HostKernelInterfaceUsedIpInventoryView struct {
	BaseInfoView
	BaseTimeView
	HostKernelInterfaceUuid string `json:"hostKernelInterfaceUuid,omitempty"`
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

