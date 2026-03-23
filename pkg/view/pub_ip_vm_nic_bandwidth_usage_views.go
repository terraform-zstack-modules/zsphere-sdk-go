// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PubIpVmNicBandwidthUsageInventoryView PubIpVmNicBandwidthUsage
type PubIpVmNicBandwidthUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	Inventory string `json:"inventory,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	BandwidthOut int64 `json:"bandwidthOut,omitempty"`
	BandwidthIn int64 `json:"bandwidthIn,omitempty"`
	VmNicIp string `json:"vmNicIp,omitempty"`
	VmNicStatus string `json:"vmNicStatus,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
}

