// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PubIpVipBandwidthUsageInventoryView PubIpVipBandwidthUsage
type PubIpVipBandwidthUsageInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
	DateInLong int64 `json:"dateInLong,omitempty"`
	Inventory string `json:"inventory,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	VipName string `json:"vipName,omitempty"`
	VipIp string `json:"vipIp,omitempty"`
	BandwidthIn int64 `json:"bandwidthIn,omitempty"`
	BandwidthOut int64 `json:"bandwidthOut,omitempty"`
	VipStatus string `json:"vipStatus,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
}

