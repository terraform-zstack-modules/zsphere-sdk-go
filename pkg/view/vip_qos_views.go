// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VipQosInventoryView VipQos
type VipQosInventoryView struct {
	BaseInfoView
	BaseTimeView
	VipUuid string `json:"vipUuid,omitempty"`
	Port int `json:"port,omitempty"`
	InboundBandwidth int64 `json:"inboundBandwidth,omitempty"`
	OutboundBandwidth int64 `json:"outboundBandwidth,omitempty"`
	Type string `json:"type,omitempty"`
}

// DeleteVipQosEventView DeleteVipQosEvent
type DeleteVipQosEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVipQosView GetVipQos
type GetVipQosView struct {
	Inventories []VipQosInventoryView `json:"inventories,omitempty"`
}

// SetVipQosEventView SetVipQosEvent
type SetVipQosEventView struct {
	Inventory VipQosInventoryView `json:"inventory,omitempty"`
}

