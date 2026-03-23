// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcHaGroupVipRefInventoryView VpcHaGroupVipRef
type VpcHaGroupVipRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VpcHaRouterUuid string `json:"vpcHaRouterUuid,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	Ip string `json:"ip,omitempty"`
	Netmask string `json:"netmask,omitempty"`
}

