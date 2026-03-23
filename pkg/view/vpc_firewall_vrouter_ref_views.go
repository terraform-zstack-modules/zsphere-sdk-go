// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcFirewallVRouterRefInventoryView VpcFirewallVRouterRef
type VpcFirewallVRouterRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	VpcFirewallUuid string `json:"vpcFirewallUuid,omitempty"`
	VRouterUuid string `json:"vRouterUuid,omitempty"`
}

// QueryVpcFirewallVRouterRefView QueryVpcFirewallVRouterRef
type QueryVpcFirewallVRouterRefView struct {
	Inventories []VpcFirewallVRouterRefInventoryView `json:"inventories,omitempty"`
}

