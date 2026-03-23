// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcFirewallInventoryView VpcFirewall
type VpcFirewallInventoryView struct {
	BaseInfoView
	BaseTimeView
	Refs []VpcFirewallRuleSetL3RefInventoryView `json:"refs,omitempty"`
	Description string `json:"description,omitempty"`
}

// UpdateVpcFirewallEventView UpdateVpcFirewallEvent
type UpdateVpcFirewallEventView struct {
	Inventory VpcFirewallInventoryView `json:"inventory,omitempty"`
}

// CreateVpcFirewallEventView CreateVpcFirewallEvent
type CreateVpcFirewallEventView struct {
	Inventory VpcFirewallInventoryView `json:"inventory,omitempty"`
}

// RefreshFirewallEventView RefreshFirewallEvent
type RefreshFirewallEventView struct {
	Inventory VpcFirewallInventoryView `json:"inventory,omitempty"`
}

// QueryVpcFirewallView QueryVpcFirewall
type QueryVpcFirewallView struct {
	Inventories []VpcFirewallInventoryView `json:"inventories,omitempty"`
}

