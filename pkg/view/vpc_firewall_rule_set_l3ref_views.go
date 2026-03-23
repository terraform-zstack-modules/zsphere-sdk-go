// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcFirewallRuleSetL3RefInventoryView VpcFirewallRuleSetL3Ref
type VpcFirewallRuleSetL3RefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	VpcFirewallUuid string `json:"vpcFirewallUuid,omitempty"`
	PacketsForwardType string `json:"packetsForwardType,omitempty"`
}

// AttachFirewallRuleSetToL3EventView AttachFirewallRuleSetToL3Event
type AttachFirewallRuleSetToL3EventView struct {
	Inventory VpcFirewallRuleSetL3RefInventoryView `json:"inventory,omitempty"`
}

// QueryFirewallRuleSetL3RefView QueryFirewallRuleSetL3Ref
type QueryFirewallRuleSetL3RefView struct {
	Inventories []VpcFirewallRuleSetL3RefInventoryView `json:"inventories,omitempty"`
}

