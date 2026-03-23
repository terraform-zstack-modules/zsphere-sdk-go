// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcFirewallRuleInventoryView VpcFirewallRule
type VpcFirewallRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	RuleSetUuid string `json:"ruleSetUuid,omitempty"`
	Action string `json:"action,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	DestPort string `json:"destPort,omitempty"`
	SourcePort string `json:"sourcePort,omitempty"`
	SourceIp string `json:"sourceIp,omitempty"`
	DestIp string `json:"destIp,omitempty"`
	RuleNumber int `json:"ruleNumber,omitempty"`
	AllowStates string `json:"allowStates,omitempty"`
	TcpFlag string `json:"tcpFlag,omitempty"`
	IcmpTypeName string `json:"icmpTypeName,omitempty"`
	IsApplied bool `json:"isApplied,omitempty"`
	Expired bool `json:"expired,omitempty"`
	State string `json:"state,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Description string `json:"description,omitempty"`
}

// ChangeFirewallRuleStateEventView ChangeFirewallRuleStateEvent
type ChangeFirewallRuleStateEventView struct {
	Inventory VpcFirewallRuleInventoryView `json:"inventory,omitempty"`
}

// CreateFirewallRuleEventView CreateFirewallRuleEvent
type CreateFirewallRuleEventView struct {
	Inventory VpcFirewallRuleInventoryView `json:"inventory,omitempty"`
}

// QueryFirewallRuleView QueryFirewallRule
type QueryFirewallRuleView struct {
	Inventories []VpcFirewallRuleInventoryView `json:"inventories,omitempty"`
}

// UpdateFirewallRuleEventView UpdateFirewallRuleEvent
type UpdateFirewallRuleEventView struct {
	Inventory VpcFirewallRuleInventoryView `json:"inventory,omitempty"`
}

