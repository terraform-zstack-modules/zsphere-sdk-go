// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VpcFirewallRuleTemplateInventoryView VpcFirewallRuleTemplate
type VpcFirewallRuleTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Action string `json:"action,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	DestPort string `json:"destPort,omitempty"`
	SourcePort string `json:"sourcePort,omitempty"`
	SourceIp string `json:"sourceIp,omitempty"`
	DestIp string `json:"destIp,omitempty"`
	AllowStates string `json:"allowStates,omitempty"`
	TcpFlag string `json:"tcpFlag,omitempty"`
	IcmpTypeName string `json:"icmpTypeName,omitempty"`
	RuleNumber int `json:"ruleNumber,omitempty"`
	EnableLog bool `json:"enableLog,omitempty"`
	State string `json:"state,omitempty"`
	IsDefault bool `json:"isDefault,omitempty"`
	Description string `json:"description,omitempty"`
	AccountUuid string `json:"accountUuid,omitempty"`
}

// QueryFirewallRuleTemplateView QueryFirewallRuleTemplate
type QueryFirewallRuleTemplateView struct {
	Inventories []VpcFirewallRuleTemplateInventoryView `json:"inventories,omitempty"`
}

// CreateFirewallRuleTemplateEventView CreateFirewallRuleTemplateEvent
type CreateFirewallRuleTemplateEventView struct {
	Inventory VpcFirewallRuleTemplateInventoryView `json:"inventory,omitempty"`
}

// UpdateFirewallRuleTemplateEventView UpdateFirewallRuleTemplateEvent
type UpdateFirewallRuleTemplateEventView struct {
	Inventory VpcFirewallRuleTemplateInventoryView `json:"inventory,omitempty"`
}

