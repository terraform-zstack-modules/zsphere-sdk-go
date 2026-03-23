// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PortForwardingRuleInventoryView PortForwardingRule
type PortForwardingRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	VipIp string `json:"vipIp,omitempty"`
	GuestIp string `json:"guestIp,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	VipPortStart int `json:"vipPortStart,omitempty"`
	VipPortEnd int `json:"vipPortEnd,omitempty"`
	PrivatePortStart int `json:"privatePortStart,omitempty"`
	PrivatePortEnd int `json:"privatePortEnd,omitempty"`
	VmNicUuid string `json:"vmNicUuid,omitempty"`
	ProtocolType string `json:"protocolType,omitempty"`
	State string `json:"state,omitempty"`
	AllowedCidr string `json:"allowedCidr,omitempty"`
}

// DetachPortForwardingRuleEventView DetachPortForwardingRuleEvent
type DetachPortForwardingRuleEventView struct {
	Inventory PortForwardingRuleInventoryView `json:"inventory,omitempty"`
}

// DeletePortForwardingRuleEventView DeletePortForwardingRuleEvent
type DeletePortForwardingRuleEventView struct {
	Success bool `json:"success,omitempty"`
}

// ChangePortForwardingRuleStateEventView ChangePortForwardingRuleStateEvent
type ChangePortForwardingRuleStateEventView struct {
	Inventory PortForwardingRuleInventoryView `json:"inventory,omitempty"`
}

// AttachPortForwardingRuleEventView AttachPortForwardingRuleEvent
type AttachPortForwardingRuleEventView struct {
	Inventory PortForwardingRuleInventoryView `json:"inventory,omitempty"`
}

// QueryPortForwardingRuleView QueryPortForwardingRule
type QueryPortForwardingRuleView struct {
	Inventories []PortForwardingRuleInventoryView `json:"inventories,omitempty"`
}

// UpdatePortForwardingRuleEventView UpdatePortForwardingRuleEvent
type UpdatePortForwardingRuleEventView struct {
	Inventory PortForwardingRuleInventoryView `json:"inventory,omitempty"`
}

// GetVpcAttachedPortForwardingRulesView GetVpcAttachedPortForwardingRules
type GetVpcAttachedPortForwardingRulesView struct {
	Inventories []PortForwardingRuleInventoryView `json:"inventories,omitempty"`
}

// CreatePortForwardingRuleEventView CreatePortForwardingRuleEvent
type CreatePortForwardingRuleEventView struct {
	Inventory PortForwardingRuleInventoryView `json:"inventory,omitempty"`
}

