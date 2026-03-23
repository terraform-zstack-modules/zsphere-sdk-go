// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SecurityGroupRuleInventoryView SecurityGroupRule
type SecurityGroupRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
	Type string `json:"type,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	State string `json:"state,omitempty"`
	Priority int `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
	SrcIpRange string `json:"srcIpRange,omitempty"`
	DstIpRange string `json:"dstIpRange,omitempty"`
	SrcPortRange string `json:"srcPortRange,omitempty"`
	DstPortRange string `json:"dstPortRange,omitempty"`
	Action string `json:"action,omitempty"`
	RemoteSecurityGroupUuid string `json:"remoteSecurityGroupUuid,omitempty"`
	AllowedCidr string `json:"allowedCidr,omitempty"`
	StartPort int `json:"startPort,omitempty"`
	EndPort int `json:"endPort,omitempty"`
}

// ChangeSecurityGroupRuleEventView ChangeSecurityGroupRuleEvent
type ChangeSecurityGroupRuleEventView struct {
	Inventory SecurityGroupRuleInventoryView `json:"inventory,omitempty"`
}

// ValidateSecurityGroupRuleView ValidateSecurityGroupRule
type ValidateSecurityGroupRuleView struct {
	Available bool `json:"available,omitempty"`
	Code string `json:"code,omitempty"`
	Reason string `json:"reason,omitempty"`
	Success bool `json:"success,omitempty"`
}

// AddSecurityGroupRuleEventView AddSecurityGroupRuleEvent
type AddSecurityGroupRuleEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

// QuerySecurityGroupRuleView QuerySecurityGroupRule
type QuerySecurityGroupRuleView struct {
	Inventories []SecurityGroupRuleInventoryView `json:"inventories,omitempty"`
}

// DeleteSecurityGroupRuleEventView DeleteSecurityGroupRuleEvent
type DeleteSecurityGroupRuleEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

