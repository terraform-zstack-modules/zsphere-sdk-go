// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SecurityGroupInventoryView SecurityGroup
type SecurityGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
	IpVersion int `json:"ipVersion,omitempty"`
	Rules []SecurityGroupRuleInventoryView `json:"rules,omitempty"`
	AttachedL3NetworkUuids []string `json:"attachedL3NetworkUuids,omitempty"`
}

// ChangeSecurityGroupStateEventView ChangeSecurityGroupStateEvent
type ChangeSecurityGroupStateEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

// DetachSecurityGroupFromL3NetworkEventView DetachSecurityGroupFromL3NetworkEvent
type DetachSecurityGroupFromL3NetworkEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

// QuerySecurityGroupView QuerySecurityGroup
type QuerySecurityGroupView struct {
	Inventories []SecurityGroupInventoryView `json:"inventories,omitempty"`
}

// AttachSecurityGroupToL3NetworkEventView AttachSecurityGroupToL3NetworkEvent
type AttachSecurityGroupToL3NetworkEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

// ChangeSecurityGroupRuleStateEventView ChangeSecurityGroupRuleStateEvent
type ChangeSecurityGroupRuleStateEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

// CreateSecurityGroupEventView CreateSecurityGroupEvent
type CreateSecurityGroupEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteSecurityGroupEventView DeleteSecurityGroupEvent
type DeleteSecurityGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateSecurityGroupRulePriorityEventView UpdateSecurityGroupRulePriorityEvent
type UpdateSecurityGroupRulePriorityEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

// UpdateSecurityGroupEventView UpdateSecurityGroupEvent
type UpdateSecurityGroupEventView struct {
	Inventory SecurityGroupInventoryView `json:"inventory,omitempty"`
}

