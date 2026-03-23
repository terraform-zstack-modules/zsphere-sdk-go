// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccessControlRuleInventoryView AccessControlRule
type AccessControlRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Rule string `json:"rule,omitempty"`
	Strategy string `json:"strategy,omitempty"`
}

// DeleteAccessControlRuleEventView DeleteAccessControlRuleEvent
type DeleteAccessControlRuleEventView struct {
	Success bool `json:"success,omitempty"`
}

// AddAccessControlRuleEventView AddAccessControlRuleEvent
type AddAccessControlRuleEventView struct {
	Inventory AccessControlRuleInventoryView `json:"inventory,omitempty"`
}

// UpdateAccessControlRuleEventView UpdateAccessControlRuleEvent
type UpdateAccessControlRuleEventView struct {
	Inventory AccessControlRuleInventoryView `json:"inventory,omitempty"`
}

// QueryAccessControlRuleView QueryAccessControlRule
type QueryAccessControlRuleView struct {
	Inventories []AccessControlRuleInventoryView `json:"inventories,omitempty"`
}

