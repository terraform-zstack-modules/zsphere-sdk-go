// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// HostSchedulingRuleGroupInventoryView HostSchedulingRuleGroup
type HostSchedulingRuleGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
}

// UpdateHostSchedulingRuleGroupEventView UpdateHostSchedulingRuleGroupEvent
type UpdateHostSchedulingRuleGroupEventView struct {
	Inventory HostSchedulingRuleGroupInventoryView `json:"inventory,omitempty"`
}

// CreateHostSchedulingRuleGroupEventView CreateHostSchedulingRuleGroupEvent
type CreateHostSchedulingRuleGroupEventView struct {
	Inventory HostSchedulingRuleGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteHostSchedulingRuleGroupEventView DeleteHostSchedulingRuleGroupEvent
type DeleteHostSchedulingRuleGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryHostSchedulingRuleGroupView QueryHostSchedulingRuleGroup
type QueryHostSchedulingRuleGroupView struct {
	Inventories []HostSchedulingRuleGroupInventoryView `json:"inventories,omitempty"`
}

