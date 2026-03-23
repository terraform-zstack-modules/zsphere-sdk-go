// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmSchedulingRuleGroupInventoryView VmSchedulingRuleGroup
type VmSchedulingRuleGroupInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	Appliance string `json:"appliance,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
}

// UpdateVmSchedulingRuleGroupEventView UpdateVmSchedulingRuleGroupEvent
type UpdateVmSchedulingRuleGroupEventView struct {
	Inventory VmSchedulingRuleGroupInventoryView `json:"inventory,omitempty"`
}

// DeleteVmSchedulingRuleGroupEventView DeleteVmSchedulingRuleGroupEvent
type DeleteVmSchedulingRuleGroupEventView struct {
	Success bool `json:"success,omitempty"`
}

// QueryVmSchedulingRuleGroupView QueryVmSchedulingRuleGroup
type QueryVmSchedulingRuleGroupView struct {
	Inventories []VmSchedulingRuleGroupInventoryView `json:"inventories,omitempty"`
}

// CreateVmSchedulingRuleGroupEventView CreateVmSchedulingRuleGroupEvent
type CreateVmSchedulingRuleGroupEventView struct {
	Inventory VmSchedulingRuleGroupInventoryView `json:"inventory,omitempty"`
}

