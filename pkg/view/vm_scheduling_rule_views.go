// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmSchedulingRuleInventoryView VmSchedulingRule
type VmSchedulingRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	Rule string `json:"rule,omitempty"`
	Mode string `json:"mode,omitempty"`
	Description string `json:"description,omitempty"`
	Policy string `json:"policy,omitempty"`
	Version string `json:"version,omitempty"`
	Type string `json:"type,omitempty"`
	Appliance string `json:"appliance,omitempty"`
	ZoneUuid string `json:"zoneUuid,omitempty"`
	State string `json:"state,omitempty"`
	Usages []AffinityGroupUsageInventoryView `json:"usages,omitempty"`
}

// ValidateVmSchedulingRuleView ValidateVmSchedulingRule
type ValidateVmSchedulingRuleView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateVmSchedulingRuleEventView UpdateVmSchedulingRuleEvent
type UpdateVmSchedulingRuleEventView struct {
	Inventory VmSchedulingRuleInventoryView `json:"inventory,omitempty"`
}

// QueryVmSchedulingRuleView QueryVmSchedulingRule
type QueryVmSchedulingRuleView struct {
	Inventories []VmSchedulingRuleInventoryView `json:"inventories,omitempty"`
}

// ChangeVmSchedulingRuleStateEventView ChangeVmSchedulingRuleStateEvent
type ChangeVmSchedulingRuleStateEventView struct {
	Inventory VmSchedulingRuleInventoryView `json:"inventory,omitempty"`
}

