// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingRuleInventoryView AutoScalingRule
type AutoScalingRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Cooldown int64 `json:"cooldown,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
	ScalingGroupUuid string `json:"scalingGroupUuid,omitempty"`
	RuleTriggers []AutoScalingRuleTriggerInventoryView `json:"ruleTriggers,omitempty"`
}

// CreateAutoScalingRuleEventView CreateAutoScalingRuleEvent
type CreateAutoScalingRuleEventView struct {
	Inventory AutoScalingRuleInventoryView `json:"inventory,omitempty"`
}

// DeleteAutoScalingRuleEventView DeleteAutoScalingRuleEvent
type DeleteAutoScalingRuleEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateAutoScalingRuleEventView UpdateAutoScalingRuleEvent
type UpdateAutoScalingRuleEventView struct {
	Inventory AutoScalingRuleInventoryView `json:"inventory,omitempty"`
}

// QueryAutoScalingRuleView QueryAutoScalingRule
type QueryAutoScalingRuleView struct {
	Inventories []AutoScalingRuleInventoryView `json:"inventories,omitempty"`
}

