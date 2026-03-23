// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingRuleTriggerInventoryView AutoScalingRuleTrigger
type AutoScalingRuleTriggerInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	RuleUuid string `json:"ruleUuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

// QueryAutoScalingRuleTriggerView QueryAutoScalingRuleTrigger
type QueryAutoScalingRuleTriggerView struct {
	Inventories []AutoScalingRuleTriggerInventoryView `json:"inventories,omitempty"`
}

// DeleteAutoScalingRuleTriggerEventView DeleteAutoScalingRuleTriggerEvent
type DeleteAutoScalingRuleTriggerEventView struct {
	Success bool `json:"success,omitempty"`
}

// CreateAutoScalingRuleTriggerEventView CreateAutoScalingRuleTriggerEvent
type CreateAutoScalingRuleTriggerEventView struct {
	Inventory AutoScalingRuleTriggerInventoryView `json:"inventory,omitempty"`
}

