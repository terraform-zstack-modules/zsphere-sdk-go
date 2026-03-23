// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AddingNewVmRuleInventoryView AddingNewVmRule
type AddingNewVmRuleInventoryView struct {
	BaseInfoView
	BaseTimeView
	AdjustmentType string `json:"adjustmentType,omitempty"`
	AdjustmentValue int `json:"adjustmentValue,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Cooldown int64 `json:"cooldown,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	SystemTags []string `json:"systemTags,omitempty"`
	ScalingGroupUuid string `json:"scalingGroupUuid,omitempty"`
	RuleTriggers []AutoScalingRuleTriggerInventoryView `json:"ruleTriggers,omitempty"`
}

