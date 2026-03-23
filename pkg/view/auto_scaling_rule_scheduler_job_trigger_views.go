// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AutoScalingRuleSchedulerJobTriggerInventoryView AutoScalingRuleSchedulerJobTrigger
type AutoScalingRuleSchedulerJobTriggerInventoryView struct {
	BaseInfoView
	BaseTimeView
	SchedulerJobUuid string `json:"schedulerJobUuid,omitempty"`
	Type string `json:"type,omitempty"`
	RuleUuid string `json:"ruleUuid,omitempty"`
	Description string `json:"description,omitempty"`
	State string `json:"state,omitempty"`
}

