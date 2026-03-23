// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteAutoScalingRuleParamDetail DeleteAutoScalingRule detail param
type DeleteAutoScalingRuleParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingRuleParam DeleteAutoScalingRule request param
type DeleteAutoScalingRuleParam struct {
	BaseParam
	Params DeleteAutoScalingRuleParamDetail `json:"deleteAutoScalingRule"`
}
// UpdateAutoScalingRuleParamDetail UpdateAutoScalingRule detail param
type UpdateAutoScalingRuleParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Cooldown *int64 `json:"cooldown,omitempty"`
}

// UpdateAutoScalingRuleParam UpdateAutoScalingRule request param
type UpdateAutoScalingRuleParam struct {
	BaseParam
	Params UpdateAutoScalingRuleParamDetail `json:"updateAutoScalingRule"`
}
