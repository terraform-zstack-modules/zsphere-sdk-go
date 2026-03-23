// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateHaStrategyConditionParamDetail UpdateHaStrategyCondition detail param
type UpdateHaStrategyConditionParamDetail struct {
	Name string `json:"name,omitempty"`
	State *string `json:"state,omitempty"`
}

// UpdateHaStrategyConditionParam UpdateHaStrategyCondition request param
type UpdateHaStrategyConditionParam struct {
	BaseParam
	Params UpdateHaStrategyConditionParamDetail `json:"updateHaStrategyCondition"`
}
