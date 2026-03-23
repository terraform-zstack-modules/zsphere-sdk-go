// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteAutoScalingRuleTriggerParamDetail DeleteAutoScalingRuleTrigger detail param
type DeleteAutoScalingRuleTriggerParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteAutoScalingRuleTriggerParam DeleteAutoScalingRuleTrigger request param
type DeleteAutoScalingRuleTriggerParam struct {
	BaseParam
	Params DeleteAutoScalingRuleTriggerParamDetail `json:"deleteAutoScalingRuleTrigger"`
}
